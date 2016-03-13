package main

import (
   "flag"
   "fmt"
   "io/ioutil"
   "net/http"
   "log"
   "time"
   "os"
   "os/signal"
   "syscall"
)

const (
   pidFileLocation = "/tmp/web_go.pid"
)

func greetFromHost(w http.ResponseWriter, r *http.Request) {
   hostname, _ := os.Hostname()
   fmt.Fprintf(w, "This host is %s date: %v\n", hostname, time.Now())
}

func createPidFile() {

   pid := []byte(fmt.Sprintf("%d\n", os.Getpid()))
   err := ioutil.WriteFile(pidFileLocation, pid, 0644)
   if err != nil {
     log.Fatal("Writing Pid file failed: ", err)
   }
}

func removePidFile(stopChan chan bool) {
    log.Printf("Removing pid file...")
    err := os.Remove(pidFileLocation)
    if err != nil {
       log.Printf("Unable to remove %d\n", pidFileLocation)
    }
    log.Printf("Done\n")
    stopChan <- true
}

var port int

func init() {
   flag.IntVar(&port, "port", 9090, "Port to listen on")
}

func main() {

   flag.Parse()

   createPidFile()
   signalChan := make(chan os.Signal, 1)
   stopChan := make(chan bool)
   go func() {
       defer removePidFile(stopChan)
       <-signalChan
   }()
   signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

   go func() {
       http.HandleFunc("/", greetFromHost)

       addr := fmt.Sprintf(":%d",port)
       log.Printf("Web App running on http://0.0.0.0%s\n", addr)
       err := http.ListenAndServe(addr, nil)
       if err != nil {
           log.Fatal("ListenAndServe: ", err)
       }
   }()

   <-stopChan
   log.Printf("Main done!\n")
}
