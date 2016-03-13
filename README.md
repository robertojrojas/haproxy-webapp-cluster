# haproxy-webapp-cluster
A Simple Go Web App Behind a Load Balancer (HAProxy).
A Simple Go Web App Behind a Load Balancer (HAProxy).
This app simply displays the name of the node executing it
and the current datetime.

We use vagrant to create a cluster.
The Vagrantfile creates 3 nodes (n1, n2, n3):
  n1 - Web App instance 1 172.20.20.11
  n2 - Web App instance 2 172.20.20.12
  n3 - HAProxy serving site on port :39090. This port will be forwarded
       to the host. 172.20.20.13 
       The load balancer will roundrobin between n1, n2

These nodes are connected with private network:
  172.20.20.x
  
To build:
---------

cd webapp/
./build.sh
cd ..

To Run:
-------

vagrant up


To Test:
--------

 curl http://<HOST>:39090/

Here is sample output:
   This host is n1 date: 2016-03-12 21:23:40.456791077 +0000 UTC
   ....
   This host is n2 date: 2016-03-12 21:23:41.456791077 +0000 UTC
