# -*- mode: ruby -*-
# vi: set ft=ruby :

$install_webapp_script = <<SCRIPT

mkdir /home/vagrant/go
chown vagrant:vagrant /home/vagrant/go
cd /tmp
wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.6.linux-amd64.tar.gz
echo 'export GOROOT=/usr/local/go' >> /home/vagrant/.bashrc
echo 'export GOPATH=/home/vagrant/go' >> /home/vagrant/.bashrc
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /home/vagrant/.bashrc
sudo rm /tmp/go1.6.linux-amd64.tar.gz

echo "Configuring Service..."
sudo cp /vagrant/webapp/web_go.service /etc/systemd/system
sudo systemctl enable web_go
sudo systemctl start web_go

echo "Service Configuration Done!"

SCRIPT

$install_haproxy_script = <<SCRIPT

echo "Instatlling HAProxy..."
sudo apt-get install haproxy
sudo systemctl stop haproxy
sudo cp /vagrant/haproxy.cfg /etc/haproxy/haproxy.cfg
sudo systemctl start haproxy

SCRIPT

Vagrant.configure(2) do |config|

  config.vm.box = "ubuntu/wily64"

  config.vm.define "n1" do |n1|
    n1.vm.hostname = "n1"
    n1.vm.provision "shell", inline: $install_webapp_script
    n1.vm.network "private_network", ip: "172.20.20.11"
    n1.vm.network "forwarded_port", guest: 9090, host: 19090
  end

  config.vm.define "n2" do |n2|
     n2.vm.hostname = "n2"
     n2.vm.provision "shell", inline: $install_webapp_script
     n2.vm.network "private_network", ip: "172.20.20.12"
     n2.vm.network "forwarded_port", guest: 9090, host: 29090
  end

  config.vm.define "n3" do |n3|
     n3.vm.hostname = "n3"
     n3.vm.provision "shell", inline: $install_haproxy_script
     n3.vm.network "private_network", ip: "172.20.20.13"
     n3.vm.network "forwarded_port", guest: 9090, host: 39090
  end

end
