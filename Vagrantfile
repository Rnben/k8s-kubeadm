# -*- mode: ruby -*-
# vi: set ft=ruby :

ENV["LC_ALL"] = "en_US.UTF8"

Vagrant.configure("2") do |config|
  config.vm.provision "shell", inline: $script

  config.vm.define "master" do |master|
    master.vm.box = "bento/centos-7.4"
    master.vm.hostname = "master"
    master.ssh.username = "root"
    master.ssh.password = "vagrant"
    master.ssh.insert_key = true
    master.vm.box_check_update = false
    master.vm.network "private_network", ip: "11.11.11.110"
    master.vm.provider "virtualbox" do |vb|
          vb.customize ["modifyvm", :id, "--name", "master", "--memory", "2048"]
          vb.customize ["modifyvm", :id, "--cpus", "2"]
    end
  end

  (1..2).each do |i|
    config.vm.define "node#{i}" do |node|
	node.vm.box = "bento/centos-7.4"
	node.vm.hostname="node#{i}"
	node.vm.network "private_network", ip: "11.11.11.10#{i}"
        node.vm.provider "virtualbox" do |vb|
          vb.customize ["modifyvm", :id, "--name", "node#{i}", "--memory", "1024"]
          vb.customize ["modifyvm", :id, "--cpus", "1"]
    end
	node.vm.provision "shell", inline: <<-SHELL
        echo "this is node"
	SHELL
 end
end

$script = <<-SCRIPT
swapoff -a
swapoff -a && sed -i 's/.*swap.*/#&/' /etc/fstab

setenforce 0
sed -i "s/^SELINUX=enforcing/SELINUX=disabled/g" /etc/selinux/config

# 开启forward
iptables -P FORWARD ACCEPT

# 配置转发相关参数
cat <<EOF > /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
vm.swappiness=0
EOF
sysctl --system

# hosts配置
cat <<EOF > /etc/hosts
11.11.11.110 master
11.11.11.101 node1
11.11.11.102 node2
EOF

mkdir /etc/yum.repos.d/bak && mv /etc/yum.repos.d/*.repo /etc/yum.repos.d/bak
curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.cloud.tencent.com/repo/centos7_base.repo
curl -o /etc/yum.repos.d/epel.repo http://mirrors.cloud.tencent.com/repo/epel-7.repo
yum clean all && yum makecache

cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

curl -o /etc/yum.repos.d/docker-ce.repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum install -y docker-ce-18.06.1.ce-3.el7
systemctl enable docker && systemctl start docker
docker --version

yum install -y kubectl-1.14.6 kubelet-1.14.6 kubeadm-1.14.6
systemctl enable kubelet

cat > /root/install.sh <<EOF
kubeadm init --kubernetes-version=1.14.6 \
--apiserver-advertise-address=11.11.11.110 \
--image-repository registry.aliyuncs.com/google_containers \
--service-cidr=10.1.0.0/16 \
--pod-network-cidr=10.244.0.0/16
EOF

SCRIPT
