# Creating a single-master-cluster

> two step to create a single master cluster
>* create vm
>* kubeadm init

## VM Configuration

| os  | kubernetes | CPU | MEM |
|---|---|---|---| 
| centos7.4 | 1.11.2 | 2C | 4G |

## Prerequisites
1. Install brew
```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

2. Install vagrant and virtualbox

```
brew cask install vagrant
brew cask install virtualbox
```

## Create vm

3. Run this cmd in directory with Vagrantfile
```
vagrant up
```

## Create kubernetes cluster

4. Init kubernetes master
```
kubeadm init --config=kubeadm.config

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

5. Remove master taint

```
kubectl taint nodes --all node-role.kubernetes.io/master-
```

6. Install flannel
```
kubectl create -f flannel.yaml
```

7. Check cluster

```
kubectl cluster-info
kubectl get cs
kubectl get nodes
```




