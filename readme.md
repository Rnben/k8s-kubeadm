- [Creating a single-master-cluster](#creating-a-single-master-cluster)
  - [Init cluster](#init-cluster)
    - [VM Configuration](#vm-configuration)
    - [Prerequisites](#prerequisites)
    - [Create vm](#create-vm)
    - [Create kubernetes cluster](#create-kubernetes-cluster)
  - [Create](#create)

# Creating a single-master-cluster

## Init cluster

### VM Configuration

| os  | kubernetes | CPU | MEM |
|---|---|---|---| 
| centos7.4 | 1.11.2 | 2C | 4G |

### Prerequisites
1. Install brew
```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

2. Install vagrant and virtualbox

```
brew cask install vagrant
brew cask install virtualbox
```

### Create vm

3. Run this cmd in directory with Vagrantfile
```
vagrant up
```

### Create kubernetes cluster

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
curl -O https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

kubectl apply -f  kube-flannel.yml
```

7. Test dns

```
kubectl run curl --image=radial/busyboxplus:curl -it

nslookup kubernetes.default
Server:    10.96.0.10
Address 1: 10.96.0.10 kube-dns.kube-system.svc.cluster.local

Name:      kubernetes.default
Address 1: 10.96.0.1 kubernetes.default.svc.cluster.local
```

8. kube-proxy开启ipvs

```
kubectl edit cm kube-proxy -n kube-system

kubectl get pod -n kube-system | grep kube-proxy | awk '{system("kubectl delete pod "$1" -n kube-system")}'

kubectl get pod --all-namespaces -o wide
```

9. Join node

**Option 1**
```
# join node
kubeadm join --token <token> <master-ip>:<master-port> --discovery-token-ca-cert-hash sha256:<hash>

# list token
kubeadm token list

# discovery-token-ca-cert-hash
root@ubuntu:/home/cong# openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'

```
**Option 2**
```
kubeadm token create --print-join-command
```

9. delete node

```
kubectl drain <node name> --delete-local-data --force --ignore-daemonsets
kubectl delete node <node name>
```

## Create

[pod cidr not assgned](https://github.com/coreos/flannel/issues/728)
[flannel issues 39701](https://github.com/kubernetes/kubernetes/issues/39701)