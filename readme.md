# For mac user
> Other os, the installation method for vagrant will be different

## Prerequisites
1. Install `brew`
`/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`

2. Install vagrant
`brew cask install vagrant`

## Ready master node
3. Run this cmd in directory with Vagrantfile
`vagrant up`

## Init kubernetes
4. Init cluster
`kubeadm init --config=kubeadm.config`





