
<br>

##  Instructions on how to setup a Kubernetes cluster.
You may use Ubuntu virtual machines/servers on premise or on cloud(aws,gcp,azure,hetzner,linode,digital ocean,contabo etc), even by using a mix of providers.

Once you have done a setup of all the instances with the ssh-keys:

### SSH into the instances
Search for "multiplexing/broadcast all" options on your terminal emulator software, in order to send all the commands at once to all sessions.

### Commands to run on all the nodes


 1- Disable swap
```
    swapoff -a
```
 2- Enable bridge network
    Since kubernetes has deprecated docker, we use containerd.
```
    sudo modprobe overlay
    sudo modprobe br_netfilter
```
```
    cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
    overlay
    br_netfilter
    EOF
```

 3- Setup required sysctl parameters.
```
    cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
    net.bridge.bridge-nf-call-iptables  = 1
    net.ipv4.ip_forward                 = 1
    net.bridge.bridge-nf-call-ip6tables = 1
    EOF
```

 4- Apply sysctl parameters
```
    sudo sysctl --system
```
 5- Install containerd
```
    sudo apt-get update 
    sudo apt-get install -y containerd
```
 6-Create the containerd config file(toml)
```
    sudo mkdir -p /etc/containerd
    sudo containerd config default | sudo tee /etc/containerd/config.toml
```

 7- Enable containerd to run with cgroup driver. If you do # not do this, kube init will fail.  
 On the config file find
             SystemdCgroup = false
     and change it to :
          #   SystemdCgroup = true
```              
    sudo nano /etc/containerd/config.toml
```
 8- Restart containerd to load the new config.
```
    sudo systemctl restart containerd
```
 9-Add Google's apt repository gpg key
```
    curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
```

10- Add the Kubernetes apt repository
```
    sudo bash -c 'cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
    deb https://apt.kubernetes.io/ kubernetes-xenial main
    EOF'
```
 11- Install  kubeadm, kubelet and kubectl. The version 
 below is an example. Check the available versions with 
 this command
```
    sudo apt-get update
    apt-cache policy kubelet | head -n 20 
```
```
    VERSION=1.21.0-00
```
```    
    sudo apt-get install -y kubelet=$VERSION kubeadm=$VERSION kubectl=$VERSION
    sudo apt-mark hold kubelet kubeadm kubectl containerd
```

12 - Run mark hold to stop the automatic update of these packages.
```    
    sudo apt-mark hold kubelet kubeadm kubectl containerd
```

After the above commands are successfully run on all the worker nodes. Below steps can be followed to initialize the Kubernetes cluster.


### On Master Node

Run the command kubeadm init to expose the API server. 

13- You can also pass something like `--pod-network-cidr=192.xxx.xx.xx/xx` but be careful with the # underlying pod network so that it does not overlaps with other machines on # the network.
```
    kubeadm init --apiserver-advertise-address="ip_address_of_master" 

```
Install CNI plugin

The below command can be run on the leader node to install the CNI plugin

```
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

OR

```
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
```


### On Worker Nodes

Run the join command that was printed out on step 13.