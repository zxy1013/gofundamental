Kubernetes详细教程

`https://www.bilibili.com/video/BV1Qv41167ck?from=search&seid=2483670257694650823&spm_id_from=333.337.0.0`

### 1. Kubernetes介绍

#### 1.1 应用部署方式演变

在部署应用程序的方式上，主要经历了三个时代：

- **传统部署**：互联网早期，会直接将应用程序部署在物理机上

  > 优点：简单，不需要其它技术的参与
  >
  > 缺点：不能为应用程序定义资源使用边界，很难合理地分配计算资源，而且程序之间容易产生影响。若是一个电脑部署一个太浪费资源。

- **虚拟化部署**：可以在一台物理机上运行多个虚拟机，每个虚拟机都是独立的一个环境

  > 优点：程序环境不会相互产生影响，提供了一定程度的安全性
  >
  > 缺点：增加了操作系统，浪费了部分资源

- **容器化部署**：与虚拟化类似，但是**共享了操作系统**

  > 优点：
  >
  > 1 可以保证每个容器拥有自己的文件系统、CPU、内存、进程空间等
  >
  > 2 运行应用程序所需要的资源都被容器包装，并和底层基础架构解耦
  >
  > 3 容器化的应用程序可以跨云服务商、跨Linux操作系统发行版进行部署

![image-20200505183738289](Kubenetes.assets/image-20200505183738289.png)

> 容器化部署方式给带来很多的便利，但是也会出现一些问题，比如说：
>
> - 一个容器故障停机了，怎么样让另外一个容器立刻启动去替补停机的容器
> - 当并发访问量变大的时候，怎么样做到横向扩展容器数量 增大减少
>

这些容器管理的问题统称为**容器编排**问题，为了解决这些容器编排问题，就产生了一些容器编排的软件：

> - **Swarm**：Docker自己的容器编排工具
> - **Mesos**：Apache的一个资源统一管控的工具，需要和Marathon结合使用
> - **Kubernetes**：Google开源的的容器编排工具
>

![image-20200524150339551](Kubenetes.assets/image-20200524150339551.png)

#### 1.2 kubernetes简介

![image-20200406232838722](Kubenetes.assets/image-20200406232838722.png)

kubernetes，是一个全新的基于容器技术的分布式架构领先方案，是谷歌严格保密十几年的秘密武器----Borg系统的一个开源版本，于2014年9月发布第一个版本，2015年7月发布第一个正式版本。

kubernetes的本质是**一组服务器集群**，它可以在集群的每个节点上运行特定的程序，来对节点中的容器进行管理。目的是实现资源管理的自动化，主要提供了如下的主要功能：

> - **自我修复**：一旦某一个容器崩溃，能够在1秒中左右迅速启动新的容器
>
> - **弹性伸缩**：可以根据需要，自动对集群中正在运行的容器数量进行调整
>
> - **服务发现**：服务可以通过自动发现的形式找到它所依赖的服务。比如nginx需要mysql和redis，k8s内部以自动发现的形式找。
>
> - **负载均衡**：如果一个服务起动了多个容器，能够自动实现请求的负载均衡。可以自定义策略。
>
> - **版本回退**：如果发现新发布的程序版本有问题，可以立即回退到原来的版本。mysql升级了 有问题，则可以回退。
>
> - **存储编排**：可以根据容器自身的需求自动创建存储卷，mysql在存储中申请存储卷。
>
>   ![image-20200406232838722](Kubenetes.assets/11.png)
>

#### 1.3 kubernetes组件

一个kubernetes集群主要是由**控制节点(master)**和**工作节点(node)**构成，每个节点上都会**安装不同**的组件。

**master：集群的控制平面，负责集群的决策 ( 管理 )**

> **ApiServer** : 资源操作的唯一入口，用户对集群的管理，接收用户输入的命令，提供认证、授权、API注册和发现等机制。安排k8s跑一个nginx服务。
>
> **Scheduler** : 负责集群资源调度，按照预定的调度策略将Pod调度到相应的node节点上。计算活应该安排给哪个node负责。
>
> **ControllerManager** : 负责维护集群的状态，比如程序部署安排、故障检测、自动扩展、滚动更新等。安排活给node1节点。
>
> **Etcd** ：负责存储集群中各种资源对象的信息。活完成的信息情况记录在etcd数据库中。

**node：集群的数据平面，负责为容器提供运行环境 ( 干活 )**

> **Kubelet** : 负责维护容器的生命周期，即通过控制docker，来创建、更新、销毁容器。接收ControllerManager发送的信息，控制容器将nginx跑起来。
>
> **Docker** : 负责节点上容器的各种操作
>
> **KubeProxy** : 负责提供集群内部的服务发现和负载均衡。提供nginx的对外访问。

![image-20200406184656917](Kubenetes.assets/image-20200406184656917.png)

下面，以部署一个nginx服务来说明kubernetes系统各个组件调用关系：

> 1. 首先要明确，一旦kubernetes环境启动之后，master和node都会将自身的信息存储到etcd数据库中。为了好管理，清楚领导被领导关系。
>
> 2. 一个nginx服务的安装请求会首先被发送到master节点的apiServer组件
>
> 3. apiServer组件会调用scheduler组件来决定到底应该把这个服务安装到哪个node节点上，在此时，它会从etcd中读取各个node节点的信息，然后按照一定的算法进行选择，并将结果告知apiServer。
>
> 4. apiServer调用controller-manager去调度Node节点安装nginx服务。
>
> 5. kubelet接收到指令后，会通知docker，然后由docker来启动一个nginx的pod，pod是kubernetes的最小操作单元，容器必须跑在pod中。
>
> 6. 至此，一个nginx服务就运行了，如果需要访问nginx，就需要通过kube-proxy来对pod产生访问的代理
>
> 这样，外界用户就可以访问集群中的nginx服务了
>

#### 1.4 kubernetes概念

> **Master**：集群控制节点，每个集群需要至少一个master节点负责集群的管控。
>
> **Node**：工作负载节点，由master分配容器到这些node工作节点上，然后node节点上的docker负责容器的运行。
>
> **Pod**：kubernetes的最小控制单元，容器都是运行在pod中的，一个pod中可以有1个或者多个容器。k8s通过控制pod进而控制容器进而控制程序。
>
> **Controller**：控制器，通过它来实现对pod的管理，比如启动pod、停止pod、伸缩pod的数量等等。多种控制器均有自己的使用场景。
>
> **Service**：pod对外服务的统一入口，下面可以维护者同一类的多个pod。外部访问tomcat，内部实现选择器可以通过标签，实现负载均衡和加权负载。
>
> **Label**：标签，用于对pod进行分类，同一类pod会拥有相同的标签
>
> **NameSpace**：命名空间，用来隔离pod的运行环境。默认情况下，在一个k8s中的所有pod可以相互访问，NameSpace实现了组划分。同组才可相互访问。
>
> ![image-20200406184656917](Kubenetes.assets/22.png)







### 2. kubernetes集群环境搭建

#### 2.1 前置知识点

目前生产部署Kubernetes 集群主要有两种方式：

> 一主多从：一台master和多台Node，搭建简单，有单机故障风险。
>
> 多主多从：多台master和多台Node，搭建麻烦，安全性高，适用于生产环境。

![image-20200404094800622](Kubenetes.assets/image-20200404094800622.png)

#### 2.2 kubeadm 部署方式介绍

> **minikube：**一个用于快速搭建单节点kubernetes的工具。
>
> **二进制包：**
>
> 从github 下载发行版的二进制包，手动部署每个组件，组成Kubernetes 集群。
>
> **kubeadm：**
>
> Kubeadm 是一个K8s 部署工具，用于快速部署Kubernetes 集群。官方地址：https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/
>
> Kubeadm 降低部署门槛，但屏蔽了很多细节，遇到问题很难排查。如果想更容易可控，推荐使用二进制包部署Kubernetes 集群，虽然手动部署麻烦点，期间可以学习很多工作原理，也利于后期维护。
>
> kubeadm 是官方社区推出的一个用于快速部署kubernetes 集群的工具，这个工具能通过两条指令完成一个kubernetes 集群的部署：
>
> - 创建一个Master 节点`kubeadm init`
> - 将Node 节点加入到当前集群中`$ kubeadm join <Master 节点的IP 和端口>`

#### 2.3 安装要求

在开始之前，部署Kubernetes 集群机器需要满足以下几个条件：

> - 一台或多台机器 一主两从centos7，软件选择基础设施服务器
> - 硬件配置：2GB 或更多RAM，2 个CPU 或更多CPU，硬盘50GB 或更多
> - 集群中所有机器之间网络互通
> - 可以访问外网，需要拉取镜像
> - 安装位置：自动分区
> - 配置主机和网络
>

#### 2.4 最终目标

> - 在所有节点上安装Docker 和kubeadm
> - 部署Kubernetes Master
> - 部署容器网络插件
> - 部署Kubernetes Node，将节点加入Kubernetes 集群中
> - 部署Dashboard Web 页面，可视化查看Kubernetes 资源
>

#### 2.5 准备环境

![image-20210609000002940](Kubenetes.assets/image-20210609000002940.png)

| 角色   | IP地址      | 组件                              |
| :----- | :---------- | :-------------------------------- |
| master | 192.168.5.3 | docker，kubectl，kubeadm，kubelet |
| node1  | 192.168.5.4 | docker，kubectl，kubeadm，kubelet |
| node2  | 192.168.5.5 | docker，kubectl，kubeadm，kubelet |

#### 2.6 环境初始化

三个服务器均需要

##### 2.6.1 检查操作系统的版本

```powershell
# 此方式下安装kubernetes集群要求Centos版本要在7.5或之上
[root@master ~]# cat /etc/redhat-release
Centos Linux 7.5.1804 (Core)
```

##### 2.6.2 主机名解析

为了方便集群节点间的直接调用，在这里配置一下主机名解析，企业中推荐使用内部DNS服务器`vim /etc/hosts`

```powershell
# 主机名成解析 编辑三台服务器的/etc/hosts文件，添加下面内容
192.168.5.3 master
192.168.5.4 node1
192.168.5.5 node2
```

##### 2.6.3 时间同步

kubernetes要求集群中的节点时间必须精确一直，这里使用chronyd服务从网络同步时间，企业中建议配置内部的时间同步服务器

```powershell
# 启动chronyd服务
[root@master ~]# systemctl start chronyd
# 设置chronyd服务开机自启
[root@master ~]# systemctl enable chronyd
# chronyd服务开机启动稍等几秒钟，就可以使用date命令验证时间了
[root@master ~]# date
```

##### 2.6.4  禁用iptable和firewalld服务

kubernetes和docker 在运行的中会产生大量的iptables规则，实现转发和路由，为了不让系统规则跟它们混淆，测试环境可以直接关闭系统的规则。

```powershell
# 1 关闭firewalld服务
[root@master ~]# systemctl stop firewalld
[root@master ~]# systemctl disable firewalld
# 2 关闭iptables服务
[root@master ~]# systemctl stop iptables
[root@master ~]# systemctl disable iptables
```

##### 2.6.5 禁用selinux

selinux是linux系统下的一个安全服务，如果不关闭它，在安装集群中会产生各种各样的奇葩问题

```powershell
# vim /etc/selinux/config 文件，修改SELINUX的值为disable
# 注意修改完毕之后需要重启linux服务生效
SELINUX=disabled
```

##### 2.6.6 禁用swap分区

swap分区指的是虚拟内存分区，它的作用是物理内存使用完，之后将磁盘空间虚拟成内存来使用，启用swap设备会对系统的性能产生非常负面的影响，因此kubernetes要求每个节点都要禁用swap设备，但是如果因为某些原因确实不能关闭swap分区，就需要在集群安装过程中通过明确的参数进行配置说明

```powershell
# 编辑分区配置文件/etc/fstab，注释掉swap分区一行
# 注意修改完毕之后需要重启linux服务
vim /etc/fstab
注释掉 /dev/mapper/centos-swap swap
```

##### 2.6.7 修改linux的内核参数

```powershell
# 修改linux的内核采纳数，添加网桥过滤和地址转发功能
# vim /etc/sysctl.d/kubernetes.conf文件，添加如下配置
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1


# 重新加载配置
[root@master ~]# sysctl -p
# 加载网桥过滤模块
[root@master ~]# modprobe br_netfilter
# 查看网桥过滤模块是否加载成功 
[root@master ~]# lsmod | grep br_netfilter
```

##### 2.6.8 配置ipvs功能

在Kubernetes中Service有两种代理模型，一种是基于iptables的，一种是基于ipvs的两者比较的话，ipvs的性能明显要高一些，但是如果要使用它，需要手动载入ipvs模块。

```powershell
# 1.安装ipset和ipvsadm
[root@master ~]# yum install ipset ipvsadm -y

# 2.添加需要加载的模块写入脚本文件
[root@master ~]# cat <<EOF> /etc/sysconfig/modules/ipvs.modules
#!/bin/bash
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack_ipv4
EOF

# 3.为脚本添加执行权限
[root@master ~]# chmod +x /etc/sysconfig/modules/ipvs.modules

# 4.执行脚本文件
[root@master ~]# /bin/bash /etc/sysconfig/modules/ipvs.modules

# 5.查看对应的模块是否加载成功
[root@master ~]# lsmod | grep -e ip_vs -e nf_conntrack_ipv4
```

**重启服务器**

##### 2.6.9 安装docker

```powershell
# 1、切换镜像源 默认源是国外的
[root@master ~]# wget https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo -O /etc/yum.repos.d/docker-ce.repo

# 2、查看当前镜像源中支持的docker版本
[root@master ~]# yum list docker-ce --showduplicates

# 3、安装特定版本的docker-ce
# 必须指定--setopt=obsoletes=0，否则yum会自动安装更高版本
[root@master ~]# yum install --setopt=obsoletes=0 docker-ce-18.06.3.ce-3.el7 -y

# 4、添加一个配置文件
# Docker在默认情况下使用Vgroup Driver为cgroupfs，而Kubernetes推荐使用systemd来替代cgroupfs
[root@master ~]# mkdir /etc/docker
[root@master ~]# cat <<EOF> /etc/docker/daemon.json
{
	"exec-opts": ["native.cgroupdriver=systemd"],
	"registry-mirrors": ["https://kn0t2bca.mirror.aliyuncs.com"]
}
EOF

# 5、启动dokcer 设置开机自启动
[root@master ~]# systemctl restart docker
[root@master ~]# systemctl enable docker 
```

##### 2.6.10 安装Kubernetes组件

```powershell
# 1、由于kubernetes的镜像在国外，速度比较慢，这里切换成国内的镜像源
# 2、vim /etc/yum.repos.d/kubernetes.repo,添加下面的配置
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgchech=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
			http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg

# 3、安装kubeadm、kubelet和kubectl
[root@master ~]# yum install --setopt=obsoletes=0 kubeadm-1.17.4-0 kubelet-1.17.4-0 kubectl-1.17.4-0 -y

# 4、配置kubelet的cgroup
# vim /etc/sysconfig/kubelet, 添加下面的配置
KUBELET_CGROUP_ARGS="--cgroup-driver=systemd"
KUBE_PROXY_MODE="ipvs"

# 5、设置kubelet开机自启
[root@master ~]# systemctl enable kubelet
```

##### 2.6.11 准备集群镜像

```powershell
# 在安装kubernetes集群之前，必须要提前准备好集群需要的镜像，所需镜像可以通过下面命令查看
[root@master ~]# kubeadm config images list

# 下载镜像
# 此镜像kubernetes的仓库中，由于网络原因，无法连接，下面提供了一种替换方案
images=(
	kube-apiserver:v1.17.4
	kube-controller-manager:v1.17.4
	kube-scheduler:v1.17.4
	kube-proxy:v1.17.4
	pause:3.1
	etcd:3.4.3-0
	coredns:1.6.5
)

for imageName in ${images[@]};do
	docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
	docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
	docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName 
done
```

##### 2.6.11 集群初始化

>下面的操作只需要在master节点上执行即可

```powershell
# 创建集群
[root@master ~]# kubeadm init \
	--apiserver-advertise-address=192.168.5.3 \
	--image-repository registry.aliyuncs.com/google_containers \
	--kubernetes-version=v1.17.4 \
	--service-cidr=10.96.0.0/12 \
	--pod-network-cidr=10.244.0.0/16
	
# 创建必要文件
[root@master ~]# mkdir -p $HOME/.kube
[root@master ~]# sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
[root@master ~]# sudo chown $(id -u):$(id -g) $HOME/.kube/config
# 展示所有node节点
[root@master ~]# kubectl get node
```

> 下面的操作只需要在node节点上执行即可

```powershell
kubeadm join 192.168.5.3:6443 --token awk15p.t6bamck54w69u4s8 \
    --discovery-token-ca-cert-hash sha256:a94fa09562466d32d29523ab6cff122186f1127599fa4dcd5fa0152694f17117 
```

在master上查看节点信息 ，由于未进行网络安装，所以是NotReady

```powershell
[root@master ~]# kubectl get nodes
NAME    STATUS   ROLES     AGE   VERSION
master  NotReady  master   6m    v1.17.4
node1   NotReady   <none>  22s   v1.17.4
node2   NotReady   <none>  19s   v1.17.4
```

##### 2.6.13 安装网络插件flannel，只在master节点操作即可

插件使用的是DaemonSet的控制器，他会在每个节点上都运行

```powershell
wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

# 启动fannel
kubectl apply -f kube-flannel.yml
```

由于外网不好访问，如果出现无法访问的情况，可以直接用下面的 记得文件名是kube-flannel.yml，位置：/root/kube-flannel.yml内容：

```powershell
https://github.com/flannel-io/flannel/tree/master/Documentation/kube-flannel.yml
```

##### 2.6.14 使用kubeadm reset重置集群

```shell
# 在master节点之外的节点进行操作
kubeadm reset
systemctl stop kubelet
systemctl stop docker
rm -rf /var/lib/cni/
rm -rf /var/lib/kubelet/*
rm -rf /etc/cni/
ifconfig cni0 down
ifconfig flannel.1 down
ifconfig docker0 down
ip link delete cni0
ip link delete flannel.1
# 重启kubelet
systemctl restart kubelet
# 重启docker
systemctl restart docker
```

##### 2.6.15 重启kubelet和docker

```powershell
# 重启kubelet
systemctl restart kubelet
# 重启docker
systemctl restart docker
```

使用配置文件启动fannel

```powershell
kubectl apply -f kube-flannel.yml
```

等待它安装完毕 发现已经是 集群的状态已经是Ready

##### 2.6.16 kubeadm中的命令

```powershell
# 生成 新的token
[root@master ~]# kubeadm token create --print-join-command
```

#### 2.7 集群测试 master操作

##### 2.7.1 创建一个nginx服务

```powershell
kubectl create deployment nginx  --image=nginx:1.14-alpine
```

##### 2.7.2 暴露端口

```powershell
kubectl expose deployment nginx  --port=80 --target-port=80  --type=NodePort
```

##### 2.7.3 查看服务

```powershell
kubectl get pod,svc
```

##### 2.7.4 查看pod ：后面的是外界访问的端口，随机分配

![img](images/2232696-20210621233130477-111035427.png)

浏览器测试结果：

![img](images/2232696-20210621233157075-1117518703.png)











### 3. 资源管理

#### 3.1 资源管理介绍

在kubernetes中，所有的内容都抽象为资源(pod pod控制器 service)，用户需要通过操作资源来管理kubernetes。

> kubernetes的本质上就是一个集群系统，用户可以在集群中部署各种服务，所谓的部署服务，其实就是在kubernetes集群中运行一个个的容器，并将指定的程序跑在容器中。
>
> kubernetes的最小管理单元是pod而不是容器，所以只能将容器放在`Pod`中，而kubernetes一般也不会直接管理Pod，而是通过`Pod控制器`来管理Pod的。
>
> Pod可以提供服务之后，就要考虑如何访问Pod中服务，kubernetes提供了`Service`资源实现这个功能。
>
> 当然，如果Pod中程序的数据需要持久化，kubernetes还提供了各种`存储`系统。

![image-20200406225334627](Kubenetes.assets/image-20200406225334627.png)

> 学习kubernetes的核心，就是学习如何对集群上的`Pod、Pod控制器、Service、存储Volume` 等各种资源进行操作

#### 3.2 YAML语言介绍

YAML是一个类似 XML、JSON 的标记性语言。它强调以**数据**为中心，而不是以标识语言为重点。因而YAML本身的定义比较简单，号称"一种人性化的数据格式语言"。

```XML
<heima>
    <age>15</age>
    <address>Beijing</address>
</heima>
```

```YAML
heima:
  age: 15
  address: Beijing
```

YAML的语法比较简单，主要有下面几个：

> - 大小写敏感
> - 使用缩进表示层级关系
> - 缩进不允许使用tab，只允许空格( 低版本限制，高版本可以)
> - 缩进的空格数不重要，只要相同层级的元素左对齐即可
> - '#'表示注释
>

YAML支持以下几种数据类型：

> - 纯量：单个的、不可再分的值
> - 对象：键值对的集合，又称为映射（mapping）/ 哈希（hash） / 字典（dictionary）
> - 数组：一组按次序排列的值，又称为序列（sequence） / 列表（list） 
>

```yml
# 纯量, 就是指的一个简单的值，字符串、布尔值、整数、浮点数、Null、时间、日期
# 1 布尔类型
c1: true (或者True)
# 2 整型
c2: 234
# 3 浮点型
c3: 3.14
# 4 null类型 使用~表示null
c4: ~  
# 5 日期类型 日期必须使用ISO 8601格式，即yyyy-MM-dd
c5: 2018-02-17
# 6 时间类型 时间使用ISO 8601格式，时间和日期之间使用T连接，最后使用+时区
c6: 2018-02-17T15:02:31+08:00
# 7 字符串类型 简单写法，直接写值, 如果字符串中间有特殊字符，必须使用双引号或者单引号包裹 
c7: heima 
# 字符串过多的情况可以拆成多行，每一行会被转化成一个空格
c8: 'line1
    line2'
```

```yaml
# 对象
# 形式一(推荐):
heima:
  age: 15
  address: Beijing
# 形式二(了解):
heima: {age: 15,address: Beijing}
```

```yaml
# 数组
# 形式一(推荐):
address:
  - 顺义
  - 昌平  
# 形式二(了解):
address: [顺义,昌平]
```

> 小提示：
>
> 1 书写yaml切记`:` 后面要加一个空格再写值
>
> 2 如果需要将多段yaml配置放在一个文件中，中间要使用`---`分隔
>
> ```yaml
> heima: good
> ---
> heima1: very good
> ```
>
> 3 下面是一个yaml转json的网站，可以通过它验证yaml是否书写正确
>
> https://www.json2yaml.com/convert-yaml-to-json



#### 3.3 资源管理方式

> - 命令式对象管理：直接使用命令去操作kubernetes资源，以创建pod为例
>
>   ``` powershell
>   kubectl run nginx-pod --image=nginx:1.17.1 --port=80
>   ```
>
> - 命令式对象配置：通过命令配置和配置文件去操作kubernetes资源，create/patch 创建 更新 改查 后面跟参数。
>
>   ```powershell
>   kubectl create/patch -f nginx-pod.yaml
>   ```
>
> - 声明式对象配置：通过apply命令和配置文件去操作kubernetes资源。只用于创建和更新资源。
>
>   ```powershell
>   kubectl apply -f nginx-pod.yaml
>   ```
>   
>

| 类型           | 操作对象 | 适用环境 | 优点                                     | 缺点                                           |
| :------------- | :------- | :------- | :--------------------------------------- | :--------------------------------------------- |
| 命令式对象管理 | 对象     | 测试     | 简单                                     | 只能操作活动对象，无法审计、跟踪，用于临时查询 |
| 命令式对象配置 | 文件     | 开发     | 可以审计、跟踪                           | 项目大时，配置文件多，操作麻烦                 |
| 声明式对象配置 | 目录     | 开发     | 支持目录操作，可以对其下所有文件进行操作 | 意外情况下难以调试。                           |

##### 3.3.1 命令式对象管理

**kubectl命令**

> kubectl是kubernetes集群的命令行工具，通过它能够对集群本身进行管理，并能够在集群上进行容器化应用的安装部署。kubectl命令的语法如下：
>
> ```
> kubectl [command] [type] [name] [flags]
> ```
>
> **comand**：指定要对资源执行的操作，例如create、get、delete
>
> **type**：指定资源类型，比如deployment程序、pod、service
>
> **name**：指定资源的名称，名称大小写敏感
>
> **flags**：指定额外的可选参数

```shell
# 查看所有pod
kubectl get pod 

# 查看某个pod
kubectl get pod pod_name

# 查看某个pod,以yaml格式展示结果
kubectl get pod pod_name -o yaml
```

**操作**

kubernetes允许对资源进行多种操作，可以通过--help查看详细的操作命令

```
kubectl --help
```

经常使用的操作有下面这些：

| 命令分类       | 命令         | 翻译                        | 命令作用                                       |
| :------------- | :----------- | :-------------------------- | :--------------------------------------------- |
| **基本命令**   | create       | 创建                        | 创建一个资源                                   |
|                | edit         | 编辑                        | 编辑一个资源                                   |
|                | get          | 获取                        | 获取一个资源                                   |
|                | patch        | 更新                        | 更新一个资源                                   |
|                | delete       | 删除                        | 删除一个资源                                   |
|                | explain      | 解释                        | 展示资源文档                                   |
| **运行和调试** | run          | 运行                        | 在集群中运行一个指定的镜像                     |
|                | expose       | 暴露                        | 暴露资源为Service                              |
|                | describe     | 描述                        | 显示资源内部信息                               |
|                | logs         | 日志输出容器在 pod 中的日志 | 输出容器在 pod 中的日志                        |
|                | attach       | 缠绕进入运行中的容器        | 进入运行中的容器                               |
|                | exec         | 执行容器中的一个命令        | 执行容器中的一个命令                           |
|                | cp           | 复制                        | 在Pod内外复制文件                              |
|                | rollout      | 首次展示                    | 管理资源的发布                                 |
|                | scale        | 规模                        | 扩(缩)容Pod的数量                              |
|                | autoscale    | 自动调整                    | 自动调整Pod的数量                              |
| **高级命令**   | apply        | rc                          | 通过文件对资源进行配置                         |
|                | label        | 标签                        | 更新资源上的标签                               |
| **其他命令**   | cluster-info | 集群信息                    | 显示集群信息`kubectl cluster-info`             |
|                | version      | 版本                        | 显示当前Server和Client的版本 `kubectl version` |

**资源类型**

kubernetes中所有的内容都抽象为资源，资源类型可以通过下面的命令进行查看:

```
kubectl api-resources
```

经常使用的资源有下面这些：

| 资源分类          | 资源名称                 | 缩写   | 资源作用                        |
| :---------------- | :----------------------- | :----- | :------------------------------ |
| **集群级别资源**  | nodes                    | no     | 集群组成部分`kubectl get nodes` |
|                   | namespaces               | ns     | 隔离Pod                         |
| **pod资源**       | pods                     | po     | 装载容器                        |
| **pod资源控制器** | replicationcontrollers   | rc     | 控制pod资源                     |
|                   | replicasets              | rs     | 控制pod资源                     |
|                   | deployments              | deploy | 控制pod资源                     |
|                   | daemonsets               | ds     | 控制pod资源                     |
|                   | jobs                     |        | 控制pod资源                     |
|                   | cronjobs                 | cj     | 控制pod资源                     |
|                   | horizontalpodautoscalers | hpa    | 控制pod资源                     |
|                   | statefulsets             | sts    | 控制pod资源                     |
| **服务发现资源**  | services                 | svc    | 统一pod对外接口                 |
|                   | ingress                  | ing    | 统一pod对外接口                 |
| **存储资源**      | volumeattachments        |        | 存储                            |
|                   | persistentvolumes        | pv     | 存储                            |
|                   | persistentvolumeclaims   | pvc    | 存储                            |
| **配置资源**      | configmaps               | cm     | 配置                            |
|                   | secrets                  |        | 配置                            |

下面以一个namespace / pod的创建和删除简单演示下命令的使用：

```shell
# 创建一个namespace 缩写也可以 # kubectl create ns dev
[root@master ~]# kubectl create namespace dev
namespace/dev created

# 获取namespace
[root@master ~]# kubectl get ns
NAME              STATUS   AGE
default           Active   21h
dev               Active   21s
kube-node-lease   Active   21h
kube-public       Active   21h
kube-system       Active   21h

# 在此namespace dev下创建并运行一个nginx的Pod
[root@master ~]# kubectl run pod --image=nginx:latest -n dev
kubectl run --generator=deployment/apps.v1 is DEPRECATED and will be removed in a future version. Use kubectl run --generator=run-pod/v1 or kubectl create instead.
deployment.apps/pod created

# 查看新创建的pod # kubectl describe pods pod -n dev
[root@master ~]# kubectl get pods -n dev
NAME  READY   STATUS    RESTARTS   AGE
pod   1/1     Running   0          21s

# 删除指定的pod
[root@master ~]# kubectl delete pods pod -n dev
pod "pod" deleted

# 删除指定的namespace
[root@master ~]# kubectl delete ns dev
namespace "dev" deleted
```

##### 3.3.2 命令式对象配置

命令式对象配置就是使用命令配合配置文件一起来操作kubernetes资源。

1） 创建一个**nginxpod.yaml**，内容如下：

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: dev
# namespace的声明
---
# 创建pod
apiVersion: v1
kind: Pod
metadata:
  name: nginxpod
  namespace: dev
spec:
  containers:
  - name: nginx-containers
    image: nginx:latest
```

2）执行create命令，创建资源：

```powershell
[root@master ~]# kubectl create -f nginxpod.yaml
namespace/dev created
pod/nginxpod created
```

此时发现创建了两个资源对象，分别是namespace和pod

3）执行get命令，查看资源：

```shell
[root@master ~]#  kubectl get -f nginxpod.yaml
NAME            STATUS   AGE
namespace/dev   Active   18s

NAME            READY   STATUS    RESTARTS   AGE
pod/nginxpod    1/1     Running   0          17s
```

这样就显示了两个资源对象的信息

4）执行delete命令，删除资源：

```shell
[root@master ~]# kubectl delete -f nginxpod.yaml
namespace "dev" deleted
pod "nginxpod" deleted
```

此时发现两个资源对象被删除了

```shell
总结:
    命令式对象配置的方式操作资源，可以简单的认为：命令  +  yaml配置文件（里面是命令需要的各种参数）
```

##### 3.3.3 声明式对象配置

声明式对象配置跟命令式对象配置很相似，但是它只有一个命令apply。

```shell
# 首先执行一次kubectl apply -f yaml文件，发现创建了资源
[root@master ~]#  kubectl apply -f nginxpod.yaml
namespace/dev created
pod/nginxpod created

# 再次执行一次kubectl apply -f yaml文件，试图更新，发现资源没有变动
[root@master ~]#  kubectl apply -f nginxpod.yaml
namespace/dev unchanged
pod/nginxpod unchanged

# 若发生改变 可以通过以下命令查看
kubectl describe pods nginxpod -n dev
```

```powershell
总结:
    其实声明式对象配置就是使用apply描述一个资源最终的状态（在yaml中定义状态）
    使用apply操作资源：
        如果资源不存在，就创建，相当于 kubectl create
        如果资源已存在，就更新，相当于 kubectl patch
```

> 扩展：kubectl可以在node节点上运行吗 ?

kubectl的运行是需要进行配置的，它的配置文件是$HOME/.kube，如果想要在node节点运行此命令，需要将master上的.kube文件复制到node节点上，即在master节点上执行下面操作：

```shell
scp  -r  ~/.kube   node1: ~/
```

> 使用推荐: 三种方式应该怎么用 ?

> 创建/更新资源 使用声明式对象配置 `kubectl apply -f XXX.yaml`
>
> 删除资源 使用命令式对象配置 `kubectl delete -f XXX.yaml`
>
> 查询资源 使用命令式对象管理 `kubectl get(describe) 资源名称`
>

