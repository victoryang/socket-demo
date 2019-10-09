# Docker

[docker and iptables](https://docs.docker.com/network/iptables/)
[container networking model](https://success.docker.com/article/networking)

## Docker
### Version
- docker ce
- docker ee

## Dockerfile

## Docker data

## Docker network
- iptables on linux server
- routing rules on windows server
- forms and encapsulates packet
- handle encryption

## Docker compose
    使用docker-compose.yml 模板文件来定义一组相关联的应用容器为一个项目

## Docker machine
    Docker machine 是官方编排工具之一， 负责在多种平台上快速安装Docker环境

## Docker swarm
    Docker swarm 是官方提供的Docker容器集群服务

    具有以下新特性：
    - 去中心化设计
    - 内置服务发现
    - 负载均衡
    - 路由网格
    - 动态伸缩
    - 滚动更新
    - 安全传输

### 结点
    运行docker的主机可以主动初始化一个swarm集群，或者加入一个已存在的swarm集群，这样主机就成为了一个swarm的结点

    结点分为manager结点和worker结点
    manager结点用于swarm结点的管理，多个结点通过raft协议选出leader
    worker结点是任务执行结点，manager结点将服务内容下发至工作结点执行，并且manager结点默认也作为worker结点

### 服务和任务
    task是swarm中最小的调度单位，目前来说就是一个单一的容器
    service是指一组任务的集合，服务定义了任务的属性

### 创建swarm集群
1 创建manager结点

    /* 在远程机器上创建虚拟机环境的docker*/
    docker-machine create -d virtualbox manager

    /* 连接到远程机器*/
    docker-machine ssh manager

    /* 创建swarm集群*/
    docker swarm init --advertise-addr 192.168.99.100

2 增加工作结点

    /* 在远程机器上创建虚拟机环境的docker*/
    docker-machine create -d virtualbox worker1

    /* 连接到远程机器*/
    docker-machine ssh worker1

    /* 连接到刚才创建的集群中*/
    docker swarm joint --token xxx-xxx-xxx 192.168.99.100:2377

3 集群操作

    /* 查看集群*/
    docker node ls

4 服务管理
### 1. docker service

    使用docker service 命令来管理swarm集群，该命令只能在manager结点运行

    /* 在创建的集群中运行一个名为nginx的服务*/
    docker service create --replicas 3 -p 80:80 --name nginx nginx:latest

    /* 查看服务*/
    docker service ls

    /* 查看服务详情*/
    docker service ps

    /* 服务伸缩*/
    docker service scale nginx=5

### 2. docker stack

    结合docker-compose.yml来一次配置、启动多个服务

    /* 根据docker-compose部署多服务到集群中*/
    docker stack deploy -c docker-compose.yml wordpress

    /* 查看服务*/
    docker stack ls

### 3. docker secret

    集群秘钥、证书管理服务

### 4. docker config

    集群配置文件管理服务

### 5. docker service update

    滚动发行
