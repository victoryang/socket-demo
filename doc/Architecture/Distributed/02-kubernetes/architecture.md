# Architecture

[Arch](http://docs.kubernetes.org.cn/251.html)

## Architecture

### Core Component

- etcd 保存整个集群的状态
- apiserver 提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制
- controller manager 负责维护集群的状态，比如故障检测、自动扩展、滚动更新等
- scheduler 负责资源的调度，按照预定的调度策略将 Pod 调度到响应的机器上
- kubelete 负责维护容器的生命周期，同时也负责 Volume (CVI)和 network (CNI)的管理
- container runtime 负责镜像管理以及 Pod 和容器的真正运行 (CRI)
- kube-proxy 负责为 Service 提供 cluster 内部的服务发现和负载均衡

### Addons

- kube-dns 负责为整个集群提供DNS服务
- ingress controller 为服务提供外网入口
- Heapster 提供资源监控
- Dashboard 提供UI
- Federation 提供跨可用区的集群
- Fluent-elastic 提供集群日志采集、存储和查询

## Hierarchy

- 核心层：对外提供API构建高层的应用，对内提供插件式应用执行环境
- 应用层：部署(无状态应用、有状态应用、批处理任务、集群应用等) 和路由 (服务发现、DNS解析等)
- 管理层：系统度量(如基础设施、容器和网路的度量)，自动化(如自动扩展、动态 Provision等) 以及策略管理(RBAC/Quota/PSP/Network Policy等)
- 接口层：kubectl commandline，客户端 SDK 以及集群联邦
- 生态层：
    - 外部： 日志、监控、配置管理、CICD、Workflow、FaaS等
    - 内部： CRI、CNI、CVI、镜像仓库等


