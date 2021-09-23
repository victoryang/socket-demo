# CNI

[CNI Compare](https://www.hwchiu.com/cni-compare.html)

[Flannel](https://www.kubernetes.org.cn/2270.html)

[Flannel Modes](https://www.cnblogs.com/sandshell/p/11769642.html)

[libcni](https://github.com/containernetworking/cni)
[plugins](https://github.com/containernetworking/plugins)

[cni plugin issue](https://blog.csdn.net/github_35614077/article/details/98213572)

## Basic CNI

[Basic plugins](https://github.com/containernetworking/plugins/tree/master/plugins/main)

- bridge
- host-device
- ipvlan
- macvlan
- ptp
- vlan
- loopback

## Flannel

Flannel 是将多个不同子网(基于node主机) 通过被 Flannel 维护的 Overlay 网络拼接成为一张大网来实现互联的。

flannel 的这个 overlay 网络支持多种后端实现，除UDP外，还有 VxLAN 和 host-gw 等。此外，flannel 支持通过两种模式来维护隧道端点上FDB的信息，其中一种是通过连接 Etcd 来实现，另外一种是直接对接 k8s，通过 k8s 添加删除 Node 来触发更新。

