# Network

[Netstat](https://www.cnblogs.com/qianyuliang/p/10542713.html)

[Linux Kernel Network](https://blog.csdn.net/yudelian/article/details/89332638)

[Linux Network Article Catalog](https://blog.csdn.net/yudelian/article/details/89332638)

[tcp connect timeout](https://www.cnblogs.com/lanyangsh/p/10152734.html)

[tcp timewait](https://yq.aliyun.com/articles/581106)

[tcpdump抓包](https://www.jianshu.com/p/1371809155a2)


## TCP

### Tcp 连接超时问题

近期出现容器部署服务中断的问题，具体表现为对远程机器执行restful请求时，概率性出现连接超时的问题。观察log可以发现tcp的连接超时时间在3s。代码中设置的tcp的连接超时为30s，不符合预期，然后检查内核参数发现，tcp连接相关的参数分别为：

```
net.ipv4.tcp_syn_retries = 1
net.ipv4.tcp_synack_retries = 3
net.ipv4.tcp_syncookies = 1
```

其中，tcp_syn_retries设置为1，代表tcp连接重连时，syn包的重发的次数为1，对应的总超时时间为3s

主机从发出数据包到第一次TCP重传开始，这段时间被称为retransmission timeout，即RTO。

## Tcpdump

tcpdump -iany tcp port [port]