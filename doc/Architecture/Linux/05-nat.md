# NAT

[nat](https://www.cnblogs.com/imstudy/p/5458133.html)

## IP
A类地址：8位 0xxxxxxx...xxxxxxxx
B类地址：16位 10xxxxxx...xxxxxxxx
C类地址：24位 110xxxxx...xxxxxxxx
D类地址：保留 1110xxxx...xxxxxxxx

内部地址
A类：10.xxx
B类：172.xxx
C类：192.168.xxx

## NAT
### 原理
- 网络分为公网和内网两部分，NAT网管设置在内网到外网的路由出口位置，双向流量必须要经过NAT网关
- 网络连接由内网发起
- NAT网关在两个访问方向上完成两次地址的转换和翻译，出方向做原信息替换，入方向做目的信息的替换
- NAT网关的存在对通信双方是保持透明的
- NAT网关为了实现双向翻译的功能，需要维护一张关联表，把会话信息保存下来