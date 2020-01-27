# Salt

[salt](https://www.cnblogs.com/wudonghang/p/5871957.html)

## Feature

- 部署简单
- 支持大部分Linux和Windows
- 主从集中化
    - master监听4505和4506两个端口
    - 4505 为消息发布端口
    - 4506 为客户端和服务端通信的端口
- 配置简单，功能性强，扩展性强
- Master与Minion基于证书认证，安全可靠
    - Minion在第一次启动时，会在/etc/salt/pki/minion下自动生成证书(私钥)和公钥，然后发送公钥给master
    - master在接收到Minion的公钥后，通过salt-key接收，并且存到master的/etc/salt/pki/master/minions下以minio id命名的目录下
    - 完成认证过程
- 支持API及自定义模块，扩展轻松