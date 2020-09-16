# Certificate

[k8s证书](https://zhuanlan.zhihu.com/p/123858320)

[certificates](https://kubernetes.io/zh/docs/concepts/cluster-administration/certificates/)

[k8s certificate](https://www.cnblogs.com/centos-python/articles/11043570.html)

[kubernetes tls bootstrapping](https://mritd.me/2018/01/07/kubernetes-tls-bootstrapping-note/)

[trusted on MAC](https://blog.csdn.net/qq_40460909/article/details/85682595)

## Catalog

```
[root@k8s-master01 ~]# cd /etc/kubernetes/pki
[root@k8s-master01 pki]# tree
.
├── apiserver.crt
├── apiserver-etcd-client.crt
├── apiserver-etcd-client.key
├── apiserver.key
├── apiserver-kubelet-client.crt
├── apiserver-kubelet-client.key
├── ca.crt
├── ca.key
├── etcd
│   ├── ca.crt
│   ├── ca.key
│   ├── healthcheck-client.crt
│   ├── healthcheck-client.key
│   ├── peer.crt
│   ├── peer.key
│   ├── server.crt
│   └── server.key
├── front-proxy-ca.crt
├── front-proxy-ca.key
├── front-proxy-client.crt
├── front-proxy-client.key
├── sa.key
└── sa.pub
```

## 公钥、私钥和证书

**公钥和私钥是成对的，他们互相解密。
公钥加密，私钥解密。
私钥数字签名，公钥验证。**

### 根证书与证书

通常我们配置 https 服务时需要到 “权威机构” (CA)申请证书。过程是这样的：
1. 网站创建一个秘钥对，提供公钥和组织以及个人信息给权威机构
2. 权威机构颁发证书
3. 浏览网页的朋友利用权威机构的根证书公钥解密签名，对比摘要，确定合法性
4. 客户端验证域名有效信息时间等（浏览器基本都内置各大权威机构的CA公钥）

证书包括如下：
1. 申请者公钥
2. 申请者组织和个人组织
3. 签发机构CA信息，有效时间，序列号等
4. 以上信息的签名

根证书又称自签名证书，也就是自己给自己颁发的证书。 CA 被称为证书授权中心，k8s中的ca证书就是根证书


## 客户端选择证书的原理

1. 证书选择是在客户端和服务端 SSL/TLS 握手协商阶段商定的
2. 服务端如果要求客户端提供证书，则在握手时会向客户端发送一个它接受的 CA 列表
3. 客户端查找它的证书列表(一般是操作系统的证书，对于 MAC 为 keychain)，看有没有被CA签名的证书，如果有，则将它们提供给用户选择(证书的私钥)；
4. 用户选择一个证书私钥，然后客户端将使用它和服务端通信