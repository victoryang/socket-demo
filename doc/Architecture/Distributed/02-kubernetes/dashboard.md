# Dashboard

[access control](https://github.com/kubernetes/dashboard/blob/master/docs/user/accessing-dashboard/README.md)

***Important***
[Right Way to Access Dashboard-ui](https://www.cnblogs.com/rainingnight/p/deploying-k8s-dashboard-ui.html)

[github](https://github.com/kubernetes/dashboard#kubernetes-dashboard)

[docs](https://github.com/kubernetes/dashboard/tree/master/docs)

[blog share](https://blog.csdn.net/networken/article/details/85607593)

## Generate Certificate

https://github.com/kubernetes/dashboard/blob/master/docs/user/certificate-management.md

### Public trusted Certificate Authority

[Let's encrypt](https://letsencrypt.org/getting-started/)

### Self Signed Cert

OpenSSL

```bash
openssl genrsa -des3 -passout pass:over4chars -out dashboard.pass.key 2048
...
openssl rsa -passin pass:over4chars -in dashboard.pass.key -out dashboard.key
# Writing RSA key
rm dashboard.pass.key
openssl req -new -key dashboard.key -out dashboard.csr

openssl x509 -req -sha256 -days 365 -in dashboard.csr -signkey dashboard.key -out dashboard.crt
```

## Ways to Access Kubernetes-dashboard

https://github.com/kubernetes/dashboard/blob/master/docs/user/accessing-dashboard/README.md

### Proxy

Access Locally Only

```bash
$ kubectl proxy
Starting to serve on 127.0.0.1:8001

kubectl proxy --address='0.0.0.0'  --accept-hosts='^*$'
```

dashboard 只允许 localhost 使用 http 访问，其他地址使用 https

### NodePort

```
spec:
  clusterIP: 10.103.5.139
  ports:
  - port: 443
    protocol: TCP
    targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard
  sessionAffinity: None
  type: NodePort
```

通过 NodePort 开放端口给外部，但由于开启了 RBAC，故需要生成证书来访问

https://github.com/kubernetes/dashboard/tree/master/docs/user

### API Server

#### 将 k8s 集群的根证书加入到本地，设置为 always trusted
如 https://blog.csdn.net/qq_40460909/article/details/85682595


#### 使用 /etc/kubernetes/admin.conf 生成客户端访问 api-server 的证书

```bash
# 生成client-certificate-data
grep 'client-certificate-data' ~/.kube/config | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.crt

# 生成client-key-data
grep 'client-key-data' ~/.kube/config | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.key

# 生成p12
openssl pkcs12 -export -clcerts -inkey kubecfg.key -in kubecfg.crt -out kubecfg.p12 -name "kubernetes-client"
```

#### 创建 Admin User

recommand.yml 所创建的User `kubernetes-dashboard` 只有 namespace `kubernetes-dashboard` 的权限，无法做到全局的管理，根据官方提供的方式：

https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md

可以创建一个 admin 的 user 来获取全局的信息。

### URL

https://<master-ip>:<apiserver-port>/api/v1/namespaces/<kubernetes-dashboard-namespace>/services/https:kubernetes-dashboard:/proxy/

### Ingress