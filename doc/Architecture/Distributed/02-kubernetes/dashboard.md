# Dashboard

[access control](https://github.com/kubernetes/dashboard/blob/master/docs/user/accessing-dashboard/README.md)

[ways to access dashboard-ui](https://www.cnblogs.com/rainingnight/p/deploying-k8s-dashboard-ui.html)

[github](https://github.com/kubernetes/dashboard#kubernetes-dashboard)

[docs](https://github.com/kubernetes/dashboard/tree/master/docs)

[blog share](https://blog.csdn.net/networken/article/details/85607593)

## generate cert files

### Self Signed Cert

https://github.com/kubernetes/dashboard/blob/master/docs/user/certificate-management.md

```bash
# 生成client-certificate-data
grep 'client-certificate-data' ~/.kube/config | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.crt

# 生成client-key-data
grep 'client-key-data' ~/.kube/config | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.key

# 生成p12
openssl pkcs12 -export -clcerts -inkey kubecfg.key -in kubecfg.crt -out kubecfg.p12 -name "kubernetes-client"
```

## Ways to Access Kubernetes-dashboard

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

```bash
# 生成client-certificate-data
grep 'client-certificate-data' ~/.kube/config | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.crt

# 生成client-key-data
grep 'client-key-data' ~/.kube/config | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.key

# 生成p12
openssl pkcs12 -export -clcerts -inkey kubecfg.key -in kubecfg.crt -out kubecfg.p12 -name "kubernetes-client"
```

### Ingress