# KubeConfig

[kubeconfig usage](https://www.jianshu.com/p/99853cac56b8)

[Organize Cluster Access Using Kubeconfig Files](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)

## kubeconfig 配置文件的生成

an example:

```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: xxx
    server: https://xxx:6443
  name: cluster1
- cluster:
    certificate-authority-data: xxx
    server: https://xxx:6443
  name: cluster2
contexts:
- context:
    cluster: cluster1
    user: kubelet
  name: cluster1-context
- context:
    cluster: cluster2
    user: kubelet
  name: cluster2-context
current-context: cluster1-context
kind: Config
preferences: {}
users:
- name: kubelet
  user:
    client-certificate-data: xxx
    client-key-data: xxx
```

apiVersion 和 kind 标识客户端解析器的版本和模式，不应该手动编辑

### cluster 模块

cluster 中包含 Kubernetes 集群的 Endpoint 数据，包括 kubernetes apiserver 的完整 url 以及集群的证书颁发机构。

可以使用 kubectl config set-config 添加或修改 cluster 条目。

### users 模块

user 定义用于向 Kubernetes 集群进行身份验证的客户度凭证。

可用凭证有 

- client-certificate
- client-key
- token
- username/passwd

username/passwd 和 token 是二者只能选择一个， 但 client-certificate 和 client-key 可以分别与他们组合。

可以使用 kubectl config set-credentials 添加或者修改 user 条目。

### context 模块

context 定义了一个命名的 `cluster/user/namespace` 元组，用于使用提供的认证信息和命名空间将请求发送到指定的集群。

三个都是可选的，仅使用 `cluster/user/namespace` 之一指定上下文，或指定 `none`。

未指定的值或在加载的 kubeconfig 中没有相应条目的命名值将被换为默认值。

可以使用 kubectl config set-context 添加或修改 context。

### current-context 模块

current-context 是作为 `cluster/user/namespace` 元组的 key， 当 kubectl 从该文件中加载配置的时候会被默认使用。

## 使用 kubeconfig 文件配置 kubectl 跨集群认证

kubectl 作为操作 k8s 的一个客户端工具，只要为 kubectl 提供连接 apiserver 的配置(kubeconfig)，kubectl 可以在任何地方操作该集群，当然，若 kubeconfig 文件中配置多个集群，kubectl 也可以轻松的在多个集群之间切换。

