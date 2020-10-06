# Kubeadm

[github doc design](https://github.com/kubernetes/kubeadm/blob/master/docs/design/design_v1.10.md)

[kubeadm 1.18 init a cluster](https://www.jianshu.com/p/bba23a9fdee1)

[upstream installation guide](https://kubernetes.io/zh/docs/reference/setup-tools/kubeadm/kubeadm-init/)

[install deb by aliyun repo](https://www.cnblogs.com/xiaochina/p/11650520.html)

[upstream kubelet config for system](https://github.com/kubernetes/release/blob/master/cmd/kubepkg/templates/latest/deb/kubeadm/10-kubeadm.conf)

[kubelet flags](https://www.jianshu.com/p/36ad3028a710)

[cri-tools](https://github.com/kubernetes-sigs/cri-tools)

[kubernetes cni plugins](https://github.com/containernetworking/plugins)

[install](https://blog.csdn.net/qq_27374315/article/details/88720507)

[kubeadm create cluster](https://www.cnblogs.com/rainingnight/p/using-kubeadm-to-create-a-cluster.html)

## prequisite

1. 域名修改
2. 关闭 swap selinux iptables 等
3. 内核参数调整
   开启内核 ipv4 转发需要加载 br_netfilter 模块
   ```
   #每台节点
   modprobe br_netfilter
   modprobe ip_conntrack
   ```
   bridge-nf 使得netfilter可以对Linux网桥上的 IPv4/ARP/IPv6 包过滤。比如，设置net.bridge.bridge-nf-call-iptables＝1后，二层的网桥在转发包时也会被 iptables的 FORWARD 规则所过滤
4. ipvs
   从 k8s 1.8 开始，kube-proxy 引入了IPVS模式，IPVS模式与 iptables 同样基于 netfilter，但是采用的 hash 表，因此当 service 数量达到一定规模时，hash 查表的速度优势就会显现出来，从而提高 service 的服务能力
   ```
   cat > /etc/sysconfig/modules/ipvs.modules <<EOF
    #!/bin/bash
    modprobe -- ip_vs
    modprobe -- ip_vs_rr
    modprobe -- ip_vs_wrr
    modprobe -- ip_vs_sh
    modprobe -- nf_conntrack
    EOF

    chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack

    #查看是否已经正确加载所需的内核模块
   ```
5. ipset
   iptables是Linux服务器上进行网络隔离的核心技术，内核在处理网络请求时会对iptables中的策略进行逐条解析，因此当策略较多时效率较低；而是用IPSet技术可以将策略中的五元组(协议，源地址，源端口,目的地址，目的端口)合并到有限的集合中，可以大大减少iptables策略条目从而提高效率。测试结果显示IPSet方式效率将比iptables提高100倍
6. ipvsadm
7. timezone
8. cni plugins

## Kubeadm Init Phase

```
kubeadm config images pull
--image-repository=registry.aliyuncs.com/google_containers
--kubernetes-version=v1.18.3
```

### Preflight Phase

```bash
systemctl enable docker.service

apt install -y ebtables

apt install -y ethtool

apt install -y socat

systemctl enable kubelet.service

apt install -y conntrack
```

### Kubelet Phase

### Cert Phase

`/etc/kubernetes/pki`

- Root CA
- ApiServer
- KubeletClient
// Front Proxy certs
- FrontProxyCA
- FrontProxyCAClient
// Etcd certs
- EtcdCA
- EtcdServer
- EtcdPeer
- EtcdHealthCheck
- EtcdAPIClient

### Config Phase

`/etc/kubernetes`

- admin.conf
- kubelet.conf
- controller-manager.conf
- scheduler.conf

#### admin.conf

- "admin" is defined as kubeadm itself and the actual person(s) that is administering the cluster and want to control the cluster:
   - with this full control (root) over the cluster
   - inside this file, a client certificate is generated from the `ca.crt` and `ca.key`. The client cert should:
      - Be a client certificate
      - Be in `system::master` organization
      - Include a CN, but that can be anything. `kubeadm` uses the `kubernetes-admin` CN.

#### kubelet.conf

- Inside this file, a client certificate is generated from the `ca.cert` and `ca.key`. The client cert should:
   - Be a client certificate
   - Be in the `system::node` organization
   - Have the CN `system::node:<hostname-lowercase>`

#### controller-manager.conf

- Inside this file, a client certificate is generated from the `ca.crt` and `ca.key`. The client cert should:
   - Be a client certificate
   - Have the CN `system:kube-controller-manager`

#### scheduler.conf

- Inside this file, a client certificate is generated from `ca.crt` and `ca.key`. The client cert should:
   - Be a client certificate
   - Have the CN `system:kube-scheduler`

Please note that:

1. `ca.cert` is embeded in all the kubeconfig file
2. If a given kubeconfig exists, and its content is evaluated compliant with the above specs, the existing file will be used and the generation phase for the given kubeconfig skipped.
3. If `admin` is running in ExternalCA mode, all the required kueconfig must be provided by the user as well, because `kubeadm` cannot generate any of them by itself.
4. In case of `kubeadm` executed in the `--dry-run` mode, kubeconfig files are written in a temporary folder.
5. Kubeconfig files generation can be invocked individually with the `kubeadm alpha phase kubeconfig all` command.


### Control-plane phase

**Generate Static Pod Manifest for Control Plane**

- All static Pods are deployed on `kube-system` namespace
- All static Pods get `tier: control-plane` and `component:{component-name}` labels
- All static Pods get `scheduler.alpha.kubernetes.io/critical-pod` annotation. Note. this will be moved over to the proper solution of using Pod Priority and Preemption when ready.
- `hostNetwork: true` is set on all static Pods to allow control plane startup before a network is configured; accordingly:
   - The `address` that the controller-manager and the scheduler use to refer the API server is `127.0.0.1`
   - If using a local etcd server, `etcd-servers` address will be set to `127.0.0.1:2379`
- Leader election is enabled for both the controller-manager and the scheduler.
- Controller-manager and the scheduler will reference kubeconfig files with their respective, unique identities.
- All static Pods get any extra flags specified by the user.
- All static Pods get any extra Volumes specified by the user(Host Path)

#### API server

#### Scheduler

### Etcd Phase

**Generate Static Pod Manifest for local etcd**

If the user specified an external etcd this step will be skipped, otherwise a static manifest file will be generated for creating a local etcd instance running in a Pod with following attributes:

- listen on `localhost:2379` and use `HostNetwork=true`
- make a `hostPath` mount out from the `dataDir` to the host's filesystem.
- Any extra flags specified by the user.

### Wait for Control Plane Phase

### Upload ConifgMap Phase

kubeadm saves the configuration passed to `kubeadm init`, either via flags or the config file, in a ConfigMap named `kubeadm-config` under `kube-system` namespace.

### Mark Control Plane Phase

As soon as the control plane is available, kubeadm executes following actions:

- Label the master with `node-role.kubernetes.io/master=""`
- Taints the master with `node-role.kubernetes.io/master:NoSchedule`

### Bootstrap Token Phase

kubeadm uses [Authenticating with Bootstrap Tokens](https://kubernetes.io/docs/reference/access-authn-authz/bootstrap-tokens/) for joining new nodes to an existing cluster;

`kubeadm init` ensures that everything is properly configured for this process, and this includes following steps as well as setting API server and controller flags

#### Create a bootstrap token

#### Allow joining nodes to call CSR API

#### Setup auto approval for new bootstrap tokens

#### Setup nodes certificate rotation with auto approval

#### Create the public `cluster-info` ConfigMap

### Addon Phase

https://kubernetes.io/docs/concepts/cluster-administration/addons/

- CoreDNS
- kube-proxy

## Apply Pod Network

### Flannel

```bash
wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

kubectl apply -f kube-flannel.yml

```bash
root@k8s-master:~/kubernetes_master# kct apply -f kube-flannel.yml
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds-amd64 created
```

### Dashboard

```
// get dashboard
wget https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.4/aio/deploy/recommended.yaml
```

kubernetes-dashboard.yml

## Kubeadm Join Phases

### Preflight Checks

`kubeadm` executes a set of preflight checks before starting the join, with the aim to verify preconditions and avoid common cluster startup problems.

Please note that:

1. `kubeadm join` preflight checks are basically a subnet `kubeadm init` preflight checks

2. Starting from 1.9, kubeadm provides better support for CRI-generic functionality; in that case, linux specfic controls are skipped or replaced by similar controls for crictl.

3. Starting from 1.9, kubeadm provides support for joining nodes running on windows, in that case, linux specific controls are skipped.

4. In any case the user can skip specific preflight checks(or eventually all preflight checks) with `--ignore-preflight-errors` option.

### Discovery cluster-info

There are 2 main schemes for discovery. The first is to use a shared token along with the IP address of the API server. The second is to provide a file(a subset of the standard kubeconfig file)

#### Shared token discovery

If `kubeadm join` is invoked with --discovery-token, token discovery is used; in this case the node basically retrieves the cluster CA certificates from the `cluster-info` ConfigMap in the `kube-public` namespace.

In order to prevent "man in the middle" attacks, several steps are taken:

- First, the CA certificate is retrieved via insecure connection (note: this is possible because `kubeadm ini` granted access to `cluster-info` users for `system::unauthenticated`)

- Then the CA certificate goes through following validate steps:
   - "Base validation", using the token ID against a JWT signature
   - "Pub key validatation", using provided `--discovery-token-ca-cert-has`. This value is available in the output of "kubeadm init" or can be calculated using standard tools(the hash is calculated over the bytes of the Subject Public Key Info(SPKI)). The `--discovery-token-ca-cert-hash flag` may be repeated multiple times to allow more than one public key.
   - as a additional validation, the CA certificate is retrieved via secure connection and then compared with the CA retrieved initially

Please note that:

"Pub key validation" can be skipped passing `--discovery-token-unsafe-skip-ca-verification flag`; This weakens the kubeadm security model since others can potentially impersonate the Kubernetes Master.

#### File/https discovery

If `kubeadm join` is invoked with `--discovery-file`, file discovery file is used; this file can be a local file or downloaded via an HTTPS URL; in case of HTTPS, the host installed CA bundle is used to verify the connection.

With file discovery, the cluster CA certificates is provided into the file itself; in fact, the discovery file is a kubeconfig file with only `server` and `certificate-authority-data` attributes set

## Kubeadm Reset Phase

```bash
# clean up node info
kubectl drain <node name> --delete-local-data --force --ignore-daemonsets

kubectl delete node <node name>

# then run reset command on node machine
kubeadm reset

# cni network
systemctl stop kubelet
systemctl stop docker
rm -rf /var/lib/cni/
rm -rf /var/lib/kubelet/*
rm -rf /etc/cni/
ifconfig cni0 down
ifconfig flannel.1 down
ifconfig docker0 down
ip link delete cni0
ip link delete flannel.1
##重启kubelet
systemctl restart kubelet
##重启docker
systemctl restart docker
```


## Q&A

[k8s pod ping clusterip](https://blog.51cto.com/13641616/2442005)

[pod-network-cidr](https://blog.csdn.net/shida_csdn/article/details/104334372)

[podip vs clusterip vs externalip](https://blog.csdn.net/xichenguan/article/details/79141445)

[flannel PodCIDR](https://ithelp.ithome.com.tw/articles/10222753)

[flannel overlay](https://blog.laputa.io/kubernetes-flannel-networking-6a1cb1f8ec7c)