# Kubeadm

[github doc design](https://github.com/kubernetes/kubeadm/blob/master/docs/design/design_v1.10.md)

[kubeadm 1.18](https://www.jianshu.com/p/bba23a9fdee1)

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

## Phase

kubeadm init
```
kubeadm config images pull
--image-repository=registry.aliyuncs.com/google_containers
--kubernetes-version=v1.18.3
```

### Preflight Phase

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

