# Ingress

[Ingress](https://segmentfault.com/a/1190000019908991)

[kubernetes addons](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/#additional-controllers)

[kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx)

[nginx/kubernetes-ingress](https://github.com/nginxinc/kubernetes-ingress)

[aliyun version ingress](https://help.aliyun.com/document_detail/151524.html?spm=a2c4g.11186623.6.1098.7a4546e30LJG3J)

## Service

Service 的作用体现在两个方面，对集群内部，不断跟踪 pod 的变化，更新 endpoint 中对应 pod 的对象，提供了 ip 不断变化的 pod 的服务发现机制。

对集群外部，类似负载均衡器，可以在集群内外部对 pod 进行访问。但单独用 service 暴露服务的方式，存在如下问题：

- ClusterIP 的方式只能在集群内部访问
- NodePort 方式下，端口管理比较麻烦
- LoadBalance 方式受限于云平台，需要在云平台部署额外的 LB 实例

## Ingress

### ingress vs ingress-controller

- ingress
    - 指 k8s 中的一个 api 对象，一般用 yaml 配置。用来定义请求如何转发到 service 的规则，可以理解为配置模板
- ingress-controller
    - 具体实现反向代理及负载均衡的程序，对 ingress 定义的规则进行解析，根据配置的规则来实现请求转发

ingress-controller 是负责具体转发的组件，通过各种方式将它暴露在集群入口，外部对集群的请求流量会先到 ingress-controller，而 ingress 对象是用来告诉 ingress-controller 该如何转发请求，比如哪些域名哪些 path 要转发到哪些服务等。

### ingress-controller

ingress-controller 并不是 k8s 自带的组件，实际上 ingress-controller 只是一个统称，用户可以选择不同的 ingress-controller 实现，目前 k8s 维护的 ingress-controller 只有 google 云的 GCE 与 ingress-nginx 两个，其他还有很多第三方维护的 ingress-controller。无论哪种 ingress-controller，实现的机制都大同小异，只是在配置上有差异。一般来说，ingress-controller 的形式都是一个 pod，里面跑着 daemon 程序和反向代理程序。daemon 负责不断监控集群的变化，根据 ingress 对象生成配置并应用新配置到反向代理，比如 nginx-ingress 就是动态生成 nginx 配置，动态更新 upstream，并在需要的时候 reload 程序应用新配置。

### ingress

ingress 是一个 API 对象，和其他对象一样，通过 yaml 文件来配置。ingress 通过 http 或 https 暴露集群内部 service，给 service 提供外部 URL、负载均衡、SSL/TLS 能力以及基于 host 的方向代理。ingress 要依靠 ingress-controller 来具体实现以上功能。

在 ingress 的配置中，annotations 很重要。前面说 ingress-controller 有很多不同的实现，而不同的 ingress-controller 就可以根据 "kubernetes.io/ingress.class:" 来判断要使用哪些 ingress 的配置，同时，不同的 ingress-controller 也有对应的 annotations 配置，用于自定义一些参数。例如 'nginx.ingress.kubernetes.io/use-regex:"true"'，最终是在生成 nginx 配置中，采用 location~ 来表示正则匹配。

### ingress 的部署

#### Deployment + LoadBalancer Service

用 Deployment 部署 ingress-controller，创建一个 type 为 LoadBalancer 的 service 关联这组 Pod。大部分公有云，都会为 LoadBalancer 的 service 自动创建一个负载均衡器，通常还绑定了公网地址。只要把域名解析指向该地址，就实现了集群服务的对外暴露。

#### Deployment + NodePort Service

同样用 Deployment 模式部署 ingress-controller，并创建对应的服务，但是 type 为 NodePort。这样，ingress 就会暴露在集群节点 ip 的特定端口上。由于 NodePort 暴露的端口是随机端口，一般会在前面再搭建一套负载均衡器来转发请求。该方式一般用于宿主机是相对固定的环境 ip 地址不变的场景。

NodePort 方式暴露 ingress 虽然简单方便，但是 NodePort 多了一层 NAT，在请求量级很大时可能对性能会有一定影响。

#### DaemonSet + HostNetwork + NodeSelector

[DaemonSet + HostNetwork + NodeSelector](https://segmentfault.com/a/1190000019908991)

用 DaemonSet 结合 NodeSelector 来部署 ingress-controller 到特定的 node 上，然后使用 HostNetwork 直接把该 pod 与宿主机 node 的网络打通，直接使用宿主机的 80/443 端口就能访问服务。这时，ingress-controller 所在的 node 机器就很类似传统架构的边缘节点，比如机房入口的 nginx 服务器。该方式整个请求链路最简单，性能相对 NodePort 模式更好。缺点是由于直接利用宿主机节点的网络和端口，一个 node 只能部署一个 ingress-controller pod。比较适合大并发的生产环境使用。

