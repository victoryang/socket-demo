# Production

[k8s in production](https://segmentfault.com/a/1190000005345466)

[production-environment](https://kubernetes.io/zh/docs/setup/production-environment/)

## Production environment

A production-quality Kubernetes cluster requires planning and preparation. If your Kubernetes cluster is to run critical workloads, it must be configured to be resilient. This page explains steps you can take to set up a production-ready cluster, or to promote an existing cluster for production use.

## Production considerations

Typically, a production Kubernetes cluster environment has more requirements than a personal learning, development, or test environment Kubernetes. A production environment may require secure access by many users, consistent availablility, and the resources to adapt to changin demands.

As you decide where you want your production Kubernetes environment to live(on premises or in a cloud) and the amount of management you want to take on or hand to others, consider how your requirements for a Kubernetes cluster are influenced by the following issues:

- Availability: A single-machine Kubernetes learning environment has a single point of failure. Creating a highly available cluster means considering:
    - Separating the control plane from the worker nodes
    - Replicating the control plane components on multiple nodes
    - Load balancing traffic to the cluster's API server
    - Having enough worker nodes available, or able to quickly become available, as changing workloads warrant it

- Scale: If you expect your production Kubernetes environment to receive a stable amount of demand, you might be able to set up for the capacity you need and be done. However, if you expect demand to grow over time or change dramatically based on things like season or special events, you need to plan how to scale to relieve increased pressure from more requests to the control plane and worker nodes or scale down to reduce unused resources.

- Security and access management: You have full admin privileges on your own Kubernetes learning cluster. But shared clusters with important workloads, and more than one or two users, requirea more refined approach to who and what can access cluster resources. You can use role-based access control and other security mechanisms to make sure that users and workloads can get access to the resources they need, while keeping workloads, and the cluster itself, secure. You can set limits on the resources that users and workloads can access by managing policies and container resources.

Before building a Kubernetes production environment on your own, consider handing off some or all of this job to Turnkey Cloud Solutions providers or other Kubernetes Partners. Options include:

- Serveless: Just run workloads on third-party equipment without managing a cluster at all. You will be charged for things like CPU usage, memory, and disk requests.
- Managed control plane: Let the provider manage the scale and availability of the cluster's control plane, as well as handle patches and upgrades.
- Managed worker nodes: Configure pools of nodes to meet your needs, then the provider makes sure those nodes are available and ready to implement upgrades when needed.
- Integration: There are providers that integrate Kubernetes with other services you may need, such as storage, container registries, authentication methods, and development tools.

## Production cluster setup

In a production-quality Kubernetes cluster, the control plane manages the cluster from services that can be spread across multiple computers in different ways. Each worker node, however, represents a single entity that is configured to run Kubernetes pods.

### Production control plane

The simplest Kubernetes cluster has the entire control plane and worker node services running on the same machine. You can grow that environment by adding worker nodes, as reflected in the diagram illustrated in Kubernetes Components. If the cluster is meant to be available for a short period of time, or can be discarded if something goes seriously wrong,this might meet your needs.

If you need a more permanet, highly available cluster, however, you should consider ways of extending the control plane. By design, one-machine control plane services running on a single machine are not highly available. If keeping the cluster up and running and ensuring that it can be repaired if something goes wrong is important, consider these steps:

- Choose deployment tools: You can deploy a control plane using tools such as kubeadm, kops, and kubespray. See Installing Kubernetes with deployment toos to learn tips for production-quality deployments using each of those deployment methods. Different Container Runtimes are available to use with your deployments.
- Manage certificates: Secure communications between control plane services are implemented using certificates. Certificates are automatically generated during deployment or you can generate them using your own certificate authority.
- Configure load balancer for apiserver. Configure a load balancer to distirbute external API requests to the apiserver service instances running on different nodes
- Separate and backup etcd service: The etcd services can either run on the same machines as other control plane services or run on seperate machines, for extra security and availability. Because etcd stores cluster configuration data, backing up the etcd database should be done regularly to ensure that you can repair that database if needed.
- Create multiple control plane systems: Fo high availability, the control plane should not be limited to a single machine. If the control plane services are run by an init service(such as systemd), each service should run on at three machines. However, running control plane services as pods in Kubernetes ensures that the replicated number of services that you request will always be available. The scheduler should be fault toerant, but not highly available. Some deployment toos set up Raft consensus algorithm to do leader election for Kubernetes services. If the primary goes away, another service elects itself and take over. 
- Span multiple zones: If keeping your cluster available at all times is critical, consider creating a cluster that runs across multiple data centers, referred to as zones in cloud environments. Groups of zones are referred to as regions. By spreading a cluster across multiple zones in the same region, it can improve the chances that your cluster will continue to function even if one zone becomes unavailable.
- Manage on-going features: If you plan to keep your cluster over time, there are tasks you need to do to maintain its health and security. For example, if you installed with kubeadm, there are instructions to help you with Certificate Management and Upgrading kubeadm clusters.

### Production worker nodes

Production worker nodes

Production-quality workloads need to be resilient and anything they rely on needs to be resilient(such as CoreDNS). Whether you manage your own control plane or have a cloud provider do it for you, you still need to consider how you want to manage your worker nodes(also referred to simply as nodes)

- Configure nodes: Nodes can be physical or virtual machines. If you want to create and manage your own nodes, you can install a supported operating system, then add and run the appropriate Node services. Consider:
    - The demands of your workloads when you set up nodes by having appropriate memory, CPU, and disk speed and storage capacity available
    - Whether generic computer systems will do or you have workloads that need GPU processors, Windows nodes, or VM isolation
- Validate nodes: how to ensure that a node meets the requirements to join Kubernetes cluster
- Add nodes to the cluster: if you are managing your own cluster you can add nodes by setting up your own machines and either adding them manually or having them register themselves to the cluster's apiserver.
- Add Windows nodes to the cluster: Kubernetes offers support for Windows worker nodes, allowing you to run workloads implemented in Windows containers
- Scale nodes: Have a plan for expanding the capacity your cluster will eventually need. If you are managing nodes yourself, this can mean purchasing and installing your own physical equipment.
- Autoscale nodes: Most cloud providers support Cluster Autoscaler to replace unhealthy nodes or grow and shrink the number of nodes as demand requires. See the virtualization platforms that can be scripted to spin up new nodes based on demand.
- Setup node health checks: For important workloads, you want to make sure that nodes and pods runnig on those nodes are healthy. Using the Node Problem Detector daemon, you can ensure your nodes are healthy.

### Production user management

In production, you may be moving from a model where you or a small group of people are accessing the cluster to where may potentially be dozens or hundreds of people. In a learning environment or platform prototype, you might have a single administrative account for everything you do. In production, you will want more accounts with different levels of access to different namespaces.

Taking on a production-quality cluster means deciding how you want to selectively allow access by other users. In particular, you need to select strategies for validating the identities of those who try to access your cluster(authentication) and deciding if they have permissions to do what they are asking(authorization):

- Authentication: The apiserver can authenticate users using client certificate, bearer tokens, an authenticating proxy, or HTTP basic auth. You can choose which authentication methods you want to use. Using plugins, the apiserver can leverage your organization's existing authentication methods, such as LDAP or kerberos. See Authentication for a description of these different methods of authenticating Kubernetes users.

- Authorization: When you set out to authorize your regular users, you will probably choose between RBAC and ABAC authorization. See Authorization Overview to review different modes for authorizing user accounts(as well as service account access to your cluster):
    - Role-based access control(RBAC): Lets you assign access to your cluster by allowing specific sets of permissions to authenticated users. Permissions can be assigned for a specific namespace(Role) or across the entire cluster(ClusterRole). The using RoleBindings and ClusterRoleBindings, those permissions can be attached to particular users.
    - Attribute-based access control(ABAC): Lets you create policies based on resource attributes in the cluster and will allow or deny access based on those attributes. Each line of a policy file identifies versioning properties(apiVersion and kind) and a map of spec properties to match the subject (user or group), resource property, non-resource property(/version and /apis), and readonly.

As some setting up authentication and authorization on your production Kubernetes cluster, here are some things to consider:

- Set the authorization mode: When the Kubernetes API server(kube-apiserver) starts, the supported authentication modes must be set using the --authorization-mode flag. For example, that flag in the kube-adminserver.yaml file (in /etc/kubernetes/manifests) could be set to Node, RBAC. This would allow Node and RBAC authorization for authenticated requests.
- Create user certificates and role bindings (RBAC): If you are using RBAC authorization, users can create a CertificateRequest(CSR) that can be signed by the cluster CA. Then you can bind Roles and ClusterRoles to each user.
- Create policies that combine attributes(ABAC): If you are using ABAC authorization, you can assign combinations of attributes to form policies to authorize selected users or groups to access particular resources(such as pod), namespace, or apiGroup.
- Consider Admission Controllers: Additional forms of authorization for requests that can come in through the API server include Webhook Token Authentication. Webhooks and other special authorization types need to be enabled by adding Admission Controllers to the API server.

### Set limits on workload resources

Demands from production workloads can cause pressure both inside and outside of the Kubernetes control plane. Consider these items when setting up for the needs of your cluster's workloads:

- Set namespace limits: Set per-namespace quotas on things like memory and CPU. You can also set Hierachical Namespaces for inheriting limits.
- Prepare for DNS demand: If you expect workloads to massively scale up, your DNS service must be ready to scale up as well
- Create additional service accounts: User accounts determine what users can do on a cluster, while a service account defines pod access within a particular namespace. By default, a pod takes on the default service account from its namespace. See Managing Service Accounts for information on creating a new service account. For example, you might want to:
    - Add secrets that a pod could use to pull images from a particular container registry.
    -  Assign RBAC permission to a service account.