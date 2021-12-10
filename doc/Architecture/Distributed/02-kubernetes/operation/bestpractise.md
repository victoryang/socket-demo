# Best Practise

## Consideration for large clusters

A cluster is a set of nodes(physical or virtual machines) running Kubernetes agents, managed by the control plane. Kubernetes v1.22 supports clusters with up to 5000 nodes. More specifically, Kubernetes is designed to accommodate configuration that meet *all* of the following criteria:

- No more than 110 pods per node
- No more than 5000 nodes
- No more than 150000 total pods
- No more than 300000 total containers

You can scale your cluster by adding or removing nodes. The way you do this depends on how your cluster is deployed.

### Cloud provider resource quotas

To avoid running into cloud provider quota issues, when creating a cluster with many nodes, consider:

- Requesting a quota increase for cloud resources such as:
    - Computer instance
    - CPUs
    - Storage volumes
    - In-use IP addresses
    - Packet filters rule sets
    - Number of load balancers
    - Network subnets
    - Log streams
- Gating the cluster scaling actions to bring up new nodes in batches, with a pause between batches, because some cloud provides rate limit the creation of new instances.

### Control plane components

For a large cluster, you need a control plane with sufficient compute and other resources.

Typically you would run one or two control plane instances per failure zone, scaling those instances vertically first and then scaling horizontally after reaching the point of falling returns scale.

You should run at least one instance per failure zone to provide fault-tolerance. Kubernetes nodes do not automatically steer traffic towards control-plane endpoints that are in the same failure zone; however, your cloud provider might have its own mechanisms to do this.

For example, using a managed load balancer, you configure the load balancer to send traffic that originates from the kubelet and Pods in failure zone *A*, and direct that traffic only to the control plane hosts that are also in zone *A*. If a single control-plane host or endpoint failure zone *A* goes offline, that means that all the control-plane traffic for nodes in zone *A* is now being sent between zones. Running multiple control plane hosts in each zone makes that outcome less likely.

#### etcd storage

To improve performance of large clusters, you can store Event objects in a separate dedicated etcd instance.
When creating a cluster, you can(using custom tooling):
- start and configure additional etcd instance
- configure the API server to use it for storing events

See Operating etcd clusters for Kubernetes and Set up a High Availability etcd cluster with kubeadm for details on configuring and managing etcd for a large cluster.

### Addon resources

Kubernetes resource limits help to minimize the impact of memory leaks and other ways that pods and containers can impact on other components. These resource limits apply to addon resources just as they apply to application workloads

## Running in multiple zones

### Background

Kubernetes is designed so that a single Kubernetes cluster can run across multiple failure zones, typically where these zones fit within a logical grouping called a region. Major cloud providers define a region as a set of failure zones(also called availability zones) that provide a consistent set of features: within a region, each zone offers the same APIs and services.

Typically cloud architecture aim to minimize the chance that a failure in one zone also impairs services in another zone.