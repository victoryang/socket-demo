# DaemonSet

[DaemonSet Introduction](https://www.cnblogs.com/xzkzzz/p/9553321.html)

A *DaemonSet* ensures that all(or some) Nodes run a copy of a Pod. As nodes are added to the cluster, Pods are added to them. As nodes are removed from the cluster, those Pods are garbage collected. Deleting a DaemonSet will clean up the Pods it creates.

Some typical uses of a DaemonSet are:

- running a cluster storage daemon on every node
- running a log collection daemon on every node
- running a node monitoring daemon on every node

## Writing a DaemonSet Spec

### Creating a DaemonSet

## How Daemon Pods are scheduled

## Communicating with DaemonSet Pods

- **Push**
- **NodeIP and Known Port**
- **DNS**
- **Service**

## Updating a DaemonSet

