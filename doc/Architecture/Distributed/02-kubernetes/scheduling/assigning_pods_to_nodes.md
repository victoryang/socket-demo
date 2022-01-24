# Assigning Pods to Nodes

You can constrain a Pod so that it can only run on particular set of Nodes. There are serval ways to do this and the recommended approaches all use label selectors to facilitate the selection. Generally such constraints are unnecessary, as the scheduler will automatically do a reasonable placement but there are some circumstances where you may want to control which node the pod deploys to - for example to ensure that a pod ends up on a machine with an SSD attached to it, or to co-locate pods from different services that communicate a lot into the same availability zone.

## nodeSelector

`nodeSelector` is the simplest recommended form of node selection constraint. `nodeSelector` is a field of PodSpec. It specifies a map of key-value pairs. For the pod to be eligible to run on a node, the node must have each of the indicated key-value pairs as labels(it can have additional labels as well). The most common usage is one key-value pair.

## Interclude
[built-in node labels](https://kubernetes.io/docs/reference/labels-annotations-taints/)

## Node isolation/restriction

## Affinity and anti-affinity

`nodeSelector` provides a very simple way to constrain pods to nodes with particular labels. The affinity/anti-affinity feature, greatly expands the types of constraints you can express. The key enhancements are

1. The affinity/anti-affinity language is more expensive. The language offers more matching rules besides exact matches created with a logical AND operation;
2. you can indicate that the rule is "soft"/"preference" rather than a hard requirement, so if the scheduler can't satisfy it, the pod will still be scheduled
3. you can constrain against labels on other pods running on the node, rather than against labels on the node itself, which allows rules about which pods can and cannot be co-located