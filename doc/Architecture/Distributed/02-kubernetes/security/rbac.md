# RBAC

[role & rolebinding & clusterrole & clusterrolebinding](https://blog.csdn.net/hxpjava1/article/details/103779148)

[k8s.io doc rbac](https://kubernetes.io/zh/docs/reference/access-authn-authz/rbac/#%E5%88%9D%E5%A7%8B%E5%8C%96%E4%B8%8E%E9%A2%84%E9%98%B2%E6%9D%83%E9%99%90%E5%8D%87%E7%BA%A7)

[k8s rbac](https://www.cnblogs.com/weiyiming007/p/10484763.html)

[k8s rbac](https://note.abeffect.com/articles/2019/08/26/1566807667654.html)

## Access Control

- ABAC
- RBAC
- Webhook
- Node
- AlwaysDeny
- AlwaysAllow

### RBAC

Role-Based access control is a method of regulating access to computer or network resources based on the roles of individual users within your organization

RBAC authorization uses the `rbac.authorization.k8s.io` API group to drive authorization decisions, allowing you to dynamically configure policies through the Kubernetes API

To enable RBAC, start the API Server with the `--authorization-mode` flag set to a comma-seperated list that includes `RBAC`

### API Object

The RBAC API declares four kinds of Kubernetes object: *role*,*ClusterRole*,*RoleBinding* and *ClusterRoleBinding*.

#### Role and ClusterRole

An RBAC *Role* or *ClusterRole* contains rules that represent a set of permissions. Permissions are purely additive(there are no "deny" rules).

A Role always sets permissions within a particular namespace; when you create a Role,you have to specify the namespace it belongs in.

ClusterRole, by contrast, is a non-namespaced resource. The resources have different names(Role and ClusterRole) because a Kubernetes object always has to be either namespaced or not namespaced; it can't be both.

ClusterRoles have several uses. You can use a ClusterRole to:
1. define permissions on namespaced resources and are granted within individual namespaces
2. define permissions on namespaced resources and be granted across all namespaces
3. define permissions on cluster-scoped resources

If you want to define a role within a namespace, use a Role;if you want to define a role cluster-wide, use a ClusterRole.

#### RoleBinding and ClusterRoleBinding

A role binding grants the permissions defined in a role to a user or set of users. It holds a list of *subject*(users, groups or service accounts), and a reference to the role being granted. A RoleBinding grants permissions within a specific namespace whereas a ClusterRoleBinding grants that access cluster-wide.

A RoleBinding may reference any Role in the same namespace. Alternatively, a RoleBinding can reference a ClusterRole and bind that ClusterRole to the namespace of the RoleBinding. If you want to bind a ClusterRole to all the namespaces in your cluster, you use a ClusterRoleBindng. 