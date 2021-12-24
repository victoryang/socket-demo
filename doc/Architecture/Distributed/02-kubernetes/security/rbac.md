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

A RoleBinding can also reference a ClusterRole to grant the permissions defined in that ClusterRole to resources inside the RoleBinding's namespace. This kind of reference lets you define a set of common roles across your cluster, then reuse them within multiple namespaces.

```
apiVersion: rbac.authorization.k8s.io/v1
# This role binding allows "dave" to read secrets in the "development" namespace.
# You need to already have a ClusterRole named "secret-reader".
kind: RoleBinding
metadata:
  name: read-secrets
  #
  # The namespace of the RoleBinding determines where the permissions are granted.
  # This only grants permissions within the "development" namespace.
  namespace: development
subjects:
- kind: User
  name: dave # Name is case sensitive
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: secret-reader
  apiGroup: rbac.authorization.k8s.io
```

For instance, even though the following RoleBinding refers to a ClusterRole, "dave"(the subject, case sensitive) will only be able to read Secrets in the "development" namespace, because the roleBinding's namespace(in its metadata) is "development".

#### ClusterRoleBinding example

To grant permissions across a whole cluster, you can use a ClusterRoleBinding. The following ClusterRoleBinding allows any user in the group "manager" to read secrets in any namespace.

```
apiVersion: rbac.authorization.k8s.io/v1
# This cluster role binding allows anyone in the "manager" group to read secrets in any namespace.
kind: ClusterRoleBinding
metadata:
  name: read-secrets-global
subjects:
- kind: Group
  name: manager # Name is case sensitive
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: secret-reader
  apiGroup: rbac.authorization.k8s.io
```

After you create a binding, you cannot change the Role or ClusterRole that it refers to. If you try to change a binding's `roleref`, you get a validation error. If you do want to change the `roleref` for a binding, you need to remove the binding object and create a replacement.

There are two reasons for this restriction:

1. Making `roleref` immutable allows granting someone `update` permission on an existing binding object, so that they can manage the list of subjects, without being able to change the role that is granted to those subjects.

2. A binding to a different role is a fundamentally different binding. Requiring a binding to be deleted/recreated in order to change the `roleRef`

The `kubectl auth reconcile` command-line utility creates or updates a manifest file containing RBAC objects, and handles deleting and recreating binding objects if required to change the role they refer to.

#### Referring to resources

In the Kubernetes API, most resources are represented and accessed using a string representation of their object name, such as `pods` for a Pod. RBAC refers to resrouces using exactly the same name that appears in the URL for the relevant API endpoint. Some Kubernetes APIs involve a *subresource*, such as the logs for a Pod. A request for a Pod's logs looks like:

```
GET /api/v1/namespaces/{namespace}/pods/{name}/log
```

In this case, `pods` is the namespaced resource for Pod resources, and `log` is subresource of `pods`. To represent this in an RBAC role, use a slash(`/`) to delimit the resource and subresource. To allow a subject to read `pods` and also access the `log` subresource for each of those Pods, you write:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: pod-and-pod-logs-reader
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list"]
```

You can also refer to resources by name for certain requests through the `resourceNames` list. When specified, requests can be restricted to individual instances of a resource. Here is an example that restricts its subject to only `get` or `update` a ConfigMap named `my-configmap`:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: configmap-updater
rules:
- apiGroups: [""]
  #
  # at the HTTP level, the name of the resource for accessing ConfigMap
  # objects is "configmaps"
  resources: ["configmaps"]
  resourceNames: ["my-configmap"]
  verbs: ["update", "get"]
```

#### Aggregated ClusterRoles

You can *aggregate* serveral ClusterRoles into one combined ClusterRole. A controller, running as part of the cluster control plane, watches for ClusterRole objects with an `aggregationRule` set. The `aggregationRule` defines a label selector that the controller uses to match other ClusterRole objects that should be combined into the `rules` field of this one.

Here is an example aggreated ClusterRole:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: monitoring
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.example.com/aggregate-to-monitoring: "true"
rules: [] # The control plane automatically fills in the rules
```

If you create a new ClusterRole that matches the label selector of an existing aggregated ClusterRole, that change triggers adding the new rules into the aggregated ClusterRole. Here is an example that adds rules to the "monitoring" ClusterRole, by creating another ClusterRole labeled `rbac.example.com/aggregate-to-monitoring: true`.

```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: monitoring-endpoints
  labels:
    rbac.example.com/aggregate-to-monitoring: "true"
# When you create the "monitoring-endpoints" ClusterRole,
# the rules below will be added to the "monitoring" ClusterRole.
rules:
- apiGroups: [""]
  resources: ["services", "endpoints", "pods"]
  verbs: ["get", "list", "watch"]
```

The default user-facing roles use ClusterRole aggregation. This lets you, as a cluster administrator, include rules for custom resources, such as those served by CustomResourceDefinitions or aggregated API servers, to extend the default roles.

For example: the following ClusterRoles let the "admin" and "edit" default roles manage the custom resource named CronTab, whereas the "view" role can perform only read actions on CronTab resources. You can assume that CronTa objects are named "cronTabs" in URLs as seen by the API server.

#### Referring to subjects

A RoleBinding or ClusterRoleBinding binds a role to subjects. Subjects can be groups, users or ServiceAccounts.

Kubernetes represents usernames as strings. These can be: plain names, such as "alice";email-style names, like "bob@example"; or numeric user IDs represented as a string. It is up to you as a cluster administrator to configure the authentication modules so that authentication produces usernames in the format you want.

> **Caution:** The prefix `system:`is reserved for Kubernetes system use, so you should ensure that you don't have users or groups with names that start with `system:` by accident. Other than this special prefix, the RBAC authorization system does not require any format for usernames.

In Kubernetes, Authenticator modules provide group information. Groups, like users, are represented as string, and that string has no format requirements, other than that the prefix `system:` is reserved.

ServiceAccounts have names prefixed with `system:serviceaccount:`, and belong to groups that have names prefixed with `system:serviceaccounts:`.

> **Note:**
> - `system:serviceaccount:` is the prefix for service account usernames.
> - `system:serviceaccounts:` is the prefix for service account groups.

### Default roles and role bindings

API servers create a set of default ClusterRole and ClusterRoleBinding objects. Many of these are `system:` prefixed, which indicates that the resouce is directly managed by the cluster control plane. All of the default ClusterRoles and ClusterRoleBindings are labeled with `kubernetes.io/bootstrapping=rbac-defaults`.

> **Caution:** Take care when modifying ClusterRoles and ClusterRoleBinding with names that have a `system:` prefix. Modifications to these resources can result in non-functional clusters.

#### Auto-reconciliation

At each start-up, the API server updates default cluster roles with any missing permissions, and updates default cluster role bindings with any missing subjects. This allows the cluster to repair accidental modifications, and helps to keep roles and role bindings up-to-date as permissions and subjects change in new Kubernetes releases.

To opt out of this reconciliation, set the `rbac.authorization.kubernetes.io/auptupdate`, annotation on a default cluster role or rolebinding to `false`. Be aware that missing default permissions and subjects can result in non-functional clusters.

Auto-reconciliation is enabled by default if the RBAC authorizer is active.

### API discovery roles

Default role binding authorize unauthenticated and authenticated users to read API information that is deemed safe to be publicly accessible (including CustomResourceDefinitions). To disable anonymous unauthenticated access, add `--anonymous-auth=false` to the API server configuration.

To view the configuration of these roles via `kubectl` run:

```
kubectl get clusterroles system:discovery -o yaml
```

>**Note:** If you edit that ClusterRole, your changes will be overwritten on API server restart via auto-reconciliation. To avoid that overwriting, either do not manually edit the role, or disable auto-reconciliation.

||||
|-|-|-|
|Default ClusterRole|Default ClusterRoleBinding|Description|
|system:basic-user|system:authenticated group|Allows a user read-only access to basic information about themselves. Prior to v1.14, this role was also bound to `system:unauthenticated` by default.|
|system:discovery|system:authenticated group|Allows read-only access to API discovery endpoints needed to discover and negotiate API level. Prior to v1.14, this role was also bound to `system:unauthenticated` by default.|
|system:public-info-viewer|system:authenticated and system:unauthenticated groups|Allows read-only access to non-sensitive information about the cluster. Introduced in Kubernetes v1.14|

#### User-facing roles

Some of the default ClusterRoles are not `system:` prefixed. These are intended to be user-facing roles. They include super-user roles (cluster-admin), roles intended to be granted cluster-wide using ClusterRoleBindings, and roles intended to be granted within particular namespaces using RoleBindings(`admin`,`edit`,`view`).

User-facing ClusterRoles use ClusterRole aggregation to allow admins to include rules for custom resources on these ClusterRoles. To add rules to the `admin`,`edit`, or `view` roles, create a ClusterRole with one or more of the following labels:

```yaml
metadata:
    labels:
        rbac.authorization.k8s.io/aggregate-to-admin: "true"
        rbac.authorization.k8s.io/aggregate-to-edit: "true"
        rbac.authorization.k8s.io/aggregate-to-view: "true"
```

|Default ClusterRole|Default ClusterRoleBinding|Description|
|-|-|-|
|**cluster-admin**|system:masters group|Allows super-user access to perform any action on any resource. When used in a **ClusterRoleBinding**, it gives full control over every resource in the cluster and in all namespaces. When used in a **RoleBinding**, it gives full control over every resource in the role binding's namespace, including the namespace itself.|
|**admin**|None|Allows admin access,intended to be granted within a namespace using a **RoleBinding**.<br/> If used in a **RoleBinding**, allows read/write access to most resources in a namespace, including the ability to create roles and role bindings within the namespace. This role does not allow write access to resource quota or to the namespace itself. This role also does not allow write access to Endpoints in clusters created using Kubernetes v1.22+.|
|**edit**|None|Allow read/write access to most objects in a namespace. <br/>This role does not allow viewing or modifying roles or role bindings. However, this role allows accessing Secrets and running Pods as any ServiceAccount in the namespace, so it can be used to gain the API access levels of any ServiceAccount in the namespace. This role also does not allow write access to Endpoints in cluster created using Kubernetes v1.22+.|
|**view**|None|Allows read-only access to see most objects in a namespace. It does not allow viewing roles or role bindings.<br/> This role does not allow viewing Secrets, since reading the contents of Secrets enables access to ServiceAccount credentials in the namespace, which would allow API access as any ServiceAccount in the namespace(a form of privilege escalation)|

#### Core compnent roles

|Default ClusterRole|Default ClusterRoleBinding|Description|
|-|-|-|
|**system:kube-scheduler**|**system:kube-scheduler** user|Allow access to the resources required by the scheduler component|
|**system:volume-scheduler**|**system:volume-scheduler** user|Allows access to the volume resources required by the kube-scheduler component|
|**system:kube-controller-manager**|**system:kube-controller-manager** user|Allow access to resources required by the kubelet, **including read access to all secrets, and write access to all pod status objects**.<br/> You should use the Node authorizer and NodeRestriction admission plugin instead of the system:node role, and allow granting API access to kubelets based on the Pods scheduled to run on them. The system:node role only exists for compatibility with Kubernetes clusters upgraded from versions prior to v1.8.|
|**system:node-proxier**|**system:kube-proxy** user|Allow access to the resources required by the kube-proxy component|

#### Other component roles

#### Roles for built-in controllers

The Kubernetes controller manager runs controllers that are built in to the Kubernetes control plane. When invoke with `--use-service-account-credentials`, kube-controller-manager starts each controller using a separate service account. Corresponding roles exist for each built-in controller, prefixed with `system:controller`. If the controller manager is not started with `--user-service-account-credentials`, it runs all control loops using its own credential, which must be granted all the relevant roles.

### Privilege escalation prvention and bootstrapping

The RBAC API prevents users from escalating privileges by editing roles or role bindings. Because this is enforced at the API level, it applies even when the RBAC authorizer is not in use.

#### Restrictions on role creation or update

You can only create/update a role if at least one of the following things is true:

1. You already have allthe permissions contained in the role, at the same scope as the object being modified(cluster-wide for a ClusterRole, within the same namespace or cluster-wide for a Role)
2. You are granted explicit permission to perform the `escalate` verb on the `roles` or `clusterroles` resources in the `rbac.authorization.k8s.io` API group.

#### Restrictions on role binding creation or update

You can only create/update a role binding if you already have all the permissions contained in the referenced role(at the same scope as the role binding) or if you have been authorized to perform the `bind` verb on the refernced role. For example, if `user-1` does not have the ability to lit Secrets cluster-wide, they cannot create a ClusterRoleBinding to a role that grants that permission. To allow a user to create/update role bindings:

1. Grant them a role that allows them to create/update RoleBinding or ClusterRoleBinding objects, as desired.
2. Grant them permissions needed to bind a particular role:
    - implicitly, by giving them the permissions contained in the role.
    - explicitly, by giving them permission to perform the `bind` verb on the particular Role(or ClusterRole)

### ServiceAccout permissions

Default RBAC policies grant scoped permissions to control-plane components, nodes, and controllers,but grant *no permissions* to service accounts outside the `kube-system` namespace(beyond discovery permissions given to all authenticated users).

This allows you to grant particular roles to particular ServiceAccounts as needed. Fine-grained role bindings provide greater security, but require more effort to administrate. Broader grants can give unnecessary(and potentially escalating) API access to ServiceAccounts, but are easier to administrate.
