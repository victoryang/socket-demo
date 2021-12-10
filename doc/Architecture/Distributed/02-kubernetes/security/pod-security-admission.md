# Pod Security Admission

The Kubernetes Pod Security Standards define different isolation levels for Pods. These standards let you define how you want to restrict the behavior of pods in a clear, consistent fashion.

As an Alpha feature, Kubernetes offers a built-in *Pod Security* admission controller, the successor to PodSecurityPolicies. Pod security restrictions are applied at the namespace level when pods are created.

## Pod Security levels

Pod Security admission places requirements on a Pod's Security Context and other related fields according to the three levels defined by the Pod Security Standards: `priviledged`,`baseline`,and `restricted`.

|Mode|Description|
|-|-|
|enforce|Policy violations will cause the pod to be rejected|
|audit|Policy violations will trigger the addition of an audit annotation to the event recorded in the audit log, but are otherwise allowed|
|warn|Policy violations will trigger a user-facing warning, but are otherwise allowed|

## Workload resources and Pod templates

Pods are often created indirectly, by creating a workload object such as a Deployment or Job. The workload object defines a *Pod template* and a controll for the workload resource creates Pods based on that template. To help catch violation early, both the audit and warning modes are applied to the workload resources. However, enforce mode is not applied to workload resources, only to the resulting pod objects.

## Exemptions

You can define *exemptions* from pod security enforcement in order allow the creation of pods that would have otherwise been prohibited due to the policy associated with a given namespace. Exemptions can be statically configured in the Admission Controller configuration.

Exemptions must be explicitly enumerated. Requests meeting exemption criteria are ignored by the Admission Controller (all `enforce`, `audit` and `warn` behaviors are skipped). Exemption dimensions include:

- **Usernames:** requests from users with an exempt authenticated (or impersonated) username are ignored.
- **RuntimeClassNames:** pods and workload resources specifying an exempt runtime class name are ignored.
- **Namespaces:** pods and workload resources in an exempt namespace are ignored.

