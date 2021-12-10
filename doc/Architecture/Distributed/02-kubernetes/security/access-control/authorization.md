# Authorization

In Kubernetes, you must be authenticated(logged in) before your request can be authorized(granted permission to access).

Kubernetes expects attributes that are common to REST API requests. This means that Kubernetes authorization works with existing organization-wide or cloud-provider-wide access control systems which may handle other APIs beside the Kubernetes API.

## Determine Whether a Request is Allowed or Denied

Kubernetes authorizes API request using the API server. It evaluates all of the request attributes against all policies and allows or denies the request. All part of an API request must be allowed by some policy in order to proceed. This means that permissions are denied by default.

(Although Kubernetes uses the API server, access controls and policies that depend on specific fields of specific kinds of objects are handled by Admission Controllers)

When multiple authorization modules are configured, each is checked in sequence. If any authorizer approves or denies a request, that decision is immediately returned and no other authorizer is consulted. If all modules have no opinion on the request, then the request is denied. A deny retures an HTTP status code 403.

## Review Your Request Attributes

Kubernetes reviews only the following API request attributes:

- **user** - The `user` string provided during authentication
- **group** - The list of group names to which the authenticated user belongs
- **extra** - A map of arbitrary string keys to string values, provided by the authentication layer
- **API** - Indicates whether the request is for an API resource.
- **Request path** - Path to miscellaneous non-resource endpoints like `/api` or `/healthz`
- **API request verb** - API verbs like `get`,`list`,`create`,`update`,`patch`,`watch`,`delete`,and `deletecollection` are used for resource requests. To detemine the request verb for a resource API endpoint
- **HTTP request verb** - Lowercased HTTP methods like `get`,`post`,`put`,and `delete` are used for non-resource requests
- **Resource** - The ID or name of the resource that is being accessed(for resource requests only) -- For resource requests using `get`,`update`,`patch`,and `delete` verbs, you must provide the resource name.
- **Subresource** - The subresource that is being accessed(for resource requests only)
- **Namespace** - The namespace of the object is being accessed (for namespaced resource requests only)
- **API group** - The API Group being accessed(for resource requests only). An empty string designates the core API group

## Determine the Request Verb

**Non-resource requests** Requests to endpoints other than `/api/v1/...` or `/apis/<group>/<version>/...` are considered "non-resource requests", and use the lower-cased HTTP method of the request as the verb. For example, a `GET` request to endpoints like `/api` or `/healthz` would use `get` as the verb.

**Resource requests** To determine the request verb for a resource API endpoint, review the HTTP verb used and whether or not the request acts on an individual resource or a collection of resources:

|HTTP verb|request verb|
|-|-|
|POST|create|
|GET,HEAD|get(for individual resources), list(for collections, including full object content), watch(for watching an individual resource or collection of resources)|
|PUT|update|
|PATCH|patch|
|DELETE|delete(for individual resources), deletecollection(for collections)|

Kubernetes sometimes checks authorization for additional permissions using specialized verbs. For example:

- PodSecurityPolicy
    - `use` verb on `podsecuritypolicies` resources in the `policy` API group
- RBAC
    -`bind` and `escalate` verbs on `roles` and `clusterrole` resources in the `rbac.authorization.k8s.io` API group
- Authentication
    - `impersonate` verb on `users`, `groups`, and `serviceaccounts` in the core API group, and the `userextras` in the `authentication.k8s.io` API group.

## Authrization Modes

The Kubernetes API server may authorize a request using one of several authorization modes:

- **Node** - A special-purpose authorization mode that grants permissions to kubelets based on the pods they are scheduled to run.
- **ABAC** - Attribute-based access control(ABAC) defines an access control paradim whereby access rights are granted to users through the use of policies which combine attribute together. The policies can use any type of attributes(use attributes, resource attributes,object,environment attributes,etc)
- **RBAC** - Role-based access control (RBAC) is a method of regulating access to computer or network resources based on the roles of individual users within an enterprise. In this context, access is the ability of an individual user to perform user to perform a specific task, such as view, create, or modify a file.
    - When specified RBAC(Role-Based Access Control) uses the `rbac.authorization.k8s.io` API group to drive authorization decisions, allowing admins to dynamically configure permission policies through the Kubernetes API
    - To enable RBAC, start the apiserver with `--authorization-mode=RBAC`
-**Webhook** - A Webhook is an HTTP callback: an HTTP POST that occurs when something happens; a simple event-notification via HTTP POST. A web application implementing WebHooks will POST a message to a URL when certain things happen