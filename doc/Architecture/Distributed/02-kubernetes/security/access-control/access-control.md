# Access Control

https://kubernetes.io/docs/concepts/security/controlling-access/

Users access the Kubernetes API using `kubectl`, client libraries, or by making REST requests. Both human users and Kubernetes service accounts can be authorized for API access. When a request reaches the API, it goes through serveral stages, illustrated in the following steps:

- Authentication
- Authorization
- Admission Control

## Transport security

In a typical Kubernetes cluster, the API serves on port 443, protected by TLS. The API server presents a certificate. This certificate may be signed using a private certificate authority(CA), or based on a public key infrastructure linked to a generally recognized CA.

If your cluster uses a private certificate authority, you need a copy of that CA certificate configured into your `~/.kube/config` on the client, so that you can trust the connection and be confident it was not intercepted.

Your client can present a TLS client certificate at this stage.

## Authentication

Once TLS is established, the HTTP request moves to the Authentication step. The cluster creation script or cluster admin configures the API server to run one or more Authenticators modules. Authenticators are described in more detail in Authentication.

The input to the authentication step is the entire HTTP request; however, it typically examines the headers and/or client certificate.

Authentication modules include client certificates, password, and plain tokens, bootstrap tokens,and JSON Web Tokens (used for service accounts).

Multiple authentication modules can be specified, in which case each one is tried in sequence, util one of them succeeds.

If the request cannot be authenticated, it is rejected with HTTP status code 401. Otherwise, the user is authenticated as specific `username`, and the user name is available to subsequent steps to use in their decisions. Some authenticators also provide the group memberships of the user, while other authenticators do not.

While Kubernetes uses usernames for access control decisions and in request logging, it does not have a `User` object nor does it store usernames or other information about users in its API.

## Authorization

After the request is authenticated as coming from a specific user, the request must be authorized.

A request must include the username of the requester, the requested action, and the object affected by the action. The request is authorized if an existing policy declares that the user has permissions to complete the requested action.

For example, if Bob has the policy below, then he can read pods only in the namespace `projectCaribou`:

```yml
{
    "apiVersion": "abac.authorization.kubernetes.io/v1beta1",
    "kind": "Policy",
    "spec": {
        "user": "bob",
        "namespace": "projectCaribou",
        "resource": "pods",
        "readonly": true
    }
}
```

if Bob makes the following request, the request is authorized because he is allowed to read objects in the `projectCaribou` namespace:

```yml
{
  "apiVersion": "authorization.k8s.io/v1beta1",
  "kind": "SubjectAccessReview",
  "spec": {
    "resourceAttributes": {
      "namespace": "projectCaribou",
      "verb": "get",
      "group": "unicorn.example.org",
      "resource": "pods"
    }
  }
}
```

If Bob makes a request to write(`create` or `update`) to the objects in the `projectCaribou` namespace, his authorization is denied. If Bob makes a request to read (`get`) objects in a different namespace such as `projectFish`, then his authorization is denied.

Kubernetes authorization requires that you use common REST attributes to interact with existing organization-wide or cloud-provider-wide access control systems. It is important to use REST formatting because these control systems might interact with other APIs besides the Kubernetes API.

Kubernetes supports multiple authorization modules, such as ABAC mode, RBAC mode, and Webhook mode. When an administrator creates a cluster, they configure the authorization modules that should be used in the API server. If more than one authorization modules are configured, Kubernetes checks each module, and if any module authorizes the request, then the request can proceed. If all of the modules deny the reuquest, then the request is denied(HTTP status code 403).

## Admission control

Admission Control modules are software modules that can modify or reject requests. In addition to the attributes available to Authorization modules, Admission Control modules can access the contents of the object that is being created or modified.

Admission controllers act on requests that create, modify, delete, or connect to an object. Admission controller do not act on requests that merely read objects. When multiple admission controllers are configured, they are called in order.

Unlike Authentication and Authorization modules, if any admission controller module rejects, then the request is immediately rejected.

In addition to rejecting object, admission controllers can also set complex defaults for fields.

Once a requst passes all admission controllers, it is validated using the validation routinges for the corresponding API object, and then written to the object store

## API server ports and IPs

The previous discussion applies to requests sent to the secure port of the API server(the typical case). The API server can actually serve on 2 ports:

By default, the Kubernetes API server serves HTTP on 2 ports:

1. `localhost` port:
    - is intended for testing and bootstrap, and for other components of the master node(scheduler, controller-manager) to talk to the API
    - no TLS
    - default is port 8080
    - default IP is localhost, change with `--insecure-bind-address` flag.
    - request **bypasses** authentication and authorization modules.
    - request handled by admission control modules.
    - protected by need to have host access

2. "Secure port"
    - use whenever possible
    - uses TLS. Set cert with `--tls-cert-file` and key with `--tls-private-key-file` flag.
    - default is port 6443, change with `--secure-port` flag.
    - defulat IP is first non-localhost network interface, change with `--bind-address` flag.
    - request handled by authentication and authorization modules.
    - request handled by admission control module.
    - authentication and authorization modules run