# API Overview

[reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/)

[aliyun](https://zhuanlan.zhihu.com/p/94949824)

[kubernetes api chinese doc](http://docs.kubernetes.org.cn/31.html)

[go client example](https://github.com/kubernetes/client-go/tree/master/examples)

## API Version

- Alpha
- Beta
- Stable

## API Groups

API Groups make it easier to extend the Kubernetes API. The API group is specified in a REST path and in the `apiVerion` field of a serialized object.

- The *core* group is found at REST path `/api/v1`. The core group is not specified as part of the `apiVersion` field, for example, `apiVersion: v1`.
- The named groups are at REST path `/apis/$GROUP_NAME/$VERSION` and use `apiVersion: $GROUP_NAME/$VERSION` (for example, `apiVersion: batch/v1`).

## Enabling or disabling API groups

Certain resources and API groups are enabled by default. You can enable or disable them by setting --runtime-config on the API server. The --runtime-config flag accepts comma separated <key>=<value> pairs describing the runtime configuration of the API server. For example:

- to disable batch/v1, set --runtime-config=batch/v1=false
- to enable batch/v2alpha1, set --runtime-config=batch/v2alpha1

## Kubernetes API Concepts

The Kubernetes API is a resource-based (RESTful) programmatic interface provided via HTTP. It supports retrieving, creating, updating, and deleting primary resources via the standard HTTP verbs (POST, PUT, PATCH, DELETE, GET), includes additional subresources for many objects that allow fine grained authorization (such as binding a pod to a node), and can accept and serve those resources in different representations for convenience or efficiency. It also supports efficient change notifications on resources via "watches" and consistent lists to allow other components to effectively cache and synchronize the state of resources.

## Standard API terminology

Kubernetes generally leverages standard RESTful terminology to describe the API concepts:

- A resource type is the name used in the URL (`pods`, `namespaces`, `services`)
- All resource types have a concrete representation in JSON (their object schema) which is called a **kind**
- A list of instances of a resource type is known as a **collection**
- A single instance of the resource type is called a **resource**

## Efficient detection of changes

## Retrieving large results sets in chunks

On large clusters, retrieving the collection of some resource types may result in very large responses that can impact the server and client. For instance, a cluster may have tens of thousands of pods, each of which is 1-2kb of encoded JSON. Retrieving all pods across all namespaces may result in a very large response (10-20MB) and consume a large amount of server resources.

Starting in Kubernetes 1.9 the server supports the ability to break a single large collection request into many smaller chunks while preserving the consistency of the total request. Each chunk can be returned sequentially which reduces both the total size of the request and allows user-oriented clients to display results incrementally to improve responsiveness.

To retrieve a single list in chunks, two new parameters limit and continue are supported on collection requests and a new field continue is returned from all list operations in the list metadata field. A client should specify the maximum results they wish to receive in each chunk with limit and the server will return up to limit resources in the result and include a continue value if there are more resources in the collection. The client can then pass this continue value to the server on the next request to instruct the server to return the next chunk of results. By continuing until the server returns an empty continue value the client can consume the full set of results.

Like a watch operation, a continue token will expire after a short amount of time (by default 5 minutes) and return a 410 Gone if more results cannot be returned. In this case, the client will need to start from the beginning or omit the limit parameter.

## Receiving resources as Tables

## Alternate representations of resources

By default Kubernetes returns objects serialized to JSON with content type application/json. This is the default serialization format for the API. However, clients may request the more efficient Protobuf representation of these objects for better performance at scale. The Kubernetes API implements standard HTTP content type negotiation: passing an Accept header with a GET call will request that the server return objects in the provided content type, while sending an object in Protobuf to the server for a PUT or POST call takes the Content-Type header. The server will return a Content-Type header if the requested format is supported, or the 406 Not acceptable error if an invalid content type is provided.

Not all API resource types will support Protobuf, specifically those defined via Custom Resource Definitions or those that are API extensions. Clients that must work against all resource types should specify multiple content types in their Accept header to support fallback to JSON:

```Accept: application/vnd.kubernetes.protobuf, application/json```

### Protobuf encoding

Kubernetes uses an envelope wrapper to encode Protobuf responses. That wrapper starts with a 4 byte magic number to help identify content in disk or in etcd as Protobuf (as opposed to JSON), and then is followed by a Protobuf encoded wrapper message, which describes the encoding and type of the underlying object and then contains the object.

## Resource deletion

## Dry-run

The modifying verbs (POST, PUT, PATCH, and DELETE) can accept requests in a dry run mode. Dry run mode helps to evaluate a request through the typical request stages (admission chain, validation, merge conflicts) up until persisting objects to storage. The response body for the request is as close as possible to a non-dry-run response. The system guarantees that dry-run requests will not be persisted in storage or have any other side effects.

### Make a dry-run request

Dry-run is triggered by setting the dryRun query parameter. This parameter is a string, working as an enum, and the only accepted values are:

- `All`: Every stage runs as normal, except for the final storage stage. Admission controllers are run to check that the request is valid, mutating controllers mutate the request, merge is performed on PATCH, fields are defaulted, and schema validation occurs. The changes are not persisted to the underlying storage, but the final object which would have been persisted is still returned to the user, along with the normal status code. If the request would trigger an admission controller which would have side effects, the request will be failed rather than risk an unwanted side effect. All built in admission control plugins support dry-run. Additionally, admission webhooks can declare in their configuration object that they do not have side effects by setting the sideEffects field to "None". If a webhook actually does have side effects, then the sideEffects field should be set to "NoneOnDryRun", and the webhook should also be modified to understand the DryRun field in AdmissionReview, and prevent side effects on dry-run requests.
- Leave the value empty, which is also the default: Keep the default modifying behavior

## Server Side Apply

## Resource Versions

Resource versions are strings that identify the server's internal version of an object. Resource versions can be used by clients to determine when objects have changed, or to express data consistency requirements when getting, listing and watching resources. Resource versions must be treated as opaque by clients and passed unmodified back to the server. For example, clients must not assume resource versions are numeric, and may only compare two resource version for equality (i.e. must not compare resource versions for greater-than or less-than relationships).

## Kubernetes Deprecation Policy

