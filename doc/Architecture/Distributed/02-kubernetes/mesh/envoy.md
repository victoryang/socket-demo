# Envoy

## Introduction

https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy

## Listeners

### Listeners

The Envoy configuration supports any number of listeners within a single process. Generally we recommend running a single Envoy per machine regardless of the number of configured listeners. This allows for easier operation and a single source of statistics. Envoy supports both TCP and UDP listeners.

#### TCP

Each listener is independently configured with some number filter chains, where an individual chain is selected based on its match criteria. An individual filter chain is composed of one or more network level (L3/L4) filters. When a new connection is received on a listener, the appropriate filter chain is selected, and the configured connection local filter stack is instantiated and begins processing subsequent events. The generic listener architecture is used to perform the vast majority of different proxy tasks that Envoy is used for (e.g., rate limiting, TLS client authentication, HTTP connection management, MongoDB sniffing, raw TCP proxy, etc.).

Listeners are optionally also configured with some number of listener filters. These filters are processed before the network level filters, and have the opportunity to manipulate the connection metadata, usually to influence how the connection is processed by later filters or clusters.

Listeners can also be fetched dynamically via the listener discovery service (LDS).

#### UDP

### Listener filters

As discussed in the listener section, listener filters may be used to manipulate connection metadata. The main purpose of listener filters is to make adding further system integration functions easier by not requiring changes to Envoy core functionality, and also make interaction between multiple such features more explicit.

The API for listener filters is relatively simple since ultimately these filters operate on newly accepted sockets. Filters in the chain can stop and subsequently continue iteration to further filters. This allows for more complex scenarios such as calling a rate limiting service, etc. Envoy already includes several listener filters that are documented in this architecture overview as well as the configuration reference.

### Network Filter Chain

### Network (L3/L4) filters

As discussed in the listener section, network level (L3/L4) filters form the core of Envoy connection handling. The filter API allows for different sets of filters to be mixed and matched and attached to a given listener. There are three different types of network filters:

### TCP proxy

Since Envoy is fundamentally written as a L3/L4 server, basic L3/L4 proxy is easily implemented. The TCP proxy filter performs basic 1:1 network connection proxy between downstream clients and upstream clusters. 

The TCP proxy filter will respect the connection limits imposed by each upstream cluster’s global resource manager. The TCP proxy filter checks with the upstream cluster’s resource manager if it can create a connection without going over that cluster’s maximum number of connections, if it can’t the TCP proxy will not make the connection.

### UDP proxy

### DNS proxy

The DNS filter supports responding to forward queries for A and AAAA records. The answers are discovered from statically configured resources, clusters, or external DNS servers. The filter will return DNS responses up to to 512 bytes. If domains are configured with multiple addresses, or clusters with multiple endpoints, Envoy will return each discovered address up to the aforementioned size limit.

## HTTP

### HTTP connection management

HTTP is such a critical component of modern service oriented architectures that Envoy implements a large amount of HTTP specific functionality. Envoy has a built in network level filter called the HTTP connection manager. This filter translates raw bytes into HTTP level messages and events (e.g., headers received, body data received, trailers received, etc.). It also handles functionality common to all HTTP connections and requests such as access logging, request ID generation and tracing, request/response header manipulation, route table management, and statistics.

### HTTP filters

### HTTP routing

### HTTP upgrades

Envoy Upgrade support is intended mainly for WebSocket and CONNECT support, but may be used for arbitrary upgrades as well.

#### Websocket over HTTP/2 hops

#### CONNECT support

#### Tunneling TCP over HTTP

### HTTP dynamic forward proxy

Through the combination of both an HTTP filter and custom cluster, Envoy supports HTTP dynamic forward proxy. This means that Envoy can perform the role of an HTTP proxy without prior knowledge of all configured DNS addresses, while still retaining the vast majority of Envoy’s benefits including asynchronous DNS resolution. The implementation works as follows:

- The dynamic forward proxy HTTP filter is used to pause requests if the target DNS host is not already in cache.
- Envoy will begin asynchronously resolving the DNS address, unblocking any requests waiting on the response when the resolution completes.
- Any future requests will not be blocked as the DNS address is already in cache. The resolution process works similarly to the logical DNS service discovery type with a single target address being remembered at any given time.
- All known hosts are stored in the dynamic forward proxy cluster such that they can be displayed in admin output.
- A special load balancer will select the right host to use based on the HTTP host/authority header during forwarding.
- Hosts that have not been used for a period of time are subject to a TTL that will purge them.
- When the upstream cluster has been configured with a TLS context, Envoy will automatically perform SAN verification for the resolved host name as well as specify the host name via SNI.

## Upstream clusters

### Cluster manager

Envoy’s cluster manager manages all configured upstream clusters. Just as the Envoy configuration can contain any number of listeners, the configuration can also contain any number of independently configured upstream clusters.

Upstream clusters and hosts are abstracted from the network/HTTP filter stack given that upstream clusters and hosts may be used for any number of different proxy tasks. The cluster manager exposes APIs to the filter stack that allow filters to obtain a L3/L4 connection to an upstream cluster, or a handle to an abstract HTTP connection pool to an upstream cluster (whether the upstream host supports HTTP/1.1 or HTTP/2 is hidden). A filter stage determines whether it needs an L3/L4 connection or a new HTTP stream and the cluster manager handles all of the complexity of knowing which hosts are available and healthy, load balancing, thread local storage of upstream connection data (since most Envoy code is written to be single threaded), upstream connection type (TCP/IP, UDS), upstream protocol where applicable (HTTP/1.1, HTTP/2), etc.

Clusters known to the cluster manager can be configured either statically, or fetched dynamically via the cluster discovery service (CDS) API. Dynamic cluster fetches allow more configuration to be stored in a central configuration server and thus requires fewer Envoy restarts and configuration distribution.

### Cluster warming

When clusters are initialized both at server boot as well as via CDS, they are “warmed.” This means that clusters do not become available until the following operations have taken place.

- Initial service discovery load (e.g., DNS resolution, EDS update, etc.).
- Initial active health check pass if active health checking is configured. Envoy will send a health check request to each discovered host to determine its initial health status.

The previous items ensure that Envoy has an accurate view of a cluster before it begins using it for traffic serving.

When discussing cluster warming, the cluster “becoming available” means:

- For newly added cluster, the cluster will not appear to exist to the rest of Envoy until it has been warmed. I.e., HTTP routes that reference the cluster will result in either a 404 or 503 (depending on configuration).
- For updated clusters, the old cluster will continue to exist and serve traffic. When the new cluster has been warmed, it will be atomically swapped with the old cluster such that no traffic interruptions take place.

### Service discovery

When an upstream cluster is defined in the configuration, Envoy needs to know how to resolve the members of the cluster. This is known as service discovery.

#### Supported service discovery types

##### Static

##### Strict DNS

##### Logical DNS

##### Original destination

##### Endpoint discovery service

##### Custom cluster

##### On eventually consistent service discovery

Many existing RPC systems treat service discovery as a fully consistent process. To this end, they use fully consistent leader election backing stores such as Zookeeper, etcd, Consul, etc. Our experience has been that operating these backing stores at scale is painful.

### DNS resolution

### Healthy check

### Connection pooling

The utilizing filter code does not need to be aware of whether the underlying protocol supports true multiplexing or not. In practice the underlying implementations have the following high level properties:

#### HTTP/1.1

The HTTP/1.1 connection pool acquires connections as needed to an upstream host (up to the circuit breaking limit). Requests are bound to connections as they become available, either because a connection is done processing a previous request or because a new connection is ready to receive its first request. The HTTP/1.1 connection pool does not make use of pipelining so that only a single downstream request must be reset if the upstream connection is severed.

#### HTTP/2

The HTTP/2 connection pool multiplexes multiple requests over a single connection, up to the limits imposed by max concurrent streams and max requests per connection. The HTTP/2 connection pool establishes as many connections as are needed to serve requests. With no limits, this will be only a single connection. If a GOAWAY frame is received or if the connection reaches the maximum requests per connection limit, the connection pool will drain the affected connection. Once a connection reaches its maximum concurrent stream limit, it will be marked as busy until a stream is available. New connections are established anytime there is a pending request without a connection that can be dispatched to (up to circuit breaker limits for connections). HTTP/2 is the preferred communication protocol, as connections rarely, if ever, get severed.

#### Number of connection pools

Each host in each cluster will have one or more connection pools. If the cluster is HTTP/1 or HTTP/2 only, then the host may have only a single connection pool. However, if the cluster supports multiple upstream protocols, then at least one connection pool per protocol will be allocated. Separate connection pools are also allocated for each of the following features:

- Routing priority
- Socket options
- Transport socket (e.g. TLS) options

#### Health checking interactions

### Load Balancing

### Aggregate Cluster

### Outlier detection

### Circuit breaking

### 