# Developer

## gRPC naming and discovery

etcd provides a gRPC resolver to support an alternative name system that fetches endpoints from etcd for discovering gRPC services. The underlying mechanism is based on watching updates to keys prefixed with the service name. 

### Using etcd discovery with go-grpc

The etcd client provides a gRPC resolver for resolving gRPC endpoints with an etcd backend. The resolver is initialized with an etcd client and given a target for resolution:

```go
import (
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"

	"google.golang.org/grpc"
)

...

cli, cerr := clientv3.NewFromURL("http://localhost:2379")
r := &etcdnaming.GRPCResolver{Client: cli}
b := grpc.RoundRobin(r)
conn, gerr := grpc.Dial("my-service", grpc.WithBalancer(b), grpc.WithBlock(), ...)
```

### Managin service endpoints

The etcd resolver treats all keys under the prefix of the resolution target following a "/"(e.g., "my-service/") with JSON-encoded go-grpc `name.Update` values as potential service endpoints. Endpoints are added to the service by creating new keys and removed from the service by deleting keys.

### Adding an endpoint

New endpoints can be added to the service through `etcdctl`:

```bash
ETCDCTL_API=3 etcdctl put my-service/1.2.3.4 '{"Addr":"1.2.3.4","Metadata":"..."}'
```

The etcd client's `GRPCResolver.Update` method can also register new endpoints with a key matching the `Addr`:

```go
r.Update(context.TODO(), "my-service", naming.Update{Op: naming.Add, Addr: "1.2.3.4", Metadata: "..."})
```

### Deleting an endpoint

Hosts can be deleted from the service through `etcdctl`:

```bash
ETCDCTL_API=3 etcdctl del my-service/1.2.3.4
```

The etcd client's `GRPCResolver.Update` method also supports deleting endpoints:

```go
r.Update(context.TODO(), "my-service", naming.Update{Op: naming.Delete, Addr: "1.2.3.4"})
```

### Registering an endpoint with a lease

Registering an endpoint with a lease ensures that if the host can't maintain a keepalive heartbeat(e.g., its machine fails), it will be removed from the service:

```bash
lease=`ETCDCTL_API=3 etcdctl lease grant 5 | cut -f2 -d' '`
ETCDCTL_API=3 etcdctl put --lease=$lease my-service/1.2.3.4 '{"Addr":"1.2.3.4","Metadata":"..."}'
ETCDCTL_API=3 etcdctl lease keep-alive $lease
```

## Interacting with etcd

### Read past version of keys

Application may want to read superseded version of a key. For example, an application may wish to roll back to an old configuration by accessing an earlier version of a key. Alternatively, an application may want a consistent view over multiple keys through multiple requests by accessing key history. Since every modification to the etcd cluster key-value store increments the global revision of an etcd cluster, an application can read superseded keys by providing an older etcd revision.

```bash
etcdctl get --prefix --rev=3 foo
```

### Watch key changes

Applications can watch on a key or a range of keys to monitor for any updates.

```bash
etcdctl watch foo
```

### Watch historical changes of keys

Applications may want to watch for historical changes of keys in etcd. For example, an application may wish to receive all the modifications of a key; if the application stays connected to etcd, then `watch` is good enough. However, if the application or etcd fails, a change may happen during the failure, and the application will not receive the update in real time. To grarantee the update is delivered, the application must be able to watch for historical changes to keys. To do this, an application can specify a historical revision on a watch, just like reading past version of keys.

```bash
# watch for changes on key `foo` since revision 2
$ etcdctl watch --rev=2 foo

# watch for changes on key `foo` and return last revision value along with modified value
$ etcdctl watch --prev-kv foo
```

### Watch progress

Applications may want to check the progress of a watch to determine how up-to-date the watch stream is. For example, if a watch is used to update a cache, it can be useful to know if the cache is stale compared to the revision from a quorum read.

Progress requests can be issued using the "progress" commands in interactive watch session to ask the etcd server to send a progress notify update in the watch stream:

```bash
$ etcdctl watch -i
$ watch a
$ progress
progress notify: 1
# in another terminal: etcdctl put x 0
# in another terminal: etcdctl put y 1
$ progress
progress notify: 3
```

Note: The revision number in the progress notify response is the revision from the local etcd server node that the watch stream is connected to. If this node is partitioned and not part of quorum, this progress notify revision might be lower than the revision returned by a quorum read against a non-partitioned etcd server node.

### Compacted revisions

As we mentioned, etcd keeps revisions so that applications can be read past versions of keys. However, to avoid accumulating an unbounded amount of history, it is important to compact past revisions. After compacting, etcd removes historical revisions, releasing resources for future use. All superseded data with revisions before the compacted revision will be unavailable.

```bash
$ etcdctl compact 5
compacted revision 5

# any revisions before the compacted one are not accessible
$ etcdctl get --rev=4 foo
Error:  rpc error: code = 11 desc = etcdserver: mvcc: required revision has been compacted
```

Note: The current revision of etcd server can be found using get command on any key(existent or non-existent) in json format.

### Grant leases

Applications can grant leases for keys from an etcd cluster. When a key is attached to a lease, its lifetime is bound to the lease's lifetime which in turn is governed by a time-to-live(TTL). Each lease has a minimum time-to-live(TTL) value specified by the application at grant time. The lease's actural TTL value is at least the minumum TTL and is chosen by the etcd cluster. Once a lease's TTL elapses, the lease expires and all attached keys are deleted.

Here is the command to grant a lease:

```bash
# grant a lease with 60 second TTL
$ etcdctl lease grant 60
lease 32695410dcc0ca06 granted with TTL(60s)

# attach key foo to lease 32695410dcc0ca06
$ etcdctl put --lease=32695410dcc0ca06 foo bar
OK
```

### Revoke leases

Applications revoke leases by lease ID. Revoking a lease deletes all of its attached keys.

```bash
$ etcdctl lease revoke 32695410dcc0ca06
lease 32695410dcc0ca06 revoked

$ etcdctl get foo
# empty response since foo is deleted due to lease revocation
```

### Keep leases alive

Applications can keep a lease alive by refreshing its TTL so it does not expire.

```bash
$ etcdctl lease keep-alive 32695410dcc0ca06
lease 32695410dcc0ca06 keepalived with TTL(60)
lease 32695410dcc0ca06 keepalived with TTL(60)
lease 32695410dcc0ca06 keepalived with TTL(60)
```

### Get lease information

Applications may want to know about lease information, so that they can be renewed or to check if the lease still exists or it has expired. Applications may also want to know the keys to which a particular lease is attached. 

## Set up a local cluster

### Local multi-member cluster

#### Starting a Cluster

A `Procfile` at the base of the etcd git repository is provided to easily configure a local multi-member cluster.

## Why gRPC gateway

etcd v3 uses gRPC for its messaging protocol. The etcd project includes a gRPC-based Go client and a command line utility, etcdctl, for communicating with an etcd cluster through gRPC. For languages with no gRPC support, etcd provides a JSON gRPC gateway. This gateway serves a RESTful proxy that translates HTTP/JSON requests into gRPC messages.

### Using gRPC gateway

The gateway accepts a JSON mapping for etcd’s protocol buffer message definitions. Note that key and value fields are defined as byte arrays and therefore must be base64 encoded in JSON. The following examples use curl, but any HTTP/JSON client should work all the same.