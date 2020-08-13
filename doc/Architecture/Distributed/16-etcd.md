# Etcd

[etcd性能优化实践](https://www.infoq.cn/article/S4V2cmNeKE186mQmkPVJ)

[etcd doc](https://etcd.io/docs/)

## Data model

etcd is designed to reliably store infrequently updated data and provide reliable watch queries. etcd exposes previous versions of key-value pairs to support inexpensive snapshots and watch history events("time travel queries"). A persistent, multi-version, concurrency-control data model is a good fit for these use cases.

etcd stores data in a multiversion persistent key-value store. The persistent key-value store preserves the previous version of a key-value pair when its value is superseded with new data. The key-value store is effectively immutable; its operations do not update the structure in-place, but instead always generate a new updated structure. All past versions of keys are still accessible and watchable after modification. To prevent the data store from growing indefinitely over time and from maintaining old versions, the store may be compacted to shed the oldest versions of superseded data.

### Logical view

The store's logical view is a flat binary key space. The key space has a lexically sorted index on byte string keys so range queries are inexpensive.

The key sapce maintains multiple revisions. When the store is created, the initial revision is 1. Each atomic mutative operation (e.g., a transaction operation may contain multiple operations) creates a new revision on the key space. All data held by previous revisions. If the store is compacted to save space, revisions before the compact revision will be removed. Revisions are monotonically increasing over the lifetime of a cluster.

A key's life spans a generation, from creation to deletion. Each key may have one or multiple generations. Creating a key increments the version of that key, starting at 1 if the key does not exist at the current revision. Deleting a key generates a key tombstone, concluing the key's current generation by resetting its version to 0. Each modification of a key increments its version; so, versions are monitinically increasing within a key's generation. Once a compaction happens, any generation ended before the compaction revision will be removed, and values set before the compaction revision except the latest one will be removed.

### Physical view

etcd stores the physical data as key-value pairs in a persistent b+tree. Each revision of the store's state only contains the delta from its previous revision to be effcient. A single revision may corresponding to multiple keys in the tree.

The key of key-value pair is a 3-tuple(major, sub,type). Major is the store revision holding the key. Sub differentiates among keys within the same revision. Type is an optional suffix for special value(e.g.,`t` if the value contains a tombstone). The value of the key-valjue pair contains the modification from previous revision, thus one delta from previous revision. The b+tree is ordered by key in lexical byte-order. Ranged lookups over revision deltas are fast; this enables quickly finding modifications from one specific revision to another. Compaction removes out-of-date key-value pair.

etcd also keeps a secondary in-memory btree index to speed up range queries over keys. the keys in the btree index are the keys of the store exposed to user. The value is a pointer to the modification of the persistent b+tree. Compaction removes dead pointers.


## Overview

Starting an etcd cluster statically requires that each member knows another in the cluster.

Once an etcd cluster is up and running, adding or removing members is done via `runtime reconfiguration`.

- Static
- etcd Discovery
- DNS Discovery

### Static

As we know the cluster members, their addresses and the size of the cluster before starting, we can use an offline bootstrap configuration by setting the `initial-cluster` flag. Each machine will get either the following environment variables or commandline.

Note that the URLs specified in `initial-cluster` are the *advertised peer URLs*, i.e. they should match the value of `initial-advertise-cluster` on the respective nodes.

If spinning up multiple clusters (or creating and destroying a single cluster) with same configuration for testing purpose, it is highly recommended that each cluster is given unique `initial-cluster-token`. By doing this, etcd can generate unique IDs and member IDs for the clusters even if they otherwise have the exact same configuration. This can protect etcd from cross-cluster-interaction, which may corrupt the clusters.

etcd listens on `listen-client-urls` to accept client traffic. etcd member advertises the URLs specified in `advertise-client-urls` to other members, proxies,clients. Please make sure the `advertise-client-urls` are reachable from intended clients. A common mistake is setting `advertise-client-urls` to localhost or leave it as default if the remote clients should reach etcd.

The commandline parameters starting with `--initial-cluster` will be ignored on subsquent runs of etcd. Feel free to remove the environment variables or commandline flags after the initial bootstrap process. If the configuration needs changes later(for example, adding or removing members to/from the cluster). see the runtime configuration guide.

## Configuration flags

[example of yaml](https://github.com/etcd-io/etcd/blob/master/etcd.conf.yml.sample)

[doc](https://etcd.io/docs/v3.4.0/op-guide/configuration/)


