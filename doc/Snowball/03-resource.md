# Resource

## Queue

## Cache
- local cache
    - google guava
- centralized
    - redis
- no persistence
- no replication
    - sharding 4~8
- one port per service
- recommended structure in redis
    - kv
    - hash
    - list
    - sorted set
    - pipeline batch
- other structure should be known for its usage
    - bitset
    - multi exec transaction
    - hyperloglog
    - pubsub

## DB
- MySQL
- innodb
- avoid transaction
- Read/Write seperation
- avoid complex query
- sharding