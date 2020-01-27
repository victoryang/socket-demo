# Service Discovery

[zk vs etcd vs consul](https://www.liangzl.com/get-article-detail-12683.html)

[zk vs etcd](https://blog.csdn.net/zzhongcy/article/details/89401204)

## Comparision

|Feature|Consul|zookeeper|etcd|
|:-:|:-:|:-:|:-:|
|Health Check|service status<br>memory/disk|keepalived connection|heart beat|
|Multiple Data Center|yes|-|-|
|KV Storage|yes|yes|yes|
|Consistency|ca|cp|cp|
|Interfaces|http and dns|client|http/grpc|
|Watch|long polling|yes|long polling|
|Self Monitoring|metrics|-|metrics|
|Security|acl/https|acl|https|
|Spring Cloud|yes|yes|yes|

