# Graphite

## 相关链接
[graphite](https://graphite.readthedocs.io/en/latest/overview.html)
[graphite-web](https://github.com/graphite-project/graphite-web)
[carbon](https://github.com/graphite-project/carbon)
[whisper](https://github.com/graphite-project/whisper)
[carbon-relay-ng](https://github.com/graphite-ng/carbon-relay-ng)
[go-carbon](https://github.com/lomik/go-carbon)
[prometheus vs graphite](https://logz.io/blog/prometheus-vs-graphite/)

## What is Graphite
Graphite is a highly scalable real-time graphing system. As a user, you write an application that collect numeric time-series data that you are interested in graphing, and send it to Graphite's processing backend, carbon, which stores the data in Graphite's specialized database. The data can be visualized through graphite's web interfaces.

## How scalable is Graphite?
From a CPU perspective, Graphite scales horizontally on both frontend and backend, meaning you can simply add more machines to the mix to get more thoughput. It is also fault tolerant in the sense that losing a backend machine will cause a minimal amount of data loss(whatever that machine had cached in the memory) and will not disrupt the system if you have suffient capacity remaining to handle the load.

From an I/O perspective, under load Graphite performs lots of tiny I/O operations on lots of different files rapidly. This is because each distinct metrics sent to Graphite is stored in its own database files, similar to how many tools built on top of RRD work.

High volume(a few thousand distinct metrics updating every time) pretty much requires a good RAID array and/or SSDs. Graphite's backend caches incoming data if the disks cannot keep up with the large number fo small write operations that occur(each data point is only a few bytes, but most standard disks cannot do more than a few thousand I/O operations per second, even if they are tiny). When this occurs, Graphit's database engine, whisper, allow carbon to write multiple data points at once, thus increasing overall thoughout only at the cost of keeping excess data cached in  memory until it can be written.

## Carbon daemons
When we talk about "Carbon" we mean one or more various daemons that make up the storage backend of a Graphite installation.

All of the carbon daemon listen for time-series data and can accept it over common set of protocols.


### Carbon-cache
Carbon-cache accepts metrics over various protocols and write them to disk as efficiently as possible. This requires caching metrics values in RAM as they are received, and flushing them to disk on an interval using the underlying whisper library. It also provides a query services for in-memory metric datapoints, used by the Graphite webapp to retrieve "hot data"

### Carbon-relay
carbon-relay serves two distinct purposes: replication and sharding.

