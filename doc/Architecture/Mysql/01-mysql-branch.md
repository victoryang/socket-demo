# Mysql Branches Comparison

[链接](https://www.atlantic.net/hipaa-compliant-database-hosting/what-is-mysql-vs-mariadb-vs-percona/)

## MariaDB
It's a potentially better choice if you can't determine what kind of queries your project needs regularly.
- MariaDB includes a larger array of storage engines than either Mysql base of percona, including now a NoSQL option with Cassandra. It also includes percona's XtraDB as an option.
- As of version 10.1, MariaDB offers on-disk database encryption.
- MariaDB offers scalability features including multi-source replication, allowing a single server to replicate from serveral servers.
- The interaction of Global Transaction IDs between MariaDB and MySQL might require careful implementation if you plan to have complex replication schema bridging MySQL flavors. For example, it's possible to replicate from MySQL 5.6 into MariaDB 10.X, but not the reverse.

Overall MariaDB is a mature future-looking branch of MySQL that aims to offer new features and advantages for the database. It is appropriate for large multi-server cloud hosting applcations, especially those with changing query patterns that can take advantage of MariaDB's query optimizer.

## Percona Server
Percona has concentrated on very demanding applcations with their own high-performance alternative to the InnoDB storage engine called XtraDB, including instrumentation tools to tune it. Their features tend towards performance, availability, and scalability improvements for the largest databases with highest thoughput.