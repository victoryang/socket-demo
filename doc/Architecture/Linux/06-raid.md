# RAID

[raid全解](https://blog.csdn.net/ensp1/article/details/81318135)

## RAID
Redundant Array of Inexpensive disks
由于当时大容量磁盘比较昂贵，RAID的基本思想是将多个容量较小、相对廉价的磁盘进行有机组合，从而以较低的成本获得与昂贵大容量磁盘相当的容量、性能、可靠性。随着磁盘成本和价格的不断降低， RAID 可以使用大部分的磁盘， “廉价”已经毫无意义。
于时 RAID 变成了独立磁盘冗余阵列（ Redundant Array of Independent Disks ）。

## 基本原理
RAID是由多个独立的高性能磁盘驱动器组成的磁盘子系统，从而提供比单个磁盘更高的存储性能和数据冗余的技术。RAID是一类多磁盘管理技术，其向主机环境提供了成本适中、数据可靠性高的高性能存储。

定义：
一种磁盘阵列，部分物理存储空间用来记录保存在剩余空间上的用户数据的冗余信息。当其中某一个磁盘或访问路径发生故障时，冗余信息可用来重建用户数据。

磁盘条带化虽然与RAID定义不符，通常还是称为RAID（即RAID0）

RAID的初衷是为大型服务器提供高端的存储功能和冗余的数据安全。在整个系统中，RAID被看作是由两个或更多磁盘组成的存储空间，通过并发地在多个磁盘上读写数据来提高存储系统的I/O性能。大多数RAID等级具有完备的数据校验、纠正措施，从而提高系统的容错性，甚至镜像方式，大大增强系统的可靠性，Redundant也由此而来。

RAID的两个关键目标是提高数据可靠性和I/O性能。磁盘阵列中，数据分散在多个磁盘中，然而对于计算机系统来说，就像一个单独的磁盘。通过把相同的数据同时写入多个到多块磁盘（典型的如镜像），或将计算的校验数据写入阵列中来获得冗余能力，当单块磁盘出现故障时可以保证不会导致数据丢失。有些RAID等级允许更多的磁盘同时发生故障，比如RAID6，可以是两块磁盘同时损坏。在这样的冗余机制下，可以用新盘替换故障磁盘，RAID会自动根据剩余磁盘中的数据和校验数据重建丢失的数据，保证数据一致性和完整性。数据分散保存在RAID中的多个不同磁盘上，并发数据读写要大大优于单个磁盘，因此可以获得更高的I/O带宽。当然，磁盘阵列会减少全体磁盘的总可用存储空间，牺牲空间换取更高的可靠性和性能。比如，RAID1存储空间利用率仅有50%，RAID5会损失其中一个磁盘的存储容量，空间利用率为(n-1)/n。

## 关键技术
### 镜像
镜像是一种冗余技术，为磁盘提供保护功能，防止磁盘发生故障而造成数据丢失。

### 数据条带
 RAID 由多块磁盘组成，数据条带技术将数据以块的方式分布存储在多个磁盘中，从而可以对数据进行并发处理。这样写入和读取数据就可以在多个磁盘上同时进行，并发产生非常高的聚合 I/O ，有效提高了整体 I/O 性能，而且具有良好的线性扩展性。这对大容量数据尤其显著，如果不分块，数据只能按顺序存储在磁盘阵列的磁盘上，需要时再按顺序读取。而通过条带技术，可获得数倍与顺序访问的性能提升。

 ### 数据校验
 镜像具有高安全性、高读性能，但冗余开销太昂贵。数据条带通过并发性来大幅提高性能，然而对数据安全性、可靠性未作考虑。相对镜像，数据校验大幅缩减了冗余开销，用较小的代价换取了极佳的数据完整性和可靠性。数据条带技术提供高性能，数据校验提供数据安全性， RAID 不同等级往往同时结合使用这两种技术。

 采用数据校验时，RAID要在写入数据同时进行校验计算，并将得到的校验数据存储在 RAID 成员磁盘中。校验数据可以集中保存在某个磁盘或分散存储在多个不同磁盘中，甚至校验数据也可以分块，不同 RAID 等级实现各不相同。当其中一部分数据出错时，就可以对剩余数据和校验数据进行反校验计算重建丢失的数据。校验技术相对于镜像技术的优势在于节省大量开销，但由于每次数据读写都要进行大量的校验运算，对计算机的运算速度要求很高，必须使用硬件 RAID 控制器。在数据重建恢复方面，检验技术比镜像技术复杂得多且慢得多。

 海明校验码和 异或校验是两种最为常用的 数据校验算法。

 ## RAID等级
 ### 1 RAID 0
 RAID0 是一种简单的、无数据校验的数据条带化技术。实际上不是一种真正的 RAID ，因为它并不提供任何形式的冗余策略。 

 ### 2 RAID1
 RAID1 称为镜像，它将数据完全一致地分别写到工作磁盘和镜像 磁盘，它的磁盘空间利用率为 50% 。

 ### 3 RAID2
 RAID2 称为纠错海明码磁盘阵列，其设计思想是利用海明码实现数据校验冗余。

 ### 4 RAID3
 RAID3 是使用专用校验盘的并行访问阵列，它采用一个专用的磁盘作为校验盘，其余磁盘作为数据盘，数据按位可字节的方式交叉存储到各个数据盘中。RAID3 至少需要三块磁盘。

 ### 5 RAID4
 RAID4 与 RAID3 的原理大致相同，区别在于条带化的方式不同。

 ### 6 RAID5
RAID5 应该是目前最常见的 RAID 等级，它的原理与 RAID4 相似，区别在于校验数据分布在阵列中的所有磁盘上，而没有采用专门的校验磁盘。对于数据和校验数据，它们的写操作可以同时发生在完全不同的磁盘上。因此， RAID5 不存在 RAID4 中的并发写操作时的校验盘性能瓶颈问题。另外， RAID5 还具备很好的扩展性。当阵列磁盘 数量增加时，并行操作量的能力也随之增长，可比 RAID4 支持更多的磁盘，从而拥有更高的容量以及更高的性能。

　　RAID5 的磁盘上同时存储数据和校验数据，数据块和对应的校验信息存保存在不同的磁盘上，当一个数据盘损坏时，系统可以根据同一条带的其他数据块和对应的校验数据来重建损坏的数据。与其他 RAID 等级一样，重建数据时， RAID5 的性能会受到较大的影响。

　　RAID5 兼顾存储性能、数据安全和存储成本等各方面因素，它可以理解为 RAID0 和 RAID1 的折中方案，是目前综合性能最佳的数据保护解决方案。 RAID5 基本上可以满足大部分的存储应用需求，数据中心大多采用它作为应用数据的保护方案。