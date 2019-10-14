# Disk

[Nvme performance](http://kernelmaker.github.io/nvme_ssd_fio)

[Disk tunning](http://blog.sina.com.cn/s/blog_448574810101k1va.html)

## NVME and PCIe
历史上，大多数SSD使用如SATA、SAS等接口与计算机接口的总线连接。随着固态硬盘在大众市场上的流行，SATA已成为个人电脑中连接SSD的最典型方式；但是，SATA的设计主要是作为机械硬盘驱动器(HDD)的接口，并随着时间的推移越来越难满足速度日益提高的SSD。随着在大众市场的流行，许多固态硬盘的数据速率提升已经放缓。不同于机械硬盘，部分SSD硬盘已经受到SATA最大吞吐量的限制。

在NVMe出现之前，高端的SSD只得采用PCIe总线，但需要使用非标准规范的接口。若只用标准化的SSD接口，操作系统只需要一个驱动程序就能使用匹配规范的所有SSD。这也意味着每个SSD制造商不必使用额外的资源来设计特定接口的驱动程序。

NVMe接口优点

- 性能有数倍的提升
- 可大幅降低延迟
- NVMe可以把最大队列深度从32提升到64000，SSD的IOPS也会得到大幅提升
- 自动功耗切换和动态能耗管理功能大大降低功耗
- NVMe标准的出现解决了不同PCIe SSD之间的驱动适用性问题。