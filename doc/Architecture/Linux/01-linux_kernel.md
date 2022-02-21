# Linux Kernel

[参考链接](https://www.infoq.cn/article/Xvlq27eG_GHazRHsOTU2)

[Linux系统调优](https://yq.aliyun.com/articles/637634?spm=a2c4e.11153940.0.0.61904289RxXxtO)

[Linux系统调优2](https://yq.aliyun.com/articles/509658?spm=a2c4e.11153940.0.0.61904289RxXxtO#)

[kernel update](https://www.cnblogs.com/hezhiyao/p/8327339.html)

## Linux subsystem

- Network
- IO
- Process
- Memory
- File system


## Process Managment

### Definition
进程是计算机资源的分配单位，主要包括对系统资源的调度，如CPU，内存空间，是一个程序执行的副本。他们还包括一系列资源，如文件的打开、释放信号，内核内的数据，进程的状态，内存地址的映射，执行的多个线程，数据段的全局标量等。

### Status

- running
- interruptible
- unintertible
- stopped
- zombie

### Cycle

### Thread
线程是比进程资源更小的资源调度单位(LWP)

### Schedule

- [yesterday] 调研 & 讨论 IPC 方案
- [today] 调研Linux IPC 方案
- [blocks] 无