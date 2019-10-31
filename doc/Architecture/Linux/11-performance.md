# Performance

[Linux性能监控](https://yq.aliyun.com/articles/294523?spm=5176.13394999.0.0.18063d30MkvoN0)

## 子系统

- CPU
- Memory
- IO
- Network

- IO bound，系统会大量消耗内存和底层存储系统，并不消耗过多的CPU和网络资源(除非系统是网络的)。系统消耗CPU来接受IO请求，然后会进入sleep状态。数据库是常被认为是IO bound系统
- CPU bound，系统大量消耗cpu资源。计算型。

## 监视工具

- vmstat
- mpstat
- sar
- iostat
- netstat
- dstat
- iperf
- ethtool

## CPU 

- 运行队列长度
- CPU使用率
- 上下文切换

1. 每个CPU的运行队列长度不要超过3
2. 如果CPU满负荷运行，则应符合如下分布：
    a)User Time: 65%~70%
    b)System Time: 30%~35%
    c)Idle: 0%~5%
3. 上下文切换要结合CPU使用率来看，如果CPU使用率满足上述分布，大量的上下文切换也是可以接受的

常用监控工具

- vmstat
- top
- dstat
- mpstat

```vmstat 1```
r b swpd free buff cache si so bi bo in cs us sy id wa
0 0 104300 16800 95328 72200 0 0 5 26 7 14 4 1 95 0

- r表示运行队列的大小
- b表示由于IO等待而block的线程数量

综上：

- per cpu运行队列不大于3
- 确保CPU使用分布满足70/30原则
- 如果系统时间过长，可能是因为频繁的调度和改变优先级
- CPU Bound总是被惩罚(降低优先级)而IO Bound进程总会被奖励(提高优先级)