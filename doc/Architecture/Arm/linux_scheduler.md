# Linux scheduler

## 模块化设计

[模块化设计](https://blog.csdn.net/whh_bjqy/article/details/159168204)

### 掌握调度子系统架构的价值

|应用场景|核心价值|
|-|-|
|学术论文|深入分析CFS公平调度算法、EDF实时调度理论的形式化验证|
|工业实时系统|为ROS2/PLC/运动控制器配置最优调度策略，确保任务deadlines|
|内核开发|编写新的调度类（如GPU调度、AI推理调度），扩展现有框架|
|性能优化|定位调度延迟、上下文切换开销，优化多核负载均衡|
|安全认证|IEC 61508/SIL 认证中，证明调度器的可预测性与确定性|

### 调度子系统的5层架构

1. 调度子系统整体架构图

```
|用户空间: sched_setscheduler()系统调用|
|系统调用层|kernel/sched/core.c sys_sched_setscheduler() -> __sched_setscheduler()|
|调度类层（Scheduler Classes）|stop sched_class||deadline sched_class||rt sched_class||cfs sched_class||idle sched_class||
|核心调度器：pick_next_task() 负载均衡器：load_balance()/migrate_task()|
|底层抽象：sched_entity/task_struct 时间记账：update_curr()/scheduler_tick()|
```

2. 五大调度类详解

|调度类|优先级范围|核心算法|应用场景|
|-|-|-|-|
|Stop|最高（-1）|无（强制抢占）|CPU热插拔，负载均衡的迁移线程|
|Deadline|0-99|EDF（最早截止时间优先）|实时视频编解码、工业控制|
|RT|0-99|FIFO/RR|硬实时任务，如ROS2控制循环|
|CFS|100-139|红黑树+vruntime|普通进程，web服务器，数据库|
|Idle|最低（140）|无|空闲任务，节能|

3. 关键术语

|术语|定义|源码位置|
|-|-|-|
|sched_class|调度类抽象结构体，定义12个回调函数|kernel/sched/sched.h|
|task_struct|进程描述符，包含sched_class指针|include/linux/sched.h|
|sched_entity|CFS调度实体，维护vruntime|kernel/sched/sched.h|
|sched_dl_entity|Deadline调度实体，维护deadlines|同上|
|rq(run queue)|每个CPU的运行队列|kernel/sched/sched.h|

### 应用场景：自动驾驶实时控制系统的调度架构

在自动驾驶域控制器中，调度子系统的架构设计直接影响功能安全与实时性。典型场景包含三类关键任务：

**感知融合任务（Deadline类）**： 激光雷达点云处理周期 10ms，截止时间 8ms，适用SCHED_DEADLINE确保EDF调度，错过deadline触发安全降级。

**车辆控制任务（RT类）**：横向/纵向控制循环2ms，优先级99，使用SCHED_FIFO，禁止被抢占，确保控制指令确定性输出。

**日志与诊断任务（CFS类）** 系统状态记录、远程诊断，nice值10，在感知与控制任务空闲时运行，不影响实时性。

三类任务共存于同一系统，Stop类处理CPU热插拔时间，Idle类在空闲时节能。调度器的优先级层级确保高优先级任务始终优先，负载均衡在多核之间分配任务，避免单核过载导致的 deadline miss。


## 7种策略映射关系与适用场景全梳理

[7种策略映射关系与适用场景全梳理](https://blog.csdn.net/whh_bjqy/article/details/159168435)

### 调度策略为何是Linux性能的核心杠杆
Linux内核调度器是操作系统最核心的子系统之一，直接决定进程何时获得CPU时间、获得多长时间。在多处理器普及、异构计算兴起的今天，调度策略的选择已成为系统性能调优的首要杠杆：

- 云计算场景：容器密度提升30%往往源于CFS参数的精细化调优
- 实时控制场景：机械臂控制周期从10ms抖动降至100us确定性，依赖SCHED_FIFO的正确配置
- 桌面交互场景：chrome多标签流畅度与CFS的vruntime计算密切相关

然而，生产环境中策略滥用现象普遍：将普通计算任务绑定SCHED_FIFO导致系统starvation，或对实时任务使用CFS造成不可接受的延迟抖动。

### 核心概念：调度类与策略的层次架构

1. 调度类—策略的容器
Linux内核采用模块化调度类设计，通过sched_class结构体实现策略扩展：

调度类优先级固定，高优先级类存在可运行任务时，低优先级类完全无法获得CPU。

2. 七种调度策略详解

### 应用场景：云原生数据库的混合负载调度

在云原生分布式数据库的生产环境中，典型负载构成：

- OLTP事务处理：短查询，延迟敏感，需SCHED_NORMAL+高nice优先级
- OLAP分析查询：长扫描，吞吐优先，适用SCHED_BATCH减少缓存污染
- Raft共识协议：心跳超时严格，关键路径使用SCHED_FIFO保证10ms内响应
- compaction后台任务：磁盘IO密集型，SCHED_IDLE避免影响前台

调度策略误用案例：某金融客户将compaction设为SCHED_FIFO优先级50，导致Raft心跳任务在compaction运行时被饿死，引发频繁leader选举，P99延迟从5ms恶化至2s。修复方案：compaction降级为SCHED_IDLE，Raft保持SCHED_FIFO优先级80

### 实验案例与步骤：七策略全量实验
