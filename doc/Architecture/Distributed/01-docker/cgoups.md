# CGroups

[Linux CGroups](https://blog.csdn.net/mnasd/article/details/80588474)

[Linux内核OOM机制分析](http://blog.chinaunix.net/uid-20788636-id-4308527.html)

[深入理解linux cgroups](https://www.cnblogs.com/ryanyangcs/p/11198140.html)

## Conception
Namespace帮助进程隔离出单独的空间，CGroups限制每个空间的大小，保证进程间不会互相争抢

Linux CGroups是Linux内核的一个功能，用来限制、统计和分离一个进程组的资源(CPU、内存、磁盘输入输出等)。如果一个进程加入了某一个控制组，该控制组对Linux的系统资源都有严格的限制，进程在使用这些资源时，不能超过其最大的限制数，例如：memory资源，如果加入控制组的进程使用的memory大于其限制，可能会出现OOM错误。

通过Cgroups，可以方便地限制某个进程的资源占用，并且可以实时地监控进程和统计监控信息。

- Resource Limitation
- Prioritization
- Audit
- Control

在实践中，系统管理员一般会利用CGroups做：

- 隔离一个进程集合(比如：nginx的所有进程)，并限制他们所消费的资源，比如绑定CPU的核
- 为这组进程分配其足够使用的内存
- 为这组进程分配响应的网络带宽和磁盘存储限制
- 限制访问某些设备

## Components
在Cgroups中，task就是指系统的一个进程

### Cgroup
Cgroup中的资源控制都是以cgroup为单位实现的。cgroup表示按照某种资源控制标准划分而成的任务组，包含一个或多个subsystem。一个任务可以加入某个cgroup，也可以从某个cgroup迁移到另外一个cgroup。

### Subsystem
是cgroups的一个资源调度控制器(Resource Controller)。比如CPU子系统可以控制CPU时间分配，内存子系统可以限制cgroup内存使用量。

- blkio
- cpu
- cpuacct
- cpuset
- devices
- freezer
- memory
- net_cls
- net_prio
- ns

查看kernel支持的子系统

    lssubsys -a

### Hierachy
Hierachy是由一系列cgroup以一个树状结构排列而成，每个hierarchy通过绑定对应的subsystem进行资源调度。hierarchy中的cgroup节点可以包含零或多个节点，子节点继承父节点的属性。整个系统可以有多个hierarchy。

### Relationship
Cgroups是通过三个组件互相协作实现的。

- 系统在创建了新的hierarchy之后，系统中所有的进程都会加入这个hierarchy的cgroup根节点，这个cgroup根节点是hierarchy默认创建的
- 一个subsystem只能附加到一个hierarchy上面
- 一个hierarchy可以附加多个subsystem
- 一个进程可以作为多个cgroup成员，但是这些cgroup必须在不同的hierarchy中
- 一个进程fork出子进程时，子进程是和父进程在同一个cgroup中的，也可以根据需要将其移动到其他cgroup中

## Kernel Interface
为了让Cgroups的配置更为直观，通过一个虚拟的树状文件系统配置Cgroups的，通过层级目录虚拟出cgroup树。

1. 创建并挂载一个hierarchy(cgroup树)

> #创建hierarchy挂载点
> mkdir cgroup-test
>
> #挂载一个hierarchy
> mount -t cgroup -o none,name=cgroup-test cgroup-test ./cgroup-test
>
> #列出生成的文件
>ls ./cgroup-test

- cgroup.clone_children
- cgroup.procs
    - 树结构中当前cgroup中的进程组
- notify_on_release和release_agent
    - 最后一个进程退出的时候清理cgroup
- tasks
    - 该cgroup下面的进程ID，如果将一个进程ID写入文件，则该进程加入这个cgroup

2. 在创建好的hierarchy上扩展出两个子group

> mkdir cgroup-1
> mkdir cgroup-2

在一个cgroup目录下创建文件夹，Kernel会把文件夹标记为这个cgroup的子cgroup，并继承父cgroup的属性

3. cgroup中添加和移动进程
同一hierarchy下，一个进程只能在一个cgroup节点上存在，系统的所有进程都会默认在根节点上存在，可以将进程移动到其他cgroup节点，只需要将进程ID写到目标cgroup节点的tasks文件中即可。

4. 通过subsystem限制cgroup中进程的资源

## Docker使用Cgroups
通过为每个容器创建cgroup，并通过cgroup去配置资源的限制和资源监控