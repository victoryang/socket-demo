# NameSpace

[Linux kernel and namespace](https://cloud.tencent.com/developer/article/1096434)

## Conception
Linux命名空间机制提供了一种资源隔离的解决方案。PID，IPC，Network等系统资源不再是全局性的，而是属于特定的Namespace。Linux Namespace机制为实现基于容器的虚拟化技术提供了很好的基础，LXC(Linux Containers)就是利用这一特性实现了资源的隔离。不同Container内的进程属于不同的Namespace，彼此透明，互不干扰。

Namespace是对全局系统资源的一种封装隔离，使得处于不同Namespace的进程拥有独立的全局系统资源，改变一个Namespace中的系统资源只会影响当前Namespace里的进程，对其他Namespace中的进程没有影响。

## Linux内核支持的Namespace类型

- Cgroup(4.6)
- IPC(2.6)
- Network)(2.6)
- Mount(2.4)
- PID(2.6)
- User(2.6)
- UTS(2.6)

## 命名空间相关API

### clone
创建新的子进程，并加入新的Namespace，当前进程保持不变

    int clone(int (*child_func)(void *), void *child_stac, int flags, void *arg);

### setns
将当前进程加入到已有的Namespace中

    int setns(int fd, int nstype);

    fd： 
        指向/proc/[pid]/ns/目录里相应namespace对应的文件，表示要加入哪个namespace

    nstype：
        指定namespace的类型（上面的任意一个CLONE_NEW*）：
        1. 如果当前进程不能根据fd得到它的类型，如fd由其他进程创建，并通过UNIX domain socket传给当前进程，那么就需要通过nstype来指定fd指向的namespace的类型
        2. 如果进程能根据fd得到namespace类型，比如这个fd是由当前进程打开的，那么nstype设置为0即可

### unshare
当前进程退出指定类型的Namespace，并加入到新创建的Namespace

    int unshare(int flags);

    flags：
        指定一个或者多个namespace的类型

### Kernel
#### task_struct
进程描述符中，有对应的Namespace info
在Linux/include/linux/sched.h定义了task_struct结构体，该结构体是Linux Process完整信息的集合，其中就包含了一个指向Namespace结构体的指针nsproxy(定义见Linux/include/linux/nsproxy.h)

    struct task_struct {
      ...
      /* namespaces */
      struct nsproxy *nsproxy;
      ...
    }


## Namespace

**一个进程可以同时属于多个Namespace**

- UTS NameSpace
    - Hostname and Domain name
- IPC
- PID
    - 每个PID Namespace中的进程可以有其独立的PID；每个容器可以有其PID为1的root进程；也使得容器可以在不同的host之间迁移，因为Namespace中的进程ID和host无关。这也使得容器中的每个进程有两个PID：容器进程和host上的PID
- Mount
    - 每个容器能看到不同的文件系统层次结构
    - 类似chroot，将某一个子目录变成根节点，并且更加安全和灵活
- Network
    - 隔离网络设备，IP地址端口等网络栈
    - Network Namespace允许每个容器拥有自己独立的(虚拟)网络设备，并且容器内的应用可以绑定到自己的端口，每个Namespace内的端口都不会互相冲突
    - 实现不容容器之间的通信，和使用相同的端口
- User
    - 隔离用户和用户组ID。同pid，一个进程的user ID和group ID在User namespace内外是可以不同的；因此可以在namespace内设定某个用户为root用户

## Conclusion

- Namespace的本质就是把原来所有的进程全局共享的资源，分拆成了很多个组进程共享的资源
- 当一个Namespace里面的所有进程都退出时，Namespace也会被销毁
- UTS Namespace
    - 进程的一个属性，属性值相同的一组进程就属于同一个Namespace，跟这组进程之间有没有亲缘关系无关
    - 没有嵌套关系，即不存在说一个Namespace是另一个Namespace的父Namespace