# Systemd

[boot and systemd](https://linux.cn/article-8807-1.html)

[Linux Systemd机制](https://www.cnblogs.com/klb561/p/9245657.html)

[specification](https://www.freedesktop.org/software/systemd/man/init.html#)

[booting application after network up](https://www.freedesktop.org/wiki/Software/systemd/NetworkTarget/)

[unit specification](http://0pointer.de/public/systemd-man/systemd.unit.html)

[systemd and cgroup](https://www.imooc.com/article/72502)

[systemd arch](https://www.cnblogs.com/sparkdev/p/8448237.html)

[systemd issue](https://github.com/moby/moby/issues/16256)

[make cgroupfs as default cgroup manager](https://github.com/moby/moby/pull/17704)

## Unit

systemd基于Unit的概念。Unit是由一个与配置文件对应的名字和类型组成的。

### Unit分类

- service 守护进程
- socket
    - 封装系统和实际的socket
    - 每一个socket unit都有一个相应的服务unit
    - **相应的服务，在第一个“连接”来时启动**
- target
    - 将unit进行逻辑分组，或者在boot up阶段提供well-known同步点
- device
    - 暴露kernel devices，用于基于device的激活动作
    - 封装一个存在于linux设备树中的设备
    - 每一个udev规则标记的设备都会在systemd中作为一个设备unit存在
    - udev的属性设置可以作为配置设备unit依赖关系的配置源
- mount
    - 封装一个挂载点
- automount
- timer
- swap
    - 封装swap内存相关
- path
    - file system object changes
- slice
    - is used to group units which manage system process(such as service and scope unit) in a hierarchical tree for resource management purpose
- scope
    - is similar to service units, but manage foreign process instead of starting them as well

### cgroup

Processes systemd spawns are placed in individual Linux control groups named after the unit which they belong to in the private systemd hierarchy. Systemd uses this to effectively keep track of processes. Control group information is maintained in the kernel, and is accessible via the file system hierarchy(beneath /sys/fs/systemd/), or in tools such as systemd-cgls or ps.

### Transaction System
Systemd has a minimal transaction system: if a unit is requested to start up or shut down it will add it and all its dependencies to a temporary transaction. Then it will verify if the transaction is consistent(i.e. whether the ordering of all units is cycle-free). If it is not, systemd will try to fix it up, and removes non-essential jobs from the transaction that might remove the loop. Also, systemd tries to suppress non-essential jobs in the transaction that would stop a running service.



## Tool

- systemctl
    - 控制systemd系统和服务的状态
- systemd-cgls
    - 以树形递归显示选中的Linux控制组结构层次
- systemadm
    - 图形化前端

### systemctl

#### 输出unit列表
1. systemctl <==> systemctl list-units

2. 输出运行失败的单元:  systemctl --failed

3. systemctl list-unit-files

**所有可用的unit文件存放在/lib/systemd/system/和/etc/systemd/system/**

#### Unit

使用systemctl控制单元时，通常需要使用单元文件的全名，包括扩展名

- 无扩展名，默认扩展.service
- 挂载点自动加载为.mount, 如home.mount
- 设备自动转化为.device单元，如dev-sda2.device

- systemctl start
- systemctl stop
- systemctl restart
- systemctl reload
- systemctl status
- systemctl is-enabled
- systemctl enable
- systemctl disable
- systemctl daemon-reload

#### Unit File Schema

```
[Unit]
Description=Network Time Service

[Service]
ExecStart=/usr/bin/ntpd -n -u ntp:ntp -g

[Install]
WantedBy=multi-user.target
```

1. 所有的unit file都包含[Unit]区域，包含一般设置和简要描述

2. [service]指定当前unit file类型所包含的指定设置，如启动和终止命令等。

3. [Install]包含systemd在(反)安装时要解释的说明

#### Service

管理后台服务

systemctl --type=service

#### Target

**Dependencies**

详见[doc](http://0pointer.de/public/systemd-man/systemd.unit.html)