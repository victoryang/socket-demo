# Union FS

[Linux UnionFS](https://blog.csdn.net/songcf_faith/article/details/82787946)

[Docker AUFS](https://www.cnblogs.com/ilinuxer/p/6188654.html)

[IBM aufs performance report paper](https://domino.research.ibm.com/library/cyberdig.nsf/papers/0929052195DD819C85257D2300681E7B/$File/rc25482.pdf)

## Concept

联合文件系统(Union File System)，它可以把多个目录(也叫branch)内容联合挂载到同一个目录下，而目录的物理位置是分开的。UnionFS允许只读和可读写目录并存，就是说可同时删除和增加内容。UnionFS应用的地方很多，比如多个磁盘分区上合并不同文件系统的主目录，或把几张CD光盘合并成一个统一的光盘目录(Archived)。另外，具有写时复制(copy-on-write)功能UnionFS可以把只读和可读写文件系统合并在一起，虚拟上允许只读文件系统的修改可以保存到可写文件系统当中。

Docker镜像的设计中，引入了层(layer)的概念，也就是说，用户制作镜像的每一步操作，都会生成一个层，也就是一个增量rootfs，这样应用A和应用B所在的容器共同引用相同的Debian操作系统层、Golang环境层，而各自有各自应用程序层，和可写层。启动容器的时候通过UnionFS把相关的层挂载到一个目录，作为容器的根文件系统。

需要注意的是，rootfs只是一个操作系统所包含的文件、配置和目录，并不包括操作系统内核。这就意味着，如果你的应用程序需要配置内核参数、加载额外的内核模块，以及跟内核进行直接的交互，就需要注意：这些操作和依赖的对象，都是宿主机操作系统的内核，它对于该机器上的所有容器来说是一个“全局变量”，牵一发而动全身。

## Linux Distribution

- CentOS docker 18-ce overlay2
- Debian docker 17-ce aufs

## Docker and AUFS
AUFS是Docker选用的第一种存储驱动。AUFS具有快速启动容器，高效利用存储和内存的优点。

### Image Layer和AUFS
每一个Docker image都是由一系列的read-only layers组成。Image layers的内容都存储在Docker hosts filesystem的/var/lib/docker/aufs/diff目录下。而/var/lib/docker/aufs/layers/目录则存储着image layer如何堆栈这些layer的metadata。

- /var/lib/docker/aufs/diff/
    - 只读镜像
- /var/lib/docker/aufs/mnt/
    - 联合挂载目录
- /var/lib/docker/aufs/layers/
    - 镜像的堆栈信息

使用docker history可以查看到镜像使用的image layers。

PS: 从Docker 1.10开始，diff目录下的存储镜像layer文件夹不再与镜像ID相同。

#### 容器layer查看

- container id
- cat /sys/fs/aufs/si_{container id}/br[0-9]*

#### whiteout
一般来说只读目录都会有whiteout的属性，所谓whiteout的意思，就是如果在union中删除的某个文件，实际上是位于一个readonly的目录上，那么，在mount的union目录中你将看不到这个文件，但是readonly这个层上我们无法做任何修改，所以我们就需要对这个readonly目录里的文件做whiteout。AUFS的whiteout的实现是通过在上层的可写的目录下建立对应的whiteout隐藏文件来实现的。

- 只读层
    - 挂载方式：ro+wh
- Init层
    - 以'-init'结尾的层，夹在只读层和可读写层之间。Init层是Docker项目单独生成的一个内部层，专门用来存放/etc/hosts、/etc/resolve.conf等信息。需要这样一层的原因是，这些文件本来属于只读的系统镜像层的一部分，但是用户往往需要在启动容器时写入一些指定的值比如hostname，所以就需要在可读写层对它们进行修改。可是，这些修改往往只对当前的容器有效，我们并不希望执行docker commit时，把这些信息连同可读写层一起提交。所以，docker的做法是，在修改了这些文件之后，以一个单独的层挂载的。而用户执行docker commit只会提交可读写层，所以是不包含这些内容的。
- 可读写层
    - 容器的最上面的一层，它的挂载方式为：rw，即read write。在没有写入文件之前，这个目录是空的。而一旦在容器里做了写操作，你修改产生的内容就会以增量的方式出现在这个层中。删除ro-wh层等文件时，也会在rw层创建对应的whiteout文件，把只读层里的文件“遮挡”起来。最上面这个可读写层的作用，就是专门用来存放修改rootfs后产生的增量，无论是增删改，都放生在这里。而当我们使用完了这个被修改的容器之后，还可以使用docker commit和push指令，保存这个被修改过的可读写层，并上传到Docker Hub上，供其他人使用。而与此同时，原先的只读层里的内容则不会有任何变化

最后所有的layer都会被挂载到/var/lib/docker/aufs/mnt/里