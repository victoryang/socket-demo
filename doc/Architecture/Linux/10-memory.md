# Memory

[slab vs buddy](https://blog.csdn.net/u014645605/article/details/76617626)

[Linux系统调优2](https://yq.aliyun.com/articles/509658?spm=a2c4e.11153940.0.0.61904289RxXxtO#)

[buddy and slab](http://c.biancheng.net/view/1284.html)

[Linux内存管理之SLAB原理浅析](https://blog.csdn.net/rockrockwu/article/details/79976833)

[slab内存分配器](https://blog.csdn.net/liuhangtiant/article/details/81259293)

## Memory

<img src="memory.png" />

- buddy
- slab

### Fragmentation

#### Internal Fragmentation
内部碎片是已经被分配出去(能明确指出属于哪个进程)却不能被利用的空间；内部碎片是处于(操作系统分配的用于装载某一进程的内存)区域内部的存储块。

因为所有的内存分配必须起始于可被4、8、16整除(视处理器体系结构而定)的地址或者因为MMU的分页机制的限制，决定内存分配算法仅能把预定大小的内存块分配给客户。假设当某个客户请求一个43字节的内存块时，因为没有适合大小的内存，所以它可能会获得44字节、48字节等稍大一点的字节，因此由所需大小四舍五入而产生的多余空间叫内部碎片。

#### External Fragmentation
外部碎片指的是还没有被分配出去(不属于任何进程)，但由于太小了无法分配给申请内存空间的新进程的内存空闲区域。外部碎片是处于任何两个已分配区域或页面之间的空闲存储块。这些存储块的总和可以满足当前申请的长度要求，但是由于它们的地址不连续或其他原因，使得当前系统无法满足当前申请。

### Allocator

#### Buddy
Buddy系统是为了解决外部碎片问题，它将内存按照2的幂级(order)大小排成链表队列，存放在free_area数组。

#### Slab
每当我们要分配内存的时候，使用malloc申请大小若干的字节内存，内核也必须经常分配内存。之前描述的buddy系统支持按页分配内存，但这个单位对于内核来说太大了。如果需要为一个10字符的字符串分配空间，分配一个4KB或更多的完整页，不仅仅浪费而且不可接受，所以内核需要将页拆分成更小的单位，以便可以容纳大量的小对象。

- Slab可以对小对象空间进行分配，而不需要分配整个页面给对象，这样可以节省空间
- 内核中对于频繁使用的小对象，slab会对此作缓存，避免了频繁的内存分配和回收，提高了速度

- slab_reclaimable
- slab_unreclaimble