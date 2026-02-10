# Binder

[binder](https://blog.csdn.net/vviccc/article/details/90717764?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1.pc_relevant_antiscan_v2&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1.pc_relevant_antiscan_v2&utm_relevant_index=1)

[binder 解析](https://blog.csdn.net/yuanzhangmei1/article/details/25998209)

[binder](https://www.jianshu.com/p/4c4dcf80d412)

## Kernel Modules

- binder
- ashmem

### Binder

- module_init(binder_init);
- module_exit(binder_exit);

binder_init
```c
// 创建singlethread的任务队列 binder_deferred_workqueue
binder_deferred_workqueue = create_singlethread_workqueue("binder");

// 注册 misc 设备 binder_miscdev
misc_register(&binder_miscdev);
```

```c
static const struct file_operations binder_fops = {
	.owner = THIS_MODULE,
	.poll = binder_poll,
	.unlocked_ioctl = binder_ioctl,
	.compat_ioctl = binder_ioctl,
	.mmap = binder_mmap,
	.open = binder_open,
	.flush = binder_flush,
	.release = binder_release,
};
```

#### open

- 初始化相关队列
- 创建 binder stat
- 初始化 debugfs 相关
- filp->private_data = proc;

```c
static int binder_open(struct inode *nodp, struct file *filp)
```

#### mmap

```c
static int binder_mmap(struct file *filp, struct vm_area_struct *vma)
```

