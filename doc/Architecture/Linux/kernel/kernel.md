# Kernel

## Structure

### hash list

```
struct hlist_head { 
    struct hlist_node *first;  //指向每一个hash桶的第一个结点的指针
}; 
struct hlist_node { 
    struct hlist_node *next;   //指向下一个结点的指针
    struct hlist_node **pprev; //指向上一个结点的next指针的指针
};

struct hlist_node *node;
struct hlist_head *h;
hlist_add_head(node, h);
hlist_del(node);
```

### double-linked circular list

- LIST_HEAD_INIT
- LIST_HEAD
- INIT_LIST_HEAD

```c
struct list_head test_head;
INIT_LIST_HEAD(test_head);

struct list_head entry;
list_add(&entry, &test_head);
list_add_tail(&entry, &test_head);

list_del(&entry);

```

### waitqueue

Linux 内核的等待队列是以双循环链表为基础数据结构，与进程调度机制紧密结合，能够用于实现核心的异步事件通知机制。在Linux2.4.21中，等待队列在源代码树include/linux/wait.h中，这是一个通过list_head连接的典型双循环链表。

## function

获取当前进程信息
```c
current
```

获取调度优先级 default 10
```c
task_nice
```

[workqueue](https://blog.csdn.net/scottgly/article/details/6846828)
```c
static DECLARE_WORK(binder_deferred_work, binder_deferred_func);
void __init init_workqueues(void);

create_workqueue(name);
create_singlethread_workqueue(name);

queue_work（struct workqueue_struct *wq, struct work_struct *work);
```

[mmap](https://blog.csdn.net/ddddfang/article/details/88943612)
```c
void *mmap(void *start, size_t length, int prot, int flags, int fd, off_t offset);
```

```bash
unlikely()
```