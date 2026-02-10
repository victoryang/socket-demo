# Kernel module

## Linux kernel module

```bash
$ ls -Rl /lib/modules/5.1.15-arch1-1-ARCH //列出所有内核模块
```

## Demo

```c
#<linux/kernel.h>
#<linux/module.h>


static int __init init_my_module(void) {
    printk(KERN_INFO "Hello, my module!\n");
    return 0;
}

static void __exit exit_my_module(void) {
    printk(KERN_INFO "Bye, my module~\n");
}

module_init(init_my_module);
module_exit(exit_my_module);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("csprojectedu");
```

```bash
obj-m += my_module.o
all:
    make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules
clean:
    make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
```

### char device code

[char dev](https://segmentfault.com/a/1190000004474802?utm_source=sf-similar-article)

```c
#include <asm/uaccess.h>
#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/fs.h>

int init_module(void);
void cleanup_module(void);

static ssize_t device_read(struct file *, char *, size_t, loff_t *);
static ssize_t device_write(struct file *, const char *, size_t, loff_t *);
static int device_open(struct inode *, struct file *);
static int device_release(struct inode *, struct file *);

// 我们的设备名字
#define DEVICE_NAME "csprojectedu"

static int major_version;
static int device_is_open = 0;
static char msg[1024];
static char *pmsg;
// 定义设备允许的操作，这里只允许read(), write(), open(), release()
static struct file_operations fops = {
  .read = device_read,
  .write = device_write,
  .open = device_open,
  .release = device_release
};

// 模块注册函数
int init_module() {
  // 将我们定义的设备操作的结构体fops注册到内核中
  major_version = register_chrdev(0, DEVICE_NAME, &fops);
  if (major_version < 0) {
    printk(KERN_ALERT "Register failed, error %d.\n", major_version);
    return major_version;
  }
  // 打印设备名称和版本号，稍后我们会利用这个来挂在模块到/dev/csprojectedu
  printk(KERN_INFO "'mknod /dev/%s c %d 0'.\n", DEVICE_NAME, major_version);
  return 0;
}
// 模块退出函数
void cleanup_module() {
  unregister_chrdev(major_version, DEVICE_NAME);
}

// 处理读模块的逻辑，可以通过cat /dev/csprojectedu 来简单的读取模块
static ssize_t device_read(struct file *filp, char *buffer, size_t length, loff_t *offset) {
  int bytes = 0;
  if (*pmsg == 0) {
    return 0;
  }
  while (length && *pmsg) {
    put_user(*(pmsg++), buffer++);
    length--;
    bytes++;
  }
  return bytes;
}
// 处理写模块的逻辑，我们目前不支持
static ssize_t device_write(struct file *filp, const char *buff, size_t length, loff_t *offset) {
  return -EINVAL;
}

// 处理打开模块的逻辑
static int device_open(struct inode *inode, struct file *file) {
  static int counter = 0;
  if (device_is_open) {
    return -EBUSY;
  }
  device_is_open = 1;
  sprintf(msg, "Device open for %d times.\n", ++counter);
  pmsg = msg;
  try_module_get(THIS_MODULE);
  return 0;
}
// 处理释放/关闭模块的逻辑
static int device_release(struct inode *inode, struct file *file) {
  device_is_open = 0;
  module_put(THIS_MODULE);
  return 0;
```