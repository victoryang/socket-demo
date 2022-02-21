# Linux IPC

https://blog.csdn.net/weiwangchao_/article/details/17380031

- pipe
- signal
- message queue
- shared memory
- semaphore
- socket

## pipe

- 只支持单向数据流
- 只能用于具有亲缘关系的进程之间
- 管道的缓冲区大小有限
- 传输无格式字节流

## signal

- 软件层次上对中断机制的模拟
- 异步通信
- 来源：硬件（键盘输入）和软件（kill）

## message queue

- 有格式
- 消息的容量有限制

## shared memory

## semaphore