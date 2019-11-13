# CPU

[用户态时间](https://blog.csdn.net/u010821666/article/details/78697371)

## Bound

- Real < CPU，表明进程为计算密集型(CPU bound)，利用多核处理器的并行执行优势； 
- Real ≈ CPU，表明进程为计算密集型(CPU bound)，未并行执行； 
- Real > CPU，表明进程为I/O密集型(I/O bound)，多核并行执行优势并不明显。