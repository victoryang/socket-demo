# Greedy

[参考链接](https://blog.csdn.net/daaikuaichuan/article/details/80724787)

### 1 定义
    贪心算法在对问题求解时，总是做出在当前看来是最好的选择。也就是说，不从整体最优上加以考虑，他所做出的是在某种意义上的局部最优解.

### 2 贪心的思想
    - 建立数学模型来描述问题
    - 把求解的问题分成若干个子问题
    - 对每一子问题求解，得到子问题的局部最优解
    - 把子问题的局部最优解合成原来解问题的一个解

### 3 贪心的要素
    - 贪心的选择
        在算法中就是仅仅依据当前已有的信息就做出选择，并且以后都不会改变这次选择

    - 最优子结构
        当一个问题的最优解包括其子问题的最优解时，称此问题具有最优子结构性质

### 4 典型应用
    - 背包问题(物体可切分时的0-1背包问题)
    - Huffman编码
    - 单源最短路径
    - Prim算法
    - Kruskal算法