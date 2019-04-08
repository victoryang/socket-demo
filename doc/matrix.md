# Matrix

[矩阵](http://www2.edu-edu.com.cn/lesson_crs78/self/j_0022/soft/ch0603.html)
[Latex in markdown](http://liyangbit.com/math/jupyter-latex/)
[Latex基本语法](https://blog.csdn.net/u014630987/article/details/70156489)

## 概念
#### 1 定义


>  $由 m \times n 个数 a_{ij} (i=1,2\dots, j=1,2\dots)排成m行n列的矩阵数表$
>
>  $$
>    A = \left \{
>    \begin{matrix}
>    a_{11} & a_{12} & \dots  & a_{1n} \\
>    a_{21} & a_{22} & \dots  & a_{2n} \\
>    \dots & \dots & \dots & \dots \\
>    a_{m1} & a_{m2} & \dots  & a_{mn} \\
>    \end{matrix}
>    \right \}
>  $$
>
> $称为m行n列的矩阵，简称m \times n矩阵.$
> $这 m  \times n个数叫做矩阵A的元素, a_{ij}叫做矩阵A的第i行第j列的元素, $
  $第一个下标i表示元素所在的行, 第二个下标j表示元素所在的列， 矩阵也可以记为：A = \left \{\begin{matrix}a_{ij}\end{matrix}\right \}, 或A=\left \{ \begin{matrix} a_{ij} \end{matrix} \right \}_{m \times n}或A_{m \times n}.$
>
>  $两个矩阵的行数和列数均相等，就称它们是同型矩阵.$
>  $如果两个矩阵A和B是同型矩阵且对应元素相等, 则称为矩阵A和矩阵B相等.$ 

<br>

#### 2 几种特殊的矩阵
- 方阵：矩阵A的行数和列数相等均为n，则称A为方阵
- 行矩阵和列矩阵
- 零矩阵
- 上三角矩阵、下三角矩阵、对角矩阵与单位阵

#### 3 矩阵的初等行变换
- 1 定义
  - 换法变换： 对换矩阵的两行，对换i,j两行，记作$r_{i} \leftrightarrow r_{j}$.
  - 倍法变换： 用非零整数乘矩阵某一行的每个元素. 第i行乘k，记作$r_{i} \times k$.
  - 消法变换： 用数乘矩阵每一行的每个元素后加到另一行的对应元素上. 第j行的k倍加到第i行上，记作$r_{i}+kr_{j}$.

  当A矩阵经过初等变换变成矩阵B时，记作A -> B.
  变换后，矩阵的行列式运算结果不变，但性质不同.

- 2 行阶梯形矩阵和行最简矩阵
  - 行阶梯形矩阵
  $$
    \left \{
    \begin{matrix}
    a_{11} & a_{12} & a_{13} & a_{14} & a_{15} \\
    0 & 0 & a_{23} & a_{24} & a_{25} \\
    0 & 0 & 0 & a_{34} & a_{35} \\
    0 & 0 & 0 & 0 & 0
    \end{matrix}
    \right \}
  $$

  特点： 从第一行开始从左至右划线，遇下一行第一个非零元则下一个台阶，由此逐行向下，可划出一条阶梯线，
  线的下方元素全为零，每个台阶只有一行

  - 行最简矩阵
  $$
    \left \{
    \begin{matrix}
    1 & 0 & -1 & 0 & 4 \\
    0 & 1 & 1 & -1 & 3 \\
    0 & 0 & 0 & 1 & -3 \\
    0 & 0 & 0 & 0 & 0
    \end{matrix}
    \right \}
  $$

  是行阶梯矩阵，且非零行的第一个非零元为1，这些非零元所在的列的其它元素为0