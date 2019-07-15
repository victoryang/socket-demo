# Bazel

[参考链接](https://blog.csdn.net/elaine_bao/article/details/78668657)

[HowTo](https://blog.csdn.net/weixin_33712987/article/details/85709898)

## 定义
### WORKSPACE
    Bazel的编译是基于workspace的概念, Workspace是一个存放了所有源代码和Bazel编译输出文件的目录, 也就是整个项目的根目录.
    1. WORKSPACE文件：用于指定当前文件夹就是一个Bazel的工作区。所以WORKSPACE文件总是存在于项目的根目录下
    2. 一个或多个BUILD文件：用于告诉Bazel怎么构建项目的不同部分。

### Package和BUILD文件
    WORKSPACE中代码组织的主要单元是包. Package是相关文件的集合, 以及它们之间的依赖关系的规范. Package被定义为包含名为BUILD的文件的目录, 该文件位于WORKSPACE中的顶级目录下. Package中包含其目录中的所有文件, 以及其下的所有子目录, 除了那些本身包含BUILD文件的子目录.

    一个BUILD文件包含了几种不同类型的指令.其中最重要的编译指令, 它告诉Bazel如何编译想要的输出.

### Target
    Package是一个容器. 
    Package的元素称为target. 大多数target是两种主要类型之一, 即文件和规则. 
    文件进一步分为两种. 
    第二种目标是规则. 规则指定一组输入和一组输出文件之间的关系, 包括从输入中导出输出的必要步骤. 规则的输出始终是生成的文件. 规则的输入可以是源文件, 但也可以是生成的文件. 

### Visibility
    Package组是一组Package, 其目的是限制某些规则的可访问性. Package组由package_group函数定义. 它们有两个属性: 它们包含的包列表及其名称.
    唯一允许引用它们的方法来自规则的visibility属性或package函数的default_visibility属性; 它们不生成或使用文件.

## 使用Bazel编译项目

### Understand the BUILD file
```cpp
cc_binary(
    name = "hello-world",
    srcs = ["hello-world.cc"],
)
```
    In this example, the hello-world target instantiates Bazel's built-in cc_binary rule. The rule tells Bazel to build a self-contained executable binary from the hello-world.cc source file  with no dependencies.

    The attributes in the target explicitly state its dependencies and options. While the name attribute is mandatory, many are optional. For example, in the hello-world target, name is self-explanatory, and srcs specifies the source file from which Bazel builds the target.

### Use lables to reference targets
    //path/to/package:target-name
