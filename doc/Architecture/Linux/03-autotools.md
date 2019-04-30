# GNU Autotools

[参考链接](https://blog.csdn.net/zhengqijun_/article/details/70105077)

[Makefile.am详解](https://www.jianshu.com/p/2f5e586c3402)

## 0 预备知识
Autotools工具链
- autoscan
- aclocal
- autoconf
- autoheader
- automake

Linux系统默认会安装相关工具，安装包为automake

    一般使用步骤
    1 编写源代码项目
    2 执行autoscan, 扫描工作目录, 生成configure.scan文件. 此文件需要重新命名为configure.ac, 并且依据项目构建逻辑, 编辑此配置文件, 执行后续步骤
    3 执行aclocal, 会扫描编辑后的configure.ac, 生成aclocal.m4文件. 该文件主要处理本地的宏定义, 它根据已经安装的宏、用户宏定义和acinclude.m4文件中的宏将configure.ac文件需要的宏集中定义到aclocal.m4中
    4 执行autoheader，生成config.h.in
    5 编写Makefile.am，执行automake，生成Makefile.in
    6 执行autoconf，生成configure
    7 make

如图
<image src="autotools.svg"></image>

## 1 configure.ac 配置说明

|配置|说明|
|:---:|:---:|
|AC_PREREQ|声明autoconf要求的版本号|
|AC_INIT|定义软件名称、版本号、联系方式|
|AM_INIT_AUTOMAKE|必须要的，参数为软件名称和版本号|
|AC_CONFIG_SCRDIR|宏用来侦测所指定的源码文件是否存在, 来确定源码目录的有效性|
|AC_CONFIG_HEADER|宏用于生成config.h文件，以便 autoheader 命令使用|
|AC_PROG_CC|指定编译器，默认GCC|
|AC_CONFIG_FILES|生成相应的Makefile文件，不同文件夹下的Makefile通过空格分隔 <br> 例如 AC_CONFIG_FILES([Makefile, src/Makefile]) |
|AC_OUTPUT|用来设定 configure 所要产生的文件，如果是makefile，configure 会把它检查出来的结果带入makefile.in文件产生合适的makefile|
|||

其中AC_CONFIG_HEADER指定生成config.h，用于autoheader使用，定义来源于
- 用户定义宏
- 已安装信息定义宏，如已安装package或function
- m4内部宏定义

|配置|类型|说明|
|:---:|:---:|:---:|
|AC_ARG_ENABLE|用户定义||
|AC_ARG_WITH|用户定义||
|AC_DEFINE|用户定义||
|AC_CHECK_HEADERS|已安装信息定义宏||
|AC_CHECK_FUNCS|已安装信息定义宏||
|AC_INIT|m4定义|源信息定义相关|
||||

## 2 autoheader命令
    该命令生成config.h.in

## 3 Makefile.am
    Makefile.am 是一种比Makefile更高层次的编译规则，可以和configure.in一起通过调用automake生成configure.in， 然后调用./configure生成Makefile

|宏|说明|可能值|
|:---:|:---:|:---:|
|AUTOMAKE_OPTIONS|软件等级||
|SUBDIRS|先扫描子目录||
|bin_PROGRAMS|软件生成后的可执行文件名称||
|xxx_SOURCES|当前目录源文件||
|xxx_LDFLAGS|链接选项||
|xxx_LDADD|静态链接||
|LIBS |m4定义|动态链接|
||||

## 4 automake命令
    执行automake --add-missing命令
    该命令会依据Makfile.am生成Makefile.in

## 5 autoconf命令
    执行autoconf命令, 会将configure.ac中的宏展开，生成configure脚本
    这个过程可能需要aclocal.m4中的宏

## 6 configure命令
    执行configure命令，生成Makefile

## 7 make命令
    执行make过程