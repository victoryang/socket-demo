# ELF

## GNU 链接

[链接脚本](https://zhuanlan.zhihu.com/p/363308789)

[gnu 官方手册](https://ftp.gnu.org/old-gnu/Manuals/ld-2.9.1/html_mono/ld.html)

### 链接过程

- 空间和地址分配
- 符号解析
- 重定位

### 链接参数

链接脚本针对的主要是内存空间和地址分配部分，对于链接过程其它的部分，是连接器自动处理的，毕竟符号解析和重定位通常不会涉及到自定义的操作，同时，可以通过传入命令行参数来控制链接过程，使用 ld --help 可以查看所有的链接器参数：

- -T 指定链接脚本
- -e entry、--entry=entry, 指定程序开始的位置，通常情况下，需要指定程序内部的符号，如果给定的参数不是一个符号，链接器会尝试将参数解析成数字，表示从指定的地址开始执行程序。
- -EB，-EL 指定大小端
- -L --library-path：指定搜索的目录
- -l 链接指定的库
- -o 输出文件名
- -s 丢弃可执行文件中的符号，以减小尺寸
- -static 静态链接
- -nostdlib
- -shared

### 链接脚本

链接脚本主要的工作就是告诉链接器如何将各个目标文件、库中的段进行组织，生成输出文件，该输出文件可以是动态库、可执行文件。

通常每个输入文件都存在多个段，每个段的描述信息被单独保存在段表中，而段的信息单独保存，链接器通过分析段表中的内容，获取段相应的信息，比如段的读写属性、size、文件内偏移。

链接过程之后生成可加载的输出文件，通常是可执行文件，对于输出文件而言，内容是以 segment 的形式进行组织的，组织的依据是根据各个段的读写属性来确定的，对于代码部分通常是 读 + 执行 的权限，对于数据通常是 读 + 写 的属性，相同属性的段被组织在一起，最终在执行时被加载到内存中。

每个 segment 实际上对应两个地址，分别是加载地址 LMA(load memory address) 和虚拟地址 VMA(virtual memory address)，加载地址是程序被加载的地址，而虚拟地址是程序运行的地址，通常情况下这两者是相等的，可以想到，当执行程序时，加载器将程序直接拷贝到它的执行地址，程序就可以直接执行，完全没必要让这两个地址不同，这并没有什么问题。

至于需要区分 LMA 和 VMA 主要是针对一些特殊情况，比如在某些情况下 .data 段被加载到只读的 ROM 中而不是其需要执行的内存位置，然后由程序将这部分 data 数据拷贝到内存中，出现这种情况可能是因为 .data 对应的那块内存还没初始化完成，种种情况通常在启动代码中可以看到。

在 kernel 中同样可以看到 LMA 和 VMA 不同的情况，所有的内核代码将会被编译成一个 vmlinux 的可执行文件，该文件中的 vectors(中断向量)部分在运行时的地址应该是 0xffff0000, 但是在加载时紧随着程序部分，在后续程序执行时才会被 copy 到对应的运行位置。导致 vectors LMA 和 VMA 不同的原因在于：vmlinux 并不是最终执行的文件，它会被 strip 成一个纯数据类型的 Image，被 uboot 加载到一个统一的地址，如果要保证 vector 的 VMA 和 LMA 一致，就得把 Images 加载到 0xffff0000 附近，要不就专门为 vector 加载一次，实际情况自然是不允许的。

### 简单示例

```asm
SECTIONS
{
    . = 0x10000;
    .text : { *(.text) }
    . = 0x8000000;
    .data : { *(.data) }
    .bss : { *(.bss) }
}
```

当输入的文件只包含 .text, .data, .bss 三个段时，这就是一个完整可用的链接脚本，看起来非常简单。

其中 SECTIONS 是链接脚本的关键字，它表示个数据段布局的开始，伴随着 SECTIONS 命令的是 "."，这个 "." 是一个地址定位符，表示随后数据段对应的内存地址，在 SECTIONS 开始处，符号 "." 被定义且初始值为 0。

在上述的示例中，定位符 "." 被赋值为 0x10000，表示 .text 段的开始地址为 0x1000,而后面的 {*(.text)} 表示将所有输入文件中的 .text 段放在输出文件的 .text 段中，*是通配符。

紧接着，地址定位符赋值为 0x8000000，随后放 .data 段，紧接着就是 .bss 段，可以发现这中间没有为地址定位符赋值，在这种情况下，地址定位符 "." 的值为 0x8000000 + sizeof(.data)，即 "." 的值紧随着上一个放置段结束的位置，当然，通常情况下，中间会添加一个对齐参数。

### 脚本命令

#### 程序的入口

程序执行的第一条指令被称为程序的入口，这个入口通常就是在链接脚本指定的，链接脚本中的 ENTRY() 命令可以指定入口地址，关于程序入口地址的指定规则和优先级一次是这样的：

- 命令行通过 -e entry 指令入口地址为 entry，这个entry 可以是一个符号。
- 链接脚本中的 ENTRY(symbol) 命令，这个 symbol 是一个符号
- 如果程序中定义了 start 符号，以这个符号作为入口地址
- .text 段的起始地址
- CPU 的 0 地址处开始

## section 的输入输出

[section 的输入输出](https://zhuanlan.zhihu.com/p/363309047)

## 可重定位目标文件

[可重定位目标文件](https://zhuanlan.zhihu.com/p/363487856)

### 分析工具

readelf 和 objdump


## 通过 Section 为 Linux ELF 程序新增数据

[通过 Section 为 Linux ELF 程序新增数据](https://tinylab.org/elf-sections/)

### 背景介绍

Section 是 Linux ELF程序格式的一种核心数据表达方式，用来存放一个一个的代码块、数据块（包括控制信息块），这样一种模块化的设计为程序开发提供了很大的灵活性。

需要增加一个功能，增加一份代码或者增加一份数据都可以通过新增一个 Section 来实现。Section 的操作在 Linux 内核中有着非常广泛的应用，比如内核压缩，比如把 .config 打包后加到内核映像中。

### 通过内联汇编新增一个 Section

如何创建一个可执行的共享库 中有一个很好的例子

```
asm(".pushsection .interp,\"a\"\n"
    " .string \"/lib/i386-linux-gnu/ld-linux.so.2\"\n"
    ".popsection")
```

通过上述代码新增了一个 `.interp` Section，用于指定动态链接器。简单介绍一下这段内联汇编：

- `asm` 括号内就是汇编代码，这些代码几乎被 “原封不动”地放到汇编语言中间文件中
- 这里采用 `.pushsection`, `.popsection`, 而不是  `.section` 是为了避免之后的代码或者数据被错误地加到这里新增的Section 中来。
- `.pushsection .interp "a"`，这里的 "a"表示 Alloc，会占用内存，这种才会被加到program header table中，因为 program table 会用于创建进程映像。
- `.string` 这行用来指定动态链接器的完整路径

### 通过 gcc `__attribute__` 新增一个 Section

上面的需求可以等价地用 gcc `__attribute__` 编译属性来指定：

```
const char interp[] __attribute__((section(".interp"))) = "/lib/i386-linux-gnu/ld-linux.so.2";
```

### 通过 objcopy 把某个文件内容新增为一个 Section

```
objcopy --add-section .interp=interp.section.txt --set-section-flags .interp=alloc,readonly hello.o
```

注意，必须加上 `--set-section-flags` 配置为 alloc，否则，程序头会不纳入该 Section，结果将是缺少 INTERP 程序头而服务执行。
需要补充的是，本文介绍的 `.interp` 是一个比较特殊的 Section，链接时能自动处理，如果是新增了一个全新的 Section 类型，那么得修改链接脚本，明确告知链接器需要把 Section 放到程序头的哪个 Segment。

### 通过 objcopy 更新某个 Section

以上三种新增 Section 的方式适合不同的需求：汇编语言、C语言、链接阶段，基本能满足日常的开发需要

## ELF 转二进制

[part1](https://tinylab.org/elf2bin-part1/)
[part2](https://tinylab.org/elf2bin-part2/)
[part3](https://tinylab.org/elf2bin-part3/)
[part4](https://tinylab.org/elf2bin-part4/)

### 用 objcopy 把 elf 转成 binary 并运行

把 ELF 文件转化成二进制文件，然后把二进制文件作为一个 Section 加入到另外一个程序，然后在那个程序中访问该 Section 并运行。

```
# hello.s
# as --32 -o hello.o hello.s
# ld -melf_i386 -o hello hello.o
# objcopy -o binary hello hello.bin
```

转成 binary 后的代码和数据如下：

```
$ hexdump -C hello.bin
00000000  31 c0 b0 04 31 db 43 b9  6d 80 04 08 31 d2 b2 0d  |1...1.C.m...1...|
00000010  cd 80 31 c0 89 c3 40 cd  80 48 65 6c 6c 6f 20 57  |..1...@..Hello W|
00000020  6f 72 6c 64 0a 00 00
```

再对照 objdump

```
$ objdump -d -j .text hello
hello1:     file format elf32-i386
Disassembly of section .text:
08048054 <_start>:
 8048054:	31 c0                	xor    %eax,%eax
 8048056:	b0 04                	mov    $0x4,%al
 8048058:	31 db                	xor    %ebx,%ebx
 804805a:	43                   	inc    %ebx
 804805b:	b9 6d 80 04 08       	mov    $0x804806d,%ecx
 8048060:	31 d2                	xor    %edx,%edx
 8048062:	b2 0d                	mov    $0xd,%dl
 8048064:	cd 80                	int    $0x80
 8048066:	31 c0                	xor    %eax,%eax
 8048068:	89 c3                	mov    %eax,%ebx
 804806a:	40                   	inc    %eax
 804806b:	cd 80                	int    $0x80
```

所以，要让 hello.bin 能够运行，必须要把这段 binary 装载在指定的位置，即：

```
$ nm hello | grep " _start"
08048054 T _start
```
这样取到的数据位置才是正确的

#### 如何运行转换过后的二进制

这个是内核压缩支持的惯用做法，先要取到 Load Address，告诉 wrapper
kernel，必须把数据解压到 Load Address 开始的位置。

1. 把 hello.bin 作为一个 Section 加入到目标执行代码中，比如叫 run-bin.c
2. 然后写 ld script 明确把 hello.bin 放到 Load Address 地址上
3. 同时需要修改 ld script 中 run-bin 本身的默认加载地址，否则就覆盖了。也可以先把 hello 的 Load Address 往后搬动

把 hello.bin 作为 `.bin` Section 加入进 run-bin.o
```
$ objcopy --add-section .bin=hello.bin --set-section-flags .bin=contents,alloc,load,readonly run-bin.o
```

```
$ git diff ld.script ld.script.new
diff --git a/ld.script b/ld.script.new
index 91f8c5c..7aecbbe 100644
--- a/ld.script
+++ b/ld.script.new
@@ -11,7 +11,7 @@ SEARCH_DIR("=/usr/local/lib/i386-linux-gnu"); SEARCH_DIR("=/lib/i386-linux-gnu")
 SECTIONS
 {
   /* Read-only sections, merged into text segment: */
-  PROVIDE (__executable_start = SEGMENT_START("text-segment", 0x08048000)); . = SEGMENT_START("text-segment", 0x08048000) + SIZEOF_HEADERS;
+  PROVIDE (__executable_start = SEGMENT_START("text-segment", 0x08046000)); . = SEGMENT_START("text-segment", 0x08046000) + SIZEOF_HEADERS;
   .interp         : { *(.interp) }
   .note.gnu.build-id : { *(.note.gnu.build-id) }
   .hash           : { *(.hash) }
@@ -60,6 +60,11 @@ SECTIONS
     /* .gnu.warning sections are handled specially by elf32.em. */
     *(.gnu.warning)
   }
+  .bin 0x08048054:
+  {
+    bin_entry = .;
+    *(.bin)
+  }
   .fini           :
   {
     KEEP (*(SORT_NONE(.fini)))
```

1. 把 run-bin 的执行地址往前移动到了 0x8046000, 避免代码覆盖
2. 获取到 hello 的 _start 入口地址，并把 .bin 链接到这里。
3. 把 bin_binary 指向 .bin section 链接后的入口


### 允许把 binary 文件加载到任意位置

#### 运行时获取 eip

由于加载地址是任意的，用 .text 中的符号也不行，因为在链接时也一样是写死的，所以，唯一可能得办法是 eip, 即程序计数器。

但是 eip 是没有办法直接通过寄存器获取的，得通过一定技巧来，下面这个函数就可以：

```
eip2ecx:
    movl   (%esp), %ecx
    ret
```

这个函数能够把 eip 放到 ecx 中。



### 动态加载和运行



### 动态计算并修改数据加载地址

