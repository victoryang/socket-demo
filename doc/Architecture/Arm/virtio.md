# Virtio

https://zhuanlan.zhihu.com/p/699017345

[virtio spec](https://docs.oasis-open.org/virtio/virtio/v1.2/virtio-v1.2.html)

[vring](https://blog.csdn.net/qq_41596356/article/details/128437440)

[vhost-net](https://blog.csdn.net/qq_45698138/article/details/131450006)

[vhost-user-gpu](https://www.bilibili.com/video/BV1ch4y1T7ct/)

[eventfd and irqfd](https://www.bilibili.com/opus/767992010168598537)

[eventfd](https://zhuanlan.zhihu.com/p/662264612)

[rust virtio *important*](https://rcore-os.cn/rCore-Tutorial-Book-v3/chapter9/2device-driver-4.html)

[virtio-gpu](https://zhuanlan.zhihu.com/p/696241924)

[rust-vhost-device-gpu.spec](https://github.com/dorindabassey/rust-vhost-device-gpu-copr/blob/main/rust-vhost-device-gpu.spec)

[gpu in crosvm](https://crosvm.dev/doc/devices/virtio/gpu/struct.Gpu.html)

[mesa](https://joyxu.github.io/2021/05/13/gpu06/)

[docker rust](https://www.cnblogs.com/dibtp/p/19036320)

[virtio-gpu](https://blog.csdn.net/stray2b/article/details/123487106)

[开源mali安装](https://cloud.tencent.com/developer/article/1571948)

[wayland](https://wayland-book.com/introduction/high-level-design.html)

[guest->qemu->virglrender](https://jishuzhan.net/article/1992454802513657858)

[opengl](https://wiki.t-firefly.com/zh_CN/Firefly-Linux-Guide/manual_ubuntu.html#opengl-vulkan)

[android input vsync](https://www.androidperformance.com/2019/11/04/Android-Systrace-Input/#/%E7%B3%BB%E5%88%97%E6%96%87%E7%AB%A0%E7%9B%AE%E5%BD%95)

[android minigbm](https://deepwiki.com/feijiang1/deepwiki_minigbm)

[GPU 虚拟化典型应用场景分类](https://blog.csdn.net/stray2b/article/details/131371122)

## hypervisor

[acrn-hypervisor](https://zhuanlan.zhihu.com/p/597780462)

[windriver hypervisor](https://www.windriver.com.cn/downloads/files/HVP_Data_Sheet-CN.pdf)


## vhost-device-gpu

### rutabaga_gfx

**编译rutabaga_gfx**
```bash
https://crates.io/crates/rutabaga_gfx
```

**编译gfxstream** https://github.com/mesonbuild/meson/discussions/11292


**meson 交叉编译**

https://mesonbuild.com/Cross-compilation.html

meson setup build/ --cross-file meson-cross-file.txt

meson install -C build

**add rust project to buildroot** https://www.elebihan.com/posts/how-to-add-a-buildroot-package-for-a-cargo-crate.html

## virtio-snd

### ALSA

[ALSA 应用开发](https://zhuanlan.zhihu.com/p/443728870)

[alsa](https://geek-blogs.com/blog/advanced-linux-sound-architecture/)

[alsa 音频子系统](https://blog.csdn.net/zhilin_tang/article/details/159346643)

#### 音频相关概念

音频信号是一种连续变化的模拟信号，但计算机只能处理和记录二进制的数字信号，由自然音源得到的音频信号必须经过一定的变换，才能送到计算机中作进一步的处理。

数字音频系统通过将声波的波型转换成一系列二进制的数据，来实现对原始声音的重现，实现这一步骤的设备常被称为（A/D）。A/D 转换器以每秒钟上万次的速率对声波进行采样，每个采样点都记录下了原始模拟声波在某一时刻的状态，通常称为样本（sample），而一秒钟所采样的数目则称为采样频率，通过将一串连续样本连接起来，就可以在计算机中描述这一段声音。对于采样过程中的每一个样本来说，数字音频系统会分配一定存储来记录声波的振幅，一般称之为采样分辨率或者采样精度，采样精度越高，声音还原时就会越细腻。

数字音频涉及到的概念非常多，对于在linux下进行音频编程的程序员来说，最重要的是7406解声音数字化的两个关键步骤：采样和量化。

- 采样就是每隔一定时间就读一次声音信号的幅度，从本质上讲，采样是时间上的数字化。
- 量化则是将采样得到的声音信号幅度转化成数字值，从本质上讲，量化则是幅度上的数字化。

**采样频率**

采样频率是指将模拟声音波形进行数字化时，每秒钟抽取声波幅度样本的次数。采样频率的选择应该遵循乃奎斯特采样理论：如果对某一模拟信号进行采样，则采样后可还原的最高信号频率只有采样频率的一半，或者说只要采样频率高于输入信号最高频率的两倍，就能从采样信号系统重构原始信号。

**量化位数**

量化位数是对模拟音频信号的幅度进行数字化，它决定了模拟信号数字化以后的动态范围，常用的有8位，12位和16位。量化位越高，信号的动态范围越大，数字化后的音频信号就越可能接近原始信号，但所需要的存储空间也越大。

音频应用中常用的数字表示方法为 脉冲编码调制（pulse-code-modulated，PCM）信号。在这种表示方法中，每个采样周期用一个数字电平对模拟信号的幅度进行编码，得到的数字波形是一组采样自输入模拟波形的近似值。由于所有A/D转化器的分辨率都是有限的，所以在数字音频系统中，A/D转换器带来的量化噪声是不可避免的。

[alsa-linux 音频框架](https://zhuanlan.zhihu.com/p/695919368)

#### ALSA 框架

**PCM**
PCM 是 Pulse-Code modulation 的缩写，中文译名是脉冲编码调制。在现实生活中，人耳听到的声音是模拟信号，PCM就是要把声音从模拟转换成数字信号，他的原理简单地说就是利用一个固定的频率对模拟信号进行采样，采样后的信号在波形上看就像一串连续的幅值不一样的脉冲，把这些脉冲的幅值按一定的精度进行量化，这些量化后数值被连续地输出、传输、处理或记录到存储介质中，所有这些组成了数字音频的产生过程。

PCM信号的两个重要指标是采样频率和量化精度，目前，CD音频的采样频率通常为44100Hz,量化精度是 16bit。通常，播放音乐时，应用程序从存储介质中读取音频数据（MP3、WMA...）经过解码后，最终送到音频驱动程序中的就是PCM数据，反过来，在录音时，音频驱动不停地把采样所得的PCM数据送回到应用程序，由应用程序完成压缩、存储等任务。所以，音频驱动的两大核心任务就是：

- playback 如何把用户空间的应用程序发过来的PCM数据，转化成人耳可以辨别的模拟音频
- capture 把 mic 拾取得到模拟信号，经过采样、量化，转化成PCM信号送回给用户空间的应用程序

**ASOC**
ASOC 是 ALSA System on Chip 的缩写，是针对片上系统引入的中间层：为了适应Platform和Codec硬件上的分离，对基础 ALSA 框架实现进行了解耦。

建立在标准ALSA驱动层上，为了更好地支持嵌入式处理器和移动设备中的音频 Codec 的一套软件体系。在ASOC出现之前，内核对于SoC中的音频已经有部分的支持，不过会有一些局限性：

1. Codec 驱动与SoC CPU的底层耦合过于紧密，这种不理想会导致代码的重复。
2. 音频事件没有标准的方法来通知用户，例如耳机、麦克风的插拔和检测，这些事件在移动设备中是非常普通的，而且通常都需要特定于机器的代码进行重新对音频路径进行配置。
3. 当进行播放或录音时，驱动会让整个 codec 处于上电状态，这对于PC没问题，但对于移动设备来说，这意味着大量的电量。同时也不支持通过改变过取样频率和偏置电流来达到省电的目的。

ASOC层旨在解决这些问题并提供以下功能：

1. 编解码器独立性。允许在其他平台和机器上重用编解码驱动程序。
2. 编解码器和SoC之间的简单 I2S/PCM 音频接口设置。每个 SoC 接口和编解码都会向内核注册其音频接口功能，并在已知应用硬件参数时进行匹配和配置。
3. 动态音频电源管理（DAPM）。DAPM 始终自动将编解码器设置为最低功耗状态。这包括根据内部编解码器音频路由和任何活动流来打开/关闭内部电源模块。
4. 减少爆音和咔哒声。通过以正确的顺序打开/关闭编解码器电源（包括使用数字静音），可以减少爆裂声和咔嗒声。ASoC 向编解码器发出何时更改电源状态的信号。
5. 机器特定控制：允许机器向声卡添加控制（例如扬声器放大器的音量控制）。

为了实现这一切，ASOC 基本上将嵌入式音频系统拆分为多个可重复使用的组件驱动程序：

1. Codec class drivers：codec class driver 与平台无关，包含音频控件，音频接口功能、编解码 DAPM 定义和编解码 IO 函数。如果需要，此类可扩展到 BT、FM和MODEM IC。codec class driver 应该是可以在任何体系结构和机器上运行的通用代码。
2. Platform class drivers：平台类驱动程序包括音频 DMA引擎驱动程序、数字音频接口 DAI 驱动程序 （例如 I2S、AC97、PCM）以及该平台的任何音频 DSP 驱动程序。
3. Machine class driver：机器驱动程序类充当粘合剂，描述并将其他组件驱动程序绑定在一起以形成 ALSA 声卡设备。它处理机器特定的控制和机器级音频事件。


### Android Audio

[audio alsa driver](https://blog.csdn.net/Ciellee/article/details/100976674)

[/dev/pcm 节点创建及 open](https://blog.csdn.net/Ciellee/article/details/100997885)