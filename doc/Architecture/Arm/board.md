# Board

## RK

[资料下载](https://www.t-firefly.com/doc/download/106.html)

[RK3566](https://wiki.t-firefly.com/zh_CN/ROC-RK3566-PC/linux_compile.html#bian-yi-buildroot-gu-jian)

[Firefly Linux 开发指南](https://wiki.t-firefly.com/zh_CN/Firefly-Linux-Guide/first_use.html)

[rk3588 firmware development](https://blog.csdn.net/weixin_74253011/article/details/143433477)

[rk3399](https://blog.csdn.net/dengjin20104042056/article/details/132636148)

[uboot](https://blog.csdn.net/weixin_42744482/article/details/147272744)

[uboot 传递内存区域给 kernel](https://kingdix10.github.io/zh-cn/docs/03-boot/u-boot/uboot_mem/)

[kernel fiqdebugger](https://www.oryoy.com/news/jie-mi-ubuntu-xi-tong-zhong-chuan-kou-ttyfiq-yi-chang-zhi-mi-jiao-ni-qing-song-pai-cha-yu-jie-jue.html)

[kernel modules init debug](https://zhuanlan.zhihu.com/p/797680247)

[kernel chinese translation](https://www.kernel.org/doc/html/latest/translations/zh_CN/core-api/cpu_hotplug.html)

[获取uboot fdt](https://forums.raspberrypi.com/viewtopic.php?t=381555)

[uefi](https://kagurazakakotori.github.io/ubmp-cn/part2/loader/bootoption.html)

## uart 

minicom -D /dev/ttyUSB0

fatload mmc 1 0x10000000 bao.bin

ldr x10, =0xfeb50000
mov x11, #'A'
str x11, [x10]

bao-hypervisor main:
4a877f432f4598682ff91b6f584c6bbb49b363ed
e8df29d6491cea50dd1bc5668ebc6e7ab2871da8


Adding bank: 0x00200000 - 0x08400000 (size: 0x08200000)   130M
Adding bank: 0x09400000 - 0xf0000000 (size: 0xe6c00000)   3.5G
Adding bank: 0x100000000 - 0x3fc000000 (size: 0x2fc000000)  11G
Adding bank: 0x3fc500000 - 0x3fff00000 (size: 0x03a00000)    58M
Adding bank: 0x4f0000000 - 0x500000000 (size: 0x10000000)    256M

tftp 0x10000000 bao.bin

setenv ipaddr 192.168.1.101;setenv serverip 192.168.1.100;tftp 0xa0000000 bao.bin;go 0xa0000000

setenv ipaddr 192.168.1.101;setenv serverip 192.168.1.100;tftp 0xa0000000 bao.bin;tftp 0x20000000 linux-backend.bin;go 0xa0000000

dtc -I dtb -O dts -o output.dts input.dtb

genext2fs -B 1024 -b 32768 -d ./build/disk ./build/disk.img

find . | cpio -o -H newc | gzip > ../ramdisk-new.cpio.gz

lz4 -c -d -l ./ramdisk > test.cpio
toybox cpio -i -F test.cpio

kptr_restrict=0
kernel.printk=7

kernel/init/main.c            do_one_initcall

backend:
reg = <0 0x00200000 0 0x08200000>,
<0 0x09400000 0 0x80000000>,        // 0x09400000 ~ 0x89400000
<1 0x00000000 0 0xa0000000>;        // 0x100000000 ~ 0x1a0000000

shm:
reg = <0 0x90000000 0 0x06000000>   // 0x90000000 ~ 0x96000000

frontend1:
reg = <0x0 0x09400000 0x0 0x40000000   // 0x200000000 ~ 0x240000000
    0x0 0x90000000 0x0 0x03000000>;    // 

frontend2:
reg = <0x0 0x09400000 0x0 0x40000000   // 0x260000000 ~ 0x2a0000000
    0x0 0x90000000 0x0 0x03000000>;    // 


CMA
//reg = <0x00 0x10000000 0x00 0x8000000>;
reg = <0x00 0x00400000 0x00 0x8000000>;

CONFIG_NR_CPU=4



[   23.102833] rockchip_thermal_probe
[   23.106493] of_reset_control_array_get num:2
[   23.111089] of_reset_control_array_get, name: clock-controller@fd7c0000
[   23.118030] of_reset_control_array_get, name: clock-controller@fd7c0000
[   23.125037] reset_controller
[   23.128060] assert name: clock-controller@fd7c0000
[   23.133128] clock rockchip_softrst_assert id:193, base: 0000000003880a00
[   23.140206] assert name: clock-controller@fd7c0000
[   23.145334] clock rockchip_softrst_assert id:192, base: 0000000003880a00

tsadbc init fails
rockchip-thermal fec00000.tsadc: Missing rockchip,grf property
rockchip-thermal fec00000.tsadc: failed to register sensor 1: -19
rockchip-thermal fec00000.tsadc: failed to register sensor[1] : error = -19

kernel/drivers/thermal/rockchip_thermal.c
static const struct rockchip_tsadc_chip rk3588_tsadc_data = {
1937  	/* top, big_core0, big_core1, little_core, center, gpu, npu */
1938  	.chn_id = {0, 1, 2, 3, 4, 5, 6},
1939  	.chn_num = 7, /* seven channels for tsadc */
}
compared with id defined in dts, remove 1 and 2



注释 CONFIG_ROCKCHIP_CSU 不支持 csu

disable i2c@fd880000       关联 cpu big0 和 cpu big1


[   20.562849] /phy@fed90000: Failed to get clk index: 3 ret: -517

drm init fails
[   24.776859] [drm] failed to init overlay plane Cluster0-win1
[   24.776917] [drm] failed to init overlay plane Cluster1-win1
[   24.776972] [drm] failed to init overlay plane Cluster2-win1
[   24.777028] [drm] failed to init overlay plane Cluster3-win1

media/platform/rockchip/hdmirx/rk_hdmirx.c      big 0 -> l2
/*
4617  	 * Bind HDMIRX's FIQ and driver interrupt processing to big cpu1
4618  	 * in order to quickly respond to FIQ and prevent them from affecting
4619  	 * each other.
4620  	 */
4621  	if (sip_cpu_logical_map_mpidr(0) == 0) {
4622  		cpu_aff = sip_cpu_logical_map_mpidr(5);
4623  		hdmirx_dev->bound_cpu = 5;
4624  	} else {
4625  		cpu_aff = sip_cpu_logical_map_mpidr(1);
4626  		hdmirx_dev->bound_cpu = 1;
4627  	}


disable  watchdog@feaf0000

disabled usb@fc880000   usb@fc8c0000     fatal error
[   25.623489] usb 2-1: new high-speed USB device number 2 using ehci-platform
non-standard psci, smc_fid: 0x82000024, x1:0x3, x2:0xd2, x3:0x1
non-stand[   27.938370] ehci-platform fc880000.usb: fatal error
[   27.938452] ehci-platform fc880000.usb: HC died; cleaning up


cpu0 enters idle dues to fails of rockchip_cpufreq_driver initialition
disable cpu-idle-states = <0x10> for cpu0

CONFIG_INITRAMFS_SOURCE="/opt/src/output/rootfs/rootfs.cpio.gz"
CONFIG_INITRAMFS_ROOT_UID=0
CONFIG_INITRAMFS_ROOT_GID=0


./build.sh rootfs:ubuntu22
BR2_DEFCONFIG='' KCONFIG_AUTOCONFIG=/opt/src/buildroot/output/rockchip_rk3588/build/buildroot-config/auto.conf KCONFIG_AUTOHEADER=/opt/src/buildroot/output/rockchip_rk3588/build/buildroot-config/autoconf.h KCONFIG_TRISTATE=/opt/src/buildroot/output/rockchip_rk3588/build/buildroot-config/tristate.config BR2_CONFIG=/opt/src/buildroot/output/rockchip_rk3588/.config HOST_GCC_VERSION="9" BASE_DIR=/opt/src/buildroot/output/rockchip_rk3588 SKIP_LEGACY= CUSTOM_KERNEL_VERSION="5.10" BR2_DEFCONFIG=/opt/src/buildroot/configs/rockchip_rk3588_defconfig /opt/src/buildroot/output/rockchip_rk3588/build/buildroot-config/conf --defconfig=/opt/src/buildroot/output/rockchip_rk3588/.config.in Config.in

ip link set eth0 down

ip a a  192.168.1.101/24 brd 192.168.1.255 scope global dev eth0

ping 192.168.1.100

ip addr del 192.168.1.101/24 dev eth0

Unable to handle kernel paging request at virtual address 0000003d778ffa00
[  139.343083] Mem abort info:
[  139.345877]   ESR = 0x96000005
[  139.348935]   EC = 0x25: DABT (current EL), IL = 32 bits
[  139.354241]   SET = 0, FnV = 0
[  139.357299]   EA = 0, S1PTW = 0
[  139.360440] Data abort info:
[  139.363318]   ISV = 0, ISS = 0x00000005
[  139.367156]   CM = 0, WnR = 0
[  139.370120] user pgtable: 4k pages, 39-bit VAs, pgdp=00000001579d7000
[  139.376553] [0000003d778ffa00] pgd=0000000000000000, p4d=0000000000000000, pud=0000000000000000
[  139.385253] Internal error: Oops: 96000005 [#1] SMP
[  139.390127] Modules linked in: rtk_btusb 8723du

bluetooth fatal
BR2_PACKAGE_RKSCRIPT_USB=n   disable usbdevice

disable hdmi_tx/vop smmu

domU:
331520 ./output/build/linux-headers-custom/usr/include/linux/version.h