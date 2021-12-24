# Virtio

http://www.linux-kvm.org/page/Virtio

## Paravirtualized drivers for kvm/Linux

virtio was chosen to be the main platform for IO virtualization in KVM

- Virtio was chosen to be the main platform for IO virtualization in KVM
- The idea behind it is to have a common framework for hypervisors for IO virtualization
- More information (although not uptodate) can be found here
- At the moment network/block balloon devices are supported for kvm
- The host implementation is in userspace - qemu, so no drivers is needed in host.

## How to use Virtio

- Get kvm version >= 60
- Get Linux kernel with virtio drivers for the guest
    - Get Kernel >= 2.6.25 and activate (modules should also work, but take care of initramdisk)
    - CONFIG_VIRTIO_PCI=y (Virtualization -> PCI driver for virtio devices)
    - CONFIG_VIRTIO_BALLOON=y (Virtualization -> Virtio balloon driver)
    - CONFIG_VIRTIO_BLK=y (Device Drivers -> Block -> Virtio block driver)
    - CONFIG_VIRTIO_NET=y (Device Drivers -> Network device support -> Virtio network driver)
    - CONFIG_VIRTIO=y (automatically selected)
    - CONFIG_VIRTIO_RING=y (automatically selected)
    - you can safely disable SATA/SCSI and also all other nic drivers if you only use VIRTIO (disk/nic)
- As an alternative one can use a standard guest kernel for the guest > 2.6.18 and make use sync backward compatibility option
Backport and instructions can be found in kvm-guest-drivers-linux.git
- Use virtio-net-pci device for the network devices (or model=virtio for old -net..-net syntax) and if=virtio for disk
    - Example
```
x86_64-softmmu/qemu-system-x86_64 -boot c -drive file=/images/xpbase.qcow2,if=virtio -m 384 -netdev type=tap,script=/etc/kvm/qemu-ifup,id=net0 -device virtio-net-pci,netdev=net0
```
- -hd[ab] for disk won't work, use -drive
- Disk will show up as /dev/vd[a-z][1-9], if you migrate you need to change "root=" in Lilo/GRUB config
- At the moment the kernel modules are automatically loaded in the guest but the interface should be started manually (dhclient/ifconfig)
- Currently performance is much better when using a host kernel configured with CONFIG_HIGH_RES_TIMERS. Another option is use HPET/RTC and -clock= qemu option.
- Expected performance
    - Performance varies from host to host, kernel to kernel
    - On my laptop I measured 1.1Gbps rx throughput using 2.6.23, 850Mbps tx.
    - Ping latency is 300-500 usec