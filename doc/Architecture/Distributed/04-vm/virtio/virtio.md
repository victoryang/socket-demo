# Virtio: An I/O virtualization framework for Linux

https://developer.ibm.com/articles/l-virtio/

In a nutshell, `virtio` is an abstraction layer over devices in a paravirtualized hypervisor.

Linux is the hypervisor playground. As my article showed, Linux offers a variety of hypervisor solutions with different attributes and advantages. Examples include the Kernel-based Virtual Machine(KVM), lguest, and User-mode Linux. Having these different hypervisor solutions on Linux can tax the operating system based on their independent needs. One of the taxes is virtualization of devices. Rather than have a variety of device emulation mechanisms (for network, block, and other drivers), `virtio` provides a common front end for these device emulations to standardize the interface and increase the reuse of code across the platforms.