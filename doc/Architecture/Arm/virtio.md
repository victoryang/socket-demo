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