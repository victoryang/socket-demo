# drm


[weston 架构](https://www.cnblogs.com/arnoldlu/p/18091352)

[drm](https://doc.embedfire.com/linux/rk356x/linux_base/zh/latest/linux_app/drm/drm.html)

[drm-kms](https://events.static.linuxfound.org/sites/events/files/slides/brezillon-drm-kms.pdf)

[drm](https://www.cnblogs.com/zyly/p/17775867.html)

[开源mali安装](https://cloud.tencent.com/developer/article/1571948)

[andorid surfaceflinger](https://juejin.cn/post/7047745117267951646)

[gpu](https://zhuanlan.zhihu.com/p/1930594532157813197)

[gl vs egl vs glx ](https://zhuanlan.zhihu.com/p/400553896)

[wayland egl](https://blog.csdn.net/czg13548930186/article/details/131103472)

[wayland](https://wayland-book.com/introduction/high-level-design.html)

[swrast](https://www.oryoy.com/news/ubuntu-xi-tong-xia-de-swrast-jie-mi-tu-xing-xuan-ran-de-di-ceng-ao-mi.html)

[mesa](https://zhuanlan.zhihu.com/p/432215702)

[arm G610 GPU](https://www.cnblogs.com/arnoldlu/p/18175082)

[wayland and weston](https://zhuanlan.zhihu.com/p/690561669)

[wayland and weston](https://cloud.tencent.com/developer/article/1445734)

[weston](https://www.cnblogs.com/arnoldlu/p/18091352)

[opengl](https://zhuanlan.zhihu.com/p/1930594532157813197)

[android surface](https://juejin.cn/post/7047745117267951646)

## Linux graphic stack

### compositing

The compositor is a system service that receives each application's output buffers and draws them to an on-screen image.

A Wayland **surface** represents an application window; it is the application's handle to display its output and receive input events from the compositor. Attached to the surface is a Wayland buffer that contains the displayable pixel data plus color-format and color-format and size information. The pixel data is in the output buffer that the client application has rendered to.

The compositor maintains a list of all of the Wayland surfaces that represent application windows.

Wayland provides a protocol extension to share buffer objects via a Linux dma-buf, which represents a memory buffer that is shareable among hardware devices, drivers, and user-space programs. An application renders its scene graph via the Mesa interfaces using hardware acceleration as described in part 1, but, instead of transferring a reference to shared memory, the application sends a dma-buf object that references the buffer object while it is still located in graphics memory. The Wayland compositor uses the stored pixel data without ever reading it over the hardware bus.

### pixels to the monitor

DRM's mode-setting code controls all aspects of reading pixel data from graphics memory and sending it to an output device.

The minimum stages necessary are the framebuffer, plane, CRTC, encoder, and connector, each of which is described below.

The pipeline's first stage is the DRM framebuffer. It is the buffer object that stores the compositor's on-screen image, plus information about the image's color format and size.

Fetching the pixel data is called scanout, and the pixel data's buffer object is called the scanout buffer. The number of scanout buffers per framebuffer depends on the framebuffer's color format. Many formats, such as the common RGB-based ones, store all pixel data in a single buffer. With other formats, such as YUV-based ones, the pixel data might need to be split up into multiple buffers.

In DRM terminology, this is called a plane. It sets the scanout buffer's position, orientation, and scaling factors. Depending on the hardware, there can be multiple active planes using different framebuffers. All active planes feed their pixel output into the pipeline's third stage, which is called the cathode-ray tube controller (CRTC) for historical reasons.

The CRTC controls everything related to display-mode settings. The DRM driver programs the CRTC hardware with a display mode and connects it with all of its active planes and outputs. There can also be multiple CRTCs with different settings programmed to them. The exact configuration is only limited by hardware features.

 According to the programmed display mode and each plane's location, the CRTC hardware fetches pixel data from the planes, blends overlapping planes where necessary, and forwards the result to its outputs.

 Outputs are represented by encoders and connectors. As its name suggests, the encoder is the hardware component that encodes pixel data for an output. An encoder is associated with a specific connector, which represents the physical connection to an output device, such as HDMI or VGA ports with a connected monitor. The connector also provides information on the output device's supported display modes, physical resolution, color space, and the like. Outputs on the same CRTC mirror the CRTC's screen on different output devices.

 ### pipeline setup

 Deciding on policies for connecting and configuring the individual stages of the mode-setting pipeline is not the DRM driver's job. As part of its initial setup, the compositor opens the device file under /dev/dri, such as /dev/dri/card1, and invokes the respective ioctl() calls to program the display pipeline. It also fetches the available display modes from a connector and picks a suitable one.

 After the compositor has finished rendering the first on-screen image, it programs the mode-setting pipeline for the first time. To do so, it creates a framebuffer for the on-screen image's buffer object and attaches the framebuffer to a plane. It then sets the display mode for its on-screen buffer on the CRTC, connects all of the pipeline stages, from framebuffer to connector, and enables the display.

 It would first program the display mode in the CRTC, then upload all buffer objects into graphics memory, then set up the framebuffers and planes for scanout, and finally enable the encoders and connectors. 