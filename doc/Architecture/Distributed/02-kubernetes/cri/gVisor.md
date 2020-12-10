# gVisor

[gVisor](https://github.com/google/gvisor)

[gitee](https://gitee.com/skymysky/gVisor/)

## What is gVisor

**gVisor** is an application kernel, written in Go, that implements a substantial portion of the Linux system surface. It includes an OCI runtime called runsc that provides an isolation boundary between the application and the host kernel. The `runsc` runtime integrates with Docker and Kubernetes, making it simple to run sandboxed containers.

## Why does gVisor exist?

Containers are not a sandbox. While containers have revolutionized how we develop, package, and deploy applications, using them to run untrusted or potentially mailicious code without additional isolation is not a good idea. While using a single, shared kernel allows for efficiency and performance gains, it also means that container escape is possible with a single vulnerability.

gVisor is an application kernel for containers. It limits the host kernel surface accessible to the application while still giving the application access to all the features it expects. Unlike most kernels, gVisor does not assume or require a fixed set of physical resources; instead, it leverages existing host kernel functionality and runs as a normal process. In other words, gVisor implements Linux by way of Linux.

gVisor should not be confused with technologies and tools to harden containers against external threads, provide additional integrity checks, or limit the scope of access for a service. One should always be careful about what data is made available to a container.