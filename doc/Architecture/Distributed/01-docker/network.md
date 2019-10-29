# Docker network
- iptables on linux server
- routing rules on windows server
- forms and encapsulates packet
- handle encryption

## Network drivers
- bridge (default)
    - The default network driver. Bridge networks are usually used when your applications run in standalone containers that need to communicate.
- host
    - For standalone containers, remove network isolation between the container and the Docker host, and use the host's network. host is only available for swarm services on 17.06 and higher.
- overlay
    - Overlay networks connect multiple docker daemon together and enable swarm services to communicate with each other. You can also use overlay networks to facilitate communication between a swarm service and a standalone container, or between two standalone containers on different Docker daemons. This strategy removes the need to do OS-level routing between these containers.
- macvlan
    - Macvlan networks allow you to assign a MAC address to a container, making it appear as a physical device on your network. The Docker Daemon routes traffic to containers by their MAC address.
- none
    - disable networking
- Networks plugins

### Network driver summary
- User-defined bridge networks
    - are best when you need multiple containers to communicate on the same Docker host.
- Host networks
    - are best when the network stack should not be isolated from the docker host, but you want other aspects of the container to be isolated.
- Overlay networks
    - are best when you need containers running on different docker hosts to communicate, or when multiple applications work together using swarm services.
- Macvlan networks
    - are best when you are migrating from a VM setup or need your containers to look like physical hosts on your network, each with a unique MAC address.
- Third-party network plugins
    - allow you to integrate Docker with specialized network stack.

## Docker EE networking features
- HTTP routing mesh
- Session stickness

## Docker and iptables
On linux, Docker manipulates iptables rules to provide network isolation. This is an implementation detail, and you should not modify the rules Dockr inserts into your iptables policies.