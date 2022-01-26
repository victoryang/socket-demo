# CNI-Plugins

[libcni](https://github.com/containernetworking/cni)

[plugins](https://github.com/containernetworking/plugins)

[cni plugin issue](https://blog.csdn.net/github_35614077/article/details/98213572)

## Basic CNI plugins

### main

interface-creating

- `bridge`: Creates a bridge, adds the host and the container to it
- `ipvlan`: Adds an ipvlan interface in the container
- `loopback`: Set the state of loopback interface to up
- `macvlan`: Creates a new MAC address, forwards all traffic to that to the container
- `ptp`: Creates a veth pair
- `vlan`: Allocates a vlan device
- `host-device`: Move an already-existing device into a container

### IPAM

IP address allocation

- `dhcp`: Runs a daemon on the host to make DHCP requests on behalf of the container
- `host-local`: Maintains a local database of allocated IPs
- `static`: allocate a static IPv4/IPv6 addresses to container and it's useful in debugging purpose

### Meta

other plugins

- `tuning`: Tweaks sysctl parameters of an exisiting interface
- `portmap`: An iptables-based portmapping plugin. Maps ports from the host's address space to the container
- `bandwidth`: Allow bandwidth-limiting through use of traffic control tbf(ingress/egress)
- `sbr`: A plugin that configures source based routing for an interface(from which it is chained)
- `firewall`: A firewal plugin which uses iptables or firewalld to add rules to allow traffic to/from the container.
