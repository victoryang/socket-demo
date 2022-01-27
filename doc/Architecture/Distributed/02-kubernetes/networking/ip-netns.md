# ip netns

[ip netns](https://blog.csdn.net/supahero/article/details/100606953)

```bash
# ip netns add netns0
# ip netns add netns1

# ip link add name vnet0 type veth pair name vnet1

# ip link set vnet0 netns netns0
# ip link set vnet1 netns netns1

# ip netns exec netns ip link set vnet0 up
# ip netns exec netns ip link set vnet1 up

# ip netns exec netns ip a add 192.168.0.2 dev vnet0
# ip netns exec netns ip a add 192.168.0.3 dev vnet1

# ip netns exec netns ip route add 192.168.0.3 dev vnet0
# ip netns exec netns ip route add 192.168.0.2 dev vnet1

# ip netns exec netns0 ping 192.168.0.3
# ip netns exec netns1 ping 192.168.0.2
```