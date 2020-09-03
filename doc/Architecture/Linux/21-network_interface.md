# Network Interface

[ip tools](https://www.runoob.com/linux/linux-comm-ip.html)

## Promiscuous Mode

promiscuous mode allows a network device to intercept and read each network packet that arrives in its entirety.

```bash
# 开启网卡的混合模式
ip link set eth0 promisc on
```