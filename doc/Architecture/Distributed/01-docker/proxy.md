# Proxy

https://blog.csdn.net/mtj66/article/details/90700034

## proxy

```bash
vi /etc/systemd/system/docker.service.d/http-proxy.conf

[Service]
Environment="HTTP_PROXY=http://192.168.47.68:3128"
Environment="HTTPS_PROXY=http://192.168.47.68:3128"
```