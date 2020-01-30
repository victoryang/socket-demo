# Nginx

[高并发架构一般思路](https://www.cnblogs.com/gdsblog/p/7128497.html)

[nginx参考](http://tengine.taobao.org/book/chapter_02.html)

[nginx 1](https://blog.csdn.net/qq_29677867/article/details/90112120)

[nginx phase](https://www.centos.bz/2018/12/nginx%E8%AF%B7%E6%B1%82%E5%A4%84%E7%90%86%E6%B5%81%E7%A8%8B%E4%BD%A0%E4%BA%86%E8%A7%A3%E5%90%97%EF%BC%9F/)

[nginx一般介绍](https://www.centos.bz/2017/11/openresty%E6%9C%80%E4%BD%B3%E6%A1%88%E4%BE%8B-%E7%AC%AC1%E7%AF%87%EF%BC%9Anginx%E4%BB%8B%E7%BB%8D/)


## Nginx 应用
<img src="nginx_function.jpeg">

## Nginx 安装

### Download

```
wget -c https://nginx.org/download/nginx-1.16.1.tar.gz

# see http://nginx.org/en/docs/configure.html
./configure

make clean & make & make install
```

## Nginx 架构

### Nginx 进程模型
- 前台单进程
- 后台多进程
  - 单Master进程
  - 多Worker进程

### 多worker进程模型
<image src="nginx.svg"></image>

<img src="nginx_arch.jpeg">

<img src="nginx_arch2.jpeg">

## Nginx模块化和handle

<img src="nginx_modules.jpg">

### nginx_conf 模块

<img src="nginx_http_conf.jpg">

- main 全局设置
- server 主机设置
- upstream 负载均衡，后端服务器
- location url匹配位置

location --> server --> main

#### main

- user
- worker processors
- error_log
- pid
- worker_rlimit_nofile
- event module
- http module
- server module
- location module

### nginx_http 模块

<img src="nginx_handler.jpeg">

<img src="nginx_http_handle.jpg">

## Caching Process

<img src="nginx_cache.jpeg">

## Nginx HA

<img src="nginx_ha.jpeg">