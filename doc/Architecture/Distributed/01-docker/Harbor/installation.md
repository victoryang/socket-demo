# Installation

[Download](https://goharbor.io/docs/2.0.0/install-config/download-installer/)

[Installation](https://goharbor.io/docs/2.0.0/install-config/)

[registry-1.docker.io/v2/](https://cloud.tencent.com/developer/article/1368054)

[docker registry-1.docker.io](https://blog.csdn.net/laogouhuli/article/details/92987525)

[cloudflare.docker.com](https://blog.csdn.net/qq_35868412/article/details/102881912)


## 官方文档入口 
https://goharbor.io/docs/2.0.0/install-config

Harbor 支持 Docker Compose 和 Kubernetes 部署，本文重点偏前者

## Installation Prerequisite

- Docker Engine
- Docker Compose
- OpenSSL

## Download the Harbor Installer
可选取 Online installer 或 Offline installer

下载地址为 github releases page

https://github.com/goharbor/harbor/releases

注：可以通过asc 文件校验下载内容

## Security
可以配置使用 HTTP 协议访问 Harbor，官方建议在生产环境应该使用 HTTPS。

来自 Harbor 外部的，或者 Harbor 内部组件之间的请求，都可以通过 HTTPS 访问

## Configure Harbor YML File
Harbor 的系统参数都在 harbor.yml 中配置。修改此文件中的配置，将影响后续的安装过程。

具体配置详见 https://github.com/goharbor/harbor/blob/master/docs/install-config/configure-yml-file.md

## Run the Installer Script
运行 install.sh 将开始安装过程，如前面所说，安装将依赖 harbor.yml 的配置内容。

并且可分别通过设置如下参数来决定是否也同时安装对应的组件

--with-notary

--with-clair

--with-chartmuseum

具体参数详见 https://goharbor.io/docs/2.0.0/install-config/run-installer-script/

## TroubleShooting
详见 https://goharbor.io/docs/2.0.0/install-config/troubleshoot-installation/