# Gitlab

[official doc](https://gitlab.com/help/user/index.md)

[gitlab features](https://about.gitlab.com/features/)

[installation](https://docs.gitlab.com/ce/install/installation.html)

[gitlab ce source](https://gitlab.com/gitlab-org/gitlab-foss)

## Installation

- Omnibus GitLab package
- Kubernetes Helm
- Docker
- Source

### Omnibus GitLab package

[guide](https://about.gitlab.com/install/#ubuntu)

## Git

- 工作区
    - 当前工作路径
- 暂存区
    - stage 或 index，存放在 .git 目录下的index文件中
- 版本库
    - .git 目录，git的版本库

<img src="git_hierarchy.jpg">

HEAD 指向当前分支
objects Git的对象库，位于.git/objects，包含创建的对象和内容。

当执行 git add 时，暂存区的目录树被更新，同时工作区修改的内容吸入到对象库中的一个新对象里，而该对象的ID被记录在暂存区的文件索引里。

git commit, 暂存区的目录树被写入到版本库中，master 分支做出相应的更新。

git reset HEAD，暂存区的目录会被重写，被master分支指向的目录树所替换，但是工作区不受影响。

## Arch

[基础](https://juejin.im/post/5cf67d355188255508118def)

- repository: 代码库，可以试硬盘或NFS
- Nginx： web 入口
- 数据库：
    - repository 中的数据（元数据，issue，mr）
    - 可以登录web的用户（权限）
- Redis：缓存，负责分发任务
- sidekiq：后台任务， 主要负责发送电子邮件
- Unicorn：Gitlab自身的服务器，包含了Gitlab主进程，负责处理快速/一般任务，与redis一起工作
    - 通过检查存储在Redis中的用户会话来检查权限
    - 为 sidekiq 制作任务
    - 从仓库(warehouse)取东西或那里移动
- gitlab-shell 用于SSH交互。gitlab-shell通过Redis与sidekiq进行通信，并直接通过tcp访问Unicorn
- gitaly 后台服务，专门负责访问磁盘以高效处理git操作
- gitlab-workhorse 反向代理服务器，处理与Rails无关的请求(js,css等)，处理Git push请求，处理到Rails的链接。（修改由 Rails 发送的响应或发送给 Rails 的请求，管理 Rails 的长期 WebSocket 连接等）


### Workhouse

[workhouse](https://blog.csdn.net/weixin_34326558/article/details/91479012?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase)

### Unicorn

[unicorn](https://blog.csdn.net/weixin_34294649/article/details/91475976)

### Gitlab-shell

[gitlab-shell](https://blog.csdn.net/weixin_34025151/article/details/91475980)

### Ruby on rails

[ruby on rails](https://ihower.tw/rails/index-cn.html)