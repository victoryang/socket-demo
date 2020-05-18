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

