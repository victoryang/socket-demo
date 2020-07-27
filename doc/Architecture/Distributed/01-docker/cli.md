# Cli

[commandline](https://docs.docker.com/engine/reference/commandline)

## Commit

```
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]
```

It can be useful to commit a container’s file changes or settings into a new image. 
This allows you to debug a container by running an interactive shell, or to export a working dataset to another server.
Generally, it is better to use Dockerfiles to manage your images in a documented and maintainable way.

The commit operation will not include any data contained in volumes mounted inside the container.

**By default, the container being committed and its processes will be paused while the image is committed. This reduces the likelihood of encountering data corruption during the process of creating the commit.**