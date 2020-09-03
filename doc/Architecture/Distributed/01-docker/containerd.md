# Containerd

[official website](https://containerd.io/docs/getting-started/)

[kubernetes arch](http://dockone.io/article/9149)

[build from source](https://github.com/containerd/containerd/blob/master/BUILDING.md)

## Architecture

<img src="containerd_architecture.png">

## Container and Task

A **Container** is a metadata object that resources are allocated and attached to

A **Task** is a live, running process on the system

Tasks should be deleted after each run while a container can be used, updated, and queried multiple times.

The new task that we just created is actually a running process on your system.

### Task Wait and Start

```
If you are familiar with the OCI runtime actions, the task is currently in the “created” state. 

This means that the namespaces, root filesystem, and various container level settings have been initialized but the user defined process, in this example “redis-server”, has not been started. 

This gives users a chance to setup network interfaces or attach different tools to monitor the container. 

containerd also takes this opportunity to monitor your container as well. 

Waiting on things like the container’s exit status and cgroup metrics are setup at this point.
```

It is essential to wait for the task to finish so that we can close our example and cleanup the resources that we created. 

You always want to make sure you Wait before calling Start on a task. 