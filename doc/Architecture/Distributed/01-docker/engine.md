# Docker Engine

## Docker Edition

- Docker Engine Community
- Docker Engine Enterprise
- Docker Enterprise

### Docker Engine Community
Docker Engine Community is ideal for developers and small teams to get started with Docker and expirementing with container-based apps. Docker Engine Community has three types of update channels, stable, test and nightly.

- **Stable** gives you latest releases for general availability
- **Test** gives pre-releases that are ready for testing before general availability
- **Nightly** gives you latest builds of work in progress for the next major release

## Docker Overview

### Docker Engine
***Docker Engine*** is a client-server application with these major components:

- A server which is a type of long-running program called a daemon process(the `dockerd` command)
- A REST API which specifies interfaces that programs can use to talk to the daemon and instruct it what to do
- A command line interface (CLI) client (the `docker` command)

The CLI uses the Docker REST API to control or interact with Dokcer daemon through scripting or direct CLI command. Many other Docker applications use the underlying API and CLI.

The daemon creates and manages *Docker Objects*, such as images, containers, networks and volumes.

### Docker Architecture
Docker uses a client-server architecture. The Docker *client* talks to *daemon*, which does the heavy lifting of building, running and distributing your Docker containers. The Docker client and daemon can *run* on the same system, or you can connect a Docker client to a remote Docker Daemon. The docker client and daemon communicate using a REST API, over UNIX socket or a network interface.

#### The Docker Daemon
The Docker daemon(`dockerd`)listens for Docker API requests and manages Docker objects such as images, containers, networks and volumes. A daemon can also communicate with other daemons to manage docker services.

#### The Docker Client
The Docker client(`docker`) is the primary way that many Docker users interact with Docker. When you use command such as `docker run`, the client send these commands to `dockerd`, which carries them out.

#### The Docker Registries
A docker *registry* stores Docker images. Docker Hub is a public registry that anyone can use, and Docker is configured to look for images on Docker Hub by default. You can even run your own private registry. If you use Docker Datacenter(DDC), it includes Docker Trusted Registry(DTR).

When you use the `docker pull` or `docker run` commands, the required images are pulled from your configured registry. When you use the `docker push` command, your image is pushed to your configured registry.

#### Docker Objects

##### IMAGES

##### CONATAINERS

##### SERVICES