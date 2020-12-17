# Volume

On-disk files in a container are ephemeral, which presents some problems for non-trivial applications when running in containers. One problem is the loss of files when a container crashes. The kubelet restarts the container but with a clean state. A second problem occurs when sharing files between containers running together in a `pod`. The Kubernetes volume abstraction solves both of these problems. Familiarity with Pods is suggested.

## Background

Docker has a concept of volumes, though it is somewhat looser and less managed. A Docker volume is a directory on disk or in another container. Docker provides volume drivers, but the functionality is somewhat limited.

Kubernetes supports many types of volumes. A Pod can use any number of volume types simultaneously. Ephemeral volume types have a lifetime of a pod, but persistent volumes exist beyond the lifetime of a pod. Consequently, a volume outlives any containers that run within the pod, and data is preserved across container restarts. When a pod ceases to exist, the volume is destroyed.

At its core, a volume is just a directory, possiblely with some data in it, which is accessible to the containers in a pod. How that directory comes to be, the medium that backs it, and the contents of it are determined by the particular volume type used.

To use a volume, specify the volumes to provide for the Pod in `.spec.volumes` and declare where to mount those volumes into containers in `.spec.container[*].volumeMounts`. A process in a container sees a filesystem view composed from their Docker image and volumes. The Docker image is at the root of the filesystem hierarchy. Volumes mount at the specified paths within the image. Volumes can not mount onto other volumes or have hard links to other volumes. Each Container in the Pods's configuration must independently specify where to mount each volume.

## Types of Volumes

### cephfs

A `cephfs` volume allows an existing CephFS volume to be mounted into your Pod. Unlike `emptyDir`, which is erased when a pod is removed, the contents of a `cephfs` volume are preserved and the volume is merely unmounted. This means that a `cephfs` volume can be pre-populated with data, and that data can be shared between pods. The `cephfs` volume can be mounted by multiple writes simultaneously.

>**Note**: You must have your own Ceph server running with the share exported before you can use it.

### configMap

A ConfigMap provides a way to inject configuration data into pods. The data stored in a ConfigMap can be referenced in a volume of type `configMap` and then consumed by containerized applications running in a Pod.

When referencing a ConfigMap, you provide the name of the ConfigMap in the volume. You can customize the path to use for a specific entry in the ConfigMap. The following configuration shows how to mount the `log-config` ConfigMap onto a Pod called `configmap-pod`:

```yaml
apiVersion: v1
kind: Pod
metadata:
    name: configmap-pod
spec:
    containers:
        - image: test
            volumeMounts:
                - name: config-vol
                  mountPath: /etc/config
    volumes:
        - name: config-vol
          configMap:
            name: log-config
            items:
                - key: log_level
                  path: log_level
```

The `log-config` ConfigMap is mounted as a volume, and all contents stored in its `log_level` entry are mounted into the Pod at path `/etc/config/log_level`. Note that this path is derived from the volume's `mountPath` and the `paht` keyed with `log_level`.

>**Note:**
> - You must create a ConfigMap before you can use it.
> - A container using a ConfigMap as a `subPath` volume mount will not receive ConfigMap updates.
> - Text data is exposed as files using the UTF-8 character encoding. For other character encoding, use `binaryData`.

### emptyDir

An `emptyDir` volume is first created when a Pod is assigned to a node, and exists as long as that Pod is running on that node. As the name says, the `emptyDir` volume is initially empty. All containers in the Pod can read and write the same files in the `emptyDir` volume, though that volume can be mounted at the same or different paths in each container. When a Pod is removed from a node for any reason, the data in the `emptyDir` is deleted permanently.

>**Note:** A container crashing does not remove a Pod from a node. The data in an emptyDir volume is safe across container crashes.

Some uses for an `emptyDir` are:

- scratch space, such as for a disk-based merge sort
- checkpointing a long computation for recovery from crashes
- holding files that a content-manager container fetches while a webserver container serves the data

Depending on your environment, `emptyDir` volumes are stored on whatever medium that backs the node such as disk or SSD, or network storage. However, if you set the `emptyDir.medium` field to `"memory"`, Kubernetes mounts a tmpfs(Ram-backed filesystem) for you instead. While tmpfs is very fast, be aware that unlike disks, tmpfs is cleared on node reboot and any files you write count against your container's memory limit.

> **Note:** If the SizeMemoryBackedVolume feature gate is enabled, you can specify a size for memory backed volumes. If no size if specified, memory backed volumes are sized to 50% of the memory on a Linux host.

```yaml
apiVersion: v1
kind: Pod
metadata:
    name: test-pd
spec:
    containers:
    - image: k8s.gcr.io/test-webserver
      name: test-container
      volumeMounts:
      - mountPath: /cache
        name: cache-volume
    volumes:
    - name: cache-volume
      emptyDir: {}   
```

### glusterfs

A `glusterfs` volume allows a Glusterfs (an open source networked filesystem) volume to be mounted into your Pod. Unlike `emptyDir`, which is erased when a Pod is removed, the contents of a `glusterfs` volume are preserved and the volume is merely unmounted. This means that a glusterfs volume can be prepopulated with data, and that data can be shared between pods. GlusterFS can be mounted by multiple writes simultaneously.

### hostPath

A `hostPath` volume mounts a file or dirctory from the host node's filesystem into your Pod. This is not something that most Pods will need, but it offers a powerful escape hatch for some applications.

For example, some uses for a `hostPath` are:

- running a container that needs access to Docker internal; use a `hostPath` of `/var/lib/docker`
- running cAdvisor in a container; use a `hostPath` of `/sys`
- allowing a Pod to specify whether a given `hostPath` should exist prior to the Pod running, whether it should be created, and what it should exist as

In addition to the required `path` property, you can optionally specify a `type` for a `hostPath` volume.

The supported values for field `type` are:

|Value|behavior|
|-|-|
||Empty string(default) is for backward compatibility, which means that no checks will be performed before mounting the hostPath volume|
|`DirectoryOrCreate`|If nothing exists at the given path, an empty dirctory will be created there as needed with permission set to 0755, having the same group and ownership with Kubelet|
|`Directory`|A directory must exist at the given path|
|`FileOrCreate`|If nothing exists at the given path, an empty file will be created there as needed with permission set to 0644, having the same group and ownership with Kubelet|
|`File`|A file must exist at the given path|
|`Socket`|A UNIX socket must exist at the given path|
|`CharDevice`|A character device must exist at the given path|
|`BlockDevice`|A block device must exist at the given path|

Watch out when using this type of volume, because:

- Pods with identical configuration(such as created from a PodTemplate) may behave differently on different nodes due to different files on the nodes
- The files or directories created on the underlying hosts are only writable by root. You either need to run your process as root in a privileged Container or modify the file permissions on the host to be able to write to a `hostPath` volume

```yaml
apiVersion: v1
kind: Pod
metadata:
    name: test-pd
spec:
    containers:
    - image: k8s.gcr.io/test-webserver
      name: test-container
      volumeMounts:
      - mountPath: /test-pd
        name: test-volume
    volumes:
    - name: test-volume
      hostPath:
        path: /data
        type: Directory
```

>**Caution:**  The FileOrCreate mode does not create the parent directory of the file. If the parent directory of the mounted file does not exist, the pod fails to start. To ensure that this mode works, you can try to mount directories and files separately, as shown in the FileOrCreateconfiguration.

### local

A `local` volume represents a mounted local storage device such as a disk, partition or directory.

Local volumes can only be used as a statically created PersistentVolume. Dynamic provisioning is not supported.

Compared to `hostPath` volumes, `local` volumes are used in a durable and portable manner without manually scheduling pods to nodes. The system is aware of the volume's node constraints by looking at the node affinity on the PersistentVolume.

However, `local` volumes are subject to the availability of the underlying node and are not suitable for all applications. If a node becomes unhealthy, then the `local` volume becomes inaccessible by the pod. The pod using this volume is unable to run. Applications using `local` volumes must be able to tolerate this reduced availability, as well as potential data loss, depending on the durability characteristics of the underlying disk.

The following examples shows a PersistentVolume using a `local` volume and `nodeAffinity`:

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
    name: example-pv
spec:
    capacity:
        storage: 100Gi
    volumeMode: Filesystem
    accessModes:
    - ReadWriteOnce
    persistentVolumeReclaimPolicy: Delete
    storageClassName: local-storage
    local:
        path: /mnt/disks/ssd1
    nodeAffinity:
        required:
            nodeSelectorTerms:
            - matchExpressions:
                - key: kubernetes.io/hostname
                  operator: In
                  values: 
                  - example-node
```

You must set a PersistentVolume `nodeAffinity` when using `local` volumes. The Kubernetes scheduler uses the PersistentVolume `nodeAffinity` to schedule these Pods to the correct node.

PersistentVolume `volumeMode` can be set to "Block" (instead of the default value "Filesystem") to expose the local volume as a raw block device.

When using local volumes, it is recommended to create a StorageClass with `volumeBindingMode` set to `WaitForFirstConsumer`. For more details, see the local StorageClass example. Delaying volume binding ensures that the PersistentVolumeClaim binding decision will also be evaluated with any other node constraints the Pod may have, such as node resource requirement, node selectors, Pod affinity, and Pod anti-affinity.

An external static provisioner can be run separately for improved management of the local volume lifecycle. Note that this provisioner does not support dynamic provisioning yet. For an example on how to run an external local provisioner, see the local volume provisioner user guide.

> **Note:** The local PersistentVolume requires manual cleanup and deletion by the user if the external static provisioner is not used to manage the volume lifecycle.

### nfs

An `nfs` volume allows an existing NFS(Network File System) share to be mounted into a Pod. Unlike `emptyDir`, which is erased when a Pod is removed, the contents of an `nfs` volume are preserved and the volume is merely unmounted. This means that an NFS volume can be prepopulated with data, and that data can be shared between pods. NFS can be mounted by multiple writers simultaneously

### persistentVolumeClaim

A `persistentVolumeClaim` volume is used to mount a PersistentVolume into a Pod. PersistentVolumeClaims are a way for users to "claim" durable storage(such as a GCE PersistentDisk or an iSCSI volume) without knowing the details of the particular cloud environment

### projected

A `projected` volume maps serveral volume resources into the same directory. Currently, the following types of volume sources can be projected:

- secret
- downwardAPI
- configMap
- `serviceAccountToken`

### rbd

### secret

### vsphereVolume

## Using subPath

## Resources

The storage media(such as Disk or SSD) of an `emptyDir` volume is determined by the medium of the filesystem holding the kubelet root dir(typically `/var/lib/kubelet`). There is no limit on how much space an `emptyDir` or `hostPath` volume can consume, and no isolation between containers or between pods.

## Out-of-tree volume plugins

The out-of-tree volume plugins include Container Storage Interface(CSI) and FlexVolume. These plugins enable storage vendors to create custom storage plugins without adding their plugin source code to the Kubernetes repository.

Previously, all volume plugins were "in-tree". The "in-tree" plugins were built, linked, compiled, and shipped with the core Kubernetes binaries. This meant that adding a new storage system to Kubernetes(a volume plugin) required checking code into the core Kubernetes code repository. 

Both CSI and FlexVolume allow volume plugins to be developed independent of the Kubernetes code base, and deployed(installed) on Kubernetes clusters as extensions.

### csi

Container Storage Interface defines a standard interface for container orchestration systems(like Kubernetes) to expose arbitrary storage systems to their container workloads.

>**Note**: CSI drivers may not be compatible across all Kubernetes releases. Please check the specific CSI driver's documentation for supported deployments steps for each Kubernetes release and a compatibility matrix.

Once a CSI compatible volume driver is deployed on a Kubernetes cluster, users may use the `csi` volume type to attach or mount the volumes exposed by the CSI driver.

A `csi` volume can be used in a Pod in three different ways:

- through a reference to a PersistentVolumeClaim
- with a generic ephemeral volume
- with a CSI ephemeral volume if the driver supports that

The following fields are available to storage administrators to configure a CSI persistent volume:

- `driver`: A string value that specifies the name of the volume driver to use. This value must correspond to the value returned in the `GetPluginInfoResponse` by the CSI driver as defined in the CSI spec. It is used by Kubernetes to identity which PV objects belong to the CSI driver.
- `volumeHandle`
- `readOnly`
- `fsType`
- `volumeAttributes`
- `controllerPublishSecretRef`
- `nodeStageSecretRef`
- `nodePublishSecretRef`

#### CSI raw block volume support

Vendors with external CSI drivers can implement raw block volume support in Kubernetes workloads

### flexVolume

## Mount propagation