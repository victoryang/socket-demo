# Dockershim

kubelet.syncPod --> kuberuntime_manager.SyncPod

## kubelet.syncPod

kublete.Run --> kubelet.syncLoop --> kubelet.syncLoopIteration --> kubelet.podworkers.UpdatePod  --> kubelet.syncPod

```go
// If the pod should not be running, we request the pod's containers be stopped. This is not the same
// as termination (we want to stop the pod, but potentially restart it later if soft admission allows
// it later). Set the status and phase appropriately
runnable := kl.canRunPod(pod)

// Record the time it takes for the pod to become running.
existingStatus, ok := kl.statusManager.GetPodStatus(pod.UID)

// If pod has already been terminated then we need not create
// or update the pod's cgroup
if !kl.podWorkers.IsPodTerminationRequested(pod.UID){

}

// Create Mirror Pod for Static Pod if it doesn't already exist
if kubetypes.IsStaticPod(pod) {

}

// Make data directories for the pod
kl.makePodDataDirs(pod)

// Volume manager will not mount volumes for terminating pods
if !kl.podWorkers.IsPodTerminationRequested(pod.UID) {

}

// Fetch the pull secrets for the pod
pullSecrets := kl.getPullSecretsForPod(pod)

// Call the container runtime's SyncPod callback
result := kl.containerRuntime.SyncPod(pod, podStatus, pullSecrets, kl.backOff)

kl.reasonCache.Update(pod.UID, result)
```

## kuberuntime_manager.SyncPod

genericRuntimeManager --> dockershim --> docker

pkg/kubelet/kuberuntime/kuberuntime_manager.go

m.SyncPod

SyncPod syncs the running pod into the desired pod by executing following steps:

1. Compute sandbox and container changes.
2. Kill pod sandbox if necessary.
3. Kill any containers that should not be running.
4. Create sandbox if necessary.
5. Create ephemeral containers.
6. Create init containers.
7. Create normal containers.

```go
// Step 1: Compute sandbox and container changes.
podContainerChanges := m.computePodActions(pod, podStatus)

// Step 2: Kill the pod if the sandbox has changed.
if podContainerChanges.KillPod {
    killResult := m.killPodWithSyncResult(pod, kubecontainer.ConvertPodStatusToRunningPod(m.runtimeName, podStatus), nil)
} else {
    // Step 3: kill any running containers in this pod which are not to keep.
    err := m.killContainer(pod, containerID, containerInfo.name, containerInfo.message, containerInfo.reason, nil)
}

// Step 4: Create a sandbox for the pod if necessary.
if podContainerChanges.CreateSandbox {
    podSandboxID, msg, err = m.createPodSandbox(pod, podContainerChanges.Attempt)

    podIPs = m.determinePodSandboxIPs(pod.Namespace, pod.Name, podSandboxStatus)
}

// Get podSandboxConfig for containers to start.
podSandboxConfig, err := m.generatePodSandboxConfig(pod, podContainerChanges.Attempt)

// Helper containing boilerplate common to starting all types of containers.
start := func(typeName, metricLabel string, spec *startSpec) error {
    isInBackOff, msg, err := m.doBackOff(pod, spec.container, podStatus, backOff)

    m.startContainer(podSandboxID, podSandboxConfig, spec, pod, podStatus, pullSecrets, podIP, podIPs)
}

// Step 5: start ephemeral containers
for _, idx := range podContainerChanges.EphemeralContainersToStart {
    start("ephemeral container", metrics.EphemeralContainer, ephemeralContainerStartSpec(&pod.Spec.EphemeralContainers[idx]))
}

// Step 6: start the init container.
start("init container", metrics.InitContainer, containerStartSpec(container))

// Step 7: start containers in podContainerChanges.ContainersToStart.
start("container", metrics.Container, containerStartSpec(&pod.Spec.Containers[idx]))
```

### createPodSandbox

pkg/kubelet/kuberuntime/kuberuntime_sandbox.go

```go
podSandboxConfig, err := m.generatePodSandboxConfig(pod, attempt)

// Create pod logs directory
err = m.osInterface.MkdirAll(podSandboxConfig.LogDirectory, 0755)

runtimeHandler, err = m.runtimeClassManager.LookupRuntimeHandler(pod.Spec.RuntimeClassName)

podSandBoxID, err := m.runtimeService.RunPodSandbox(podSandboxConfig, runtimeHandler)
```

pkg/kubelet/dockershim/docker_sandbox.go
RunPodSandbox
```go
// Step 1: Pull the image for the sandbox.
image := defaultSandboxImage

// Step 2: Create the sandbox container.
createConfig, err := ds.makeSandboxDockerConfig(config, image)
createResp, err := ds.client.CreateContainer(*createConfig)
ds.setNetworkReady(createResp.ID, false)

// Step 3: Create Sandbox Checkpoint.
ds.checkpointManager.CreateCheckpoint(createResp.ID, constructPodSandboxCheckpoint(config))

// Step 4: Start the sandbox container.
err = ds.client.StartContainer(createResp.ID)

// Step 5: Setup networking for the sandbox.
// All pod networking is setup by a CNI plugin discovered at startup time.
// This plugin assigns the pod ip, sets up routes inside the sandbox,
// creates interfaces etc. In theory, its jurisdiction ends with pod
// sandbox networking, but it might insert iptables rules or open ports
// on the host as well, to satisfy parts of the pod spec that aren't
// recognized by the CNI standard yet.
err = ds.network.SetUpPod(config.GetMetadata().Namespace, config.GetMetadata().Name, cID, config.Annotations, networkOptions)
```

ds.network.SetUpPod --> pluginManager.SetUpPod --> cniNetworkPlugin.SetUpPod

pkg/kubelet/dockershim/network/cni/cni.go
SetUpPod
```go
plugin.addToNetwork(cniTimeoutCtx, plugin.getDefaultNetwork(), name, namespace, id, netnsPath, annotations, options)
```

addToNetwork
```go
rt, err := plugin.buildCNIRuntimeConf(podName, podNamespace, podSandboxID, podNetnsPath, annotations, options) {
    rt := &libcni.RuntimeConf{
		ContainerID: podSandboxID.ID,
		NetNS:       podNetnsPath,
		IfName:      network.DefaultInterfaceName,
		CacheDir:    plugin.cacheDir,
		Args: [][2]string{
			{"IgnoreUnknown", "1"},
			{"K8S_POD_NAMESPACE", podNs},
			{"K8S_POD_NAME", podName},
			{"K8S_POD_INFRA_CONTAINER_ID", podSandboxID.ID},
		},
	}
}

res, err := cniNet.AddNetworkList(ctx, netConf, rt)
```

github.com/containernetworking/cni/libcni/api.go

cni interface
```go
type CNI interface {
	AddNetworkList(ctx context.Context, net *NetworkConfigList, rt *RuntimeConf) (types.Result, error)
	CheckNetworkList(ctx context.Context, net *NetworkConfigList, rt *RuntimeConf) error
	DelNetworkList(ctx context.Context, net *NetworkConfigList, rt *RuntimeConf) error
	GetNetworkListCachedResult(net *NetworkConfigList, rt *RuntimeConf) (types.Result, error)
	GetNetworkListCachedConfig(net *NetworkConfigList, rt *RuntimeConf) ([]byte, *RuntimeConf, error)

	AddNetwork(ctx context.Context, net *NetworkConfig, rt *RuntimeConf) (types.Result, error)
	CheckNetwork(ctx context.Context, net *NetworkConfig, rt *RuntimeConf) error
	DelNetwork(ctx context.Context, net *NetworkConfig, rt *RuntimeConf) error
	GetNetworkCachedResult(net *NetworkConfig, rt *RuntimeConf) (types.Result, error)
	GetNetworkCachedConfig(net *NetworkConfig, rt *RuntimeConf) ([]byte, *RuntimeConf, error)

	ValidateNetworkList(ctx context.Context, net *NetworkConfigList) ([]string, error)
	ValidateNetwork(ctx context.Context, net *NetworkConfig) ([]string, error)
}

AddNetworkList {
    c.addNetwork {
        pluginPath, err := c.exec.FindInPath(net.Network.Type, c.Path)

        invoke.ExecPluginWithResult(ctx, pluginPath, newConf.Bytes, c.args("ADD", rt), c.exec) {
            stdoutBytes, err := exec.ExecPlugin(ctx, pluginPath, netconf, args.AsEnv())
        }
    }
}

args.AsEnv {
    env = append(env,
		"CNI_COMMAND="+args.Command,
		"CNI_CONTAINERID="+args.ContainerID,
		"CNI_NETNS="+args.NetNS,
		"CNI_ARGS="+pluginArgsStr,
		"CNI_IFNAME="+args.IfName,
		"CNI_PATH="+args.Path,
	)
}
```

### startContainer

pkg/kubelet/kuberuntime/kuberuntime_container.go

startContainer
```go
// Step 1: pull the image.
imageRef, msg, err := m.imagePuller.EnsureImageExists(pod, container, pullSecrets, podSandboxConfig)

// Step 2: create the container.
containerConfig, cleanupAction, err := m.generateContainerConfig(container, pod, restartCount, podIP, imageRef, podIPs, target)

m.internalLifecycle.PreCreateContainer(pod, container, containerConfig)

containerID, err := m.runtimeService.CreateContainer(podSandboxID, containerConfig, podSandboxConfig)

m.internalLifecycle.PreStartContainer(pod, container, containerID)

// Step 3: start the container.
err = m.runtimeService.StartContainer(containerID)

// Step 4: execute the post start hook.
msg, handlerErr := m.runner.Run(kubeContainerID, pod, container, container.Lifecycle.PostStart)
```