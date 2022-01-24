# Dockershim

## SyncPod

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

## RunPodSandbox

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