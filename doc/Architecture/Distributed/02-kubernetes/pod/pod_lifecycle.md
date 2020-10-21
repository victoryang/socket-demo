# Pod Lifecycle

[pod lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)

[Pod philosophy and lifecycle](https://www.jianshu.com/p/49c21b5feb99)

## Pod Phase

The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The phase is not intended to be a comprehensive rollup of observations of container or Pod state, nor is it intended to be a comprehensive state machine.

The number and meanings of Pod phase values are tightly guarded. Other than what is documented here, nothing should be assumed about Pods that have a given `phase` value.

Here are the possible values of `phase`:

|Value|Description|
|-|-|
|Pending|The Pod has been accepted by the Kubernetes cluster, but one or more of the containers has not been set up and made ready to run. This includes time a Pod spends waiting to be scheduled as well as the time spent downloading container images over the network|
|Running|The Pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting|
|Succeeded|All containers in the Pod have terminated in success, and will not be restarted|
|Failed|All containers in the Pod have terminated, and at least one container has terminated in failure. That is, the container either exited with non-zero status or was terminated by the system|
|Unknown|For some reason the state of the Pod could not be obtained. This phase typically occurs due to an error in communicating with the node where the Pod should be running|

If a node dies or is disconnected from the rest of the cluster, Kubernetes applies a policy for setting the `phase` of all Pods on the lost node to Failed.

## Container States

As well as the phase of the Pod overall, Kubernetes tracks the state of each container inside a Pod. You can use [container life cycle hooks](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/) to trigger events to run at certain points in a container's lifecycle.

Once the scheduler assigns a Pod to a Node,the kubelet starts creating containers for that Pod using a container runtime. There are three possible container state: `Waiting`, `Running`, and `Terminating`.

To check the state of a Pod's containers, you can use `kubectl describe pod <name-of-pod>`. The output shows the state for each container within the pod.

Each state has a specific meaning:

### Waiting

If a container is not in either `Runing` or `Terminated` state, it is `wating`. A container in `Waiting` state is still running operations it requires in order to complete start up: for example, pulling the container image from a container image registry, or applying Secret data.

When you use `kubectl` to query a Pod with a container that is `Waiting`, you also see a Reason field to summarize why the container is in that state.

### Running

The `Running` status indicates that a container is executing without issues. If there was a `PostStart` hook configured, it has already executed and finished.

When you use `kubectl` to query a Pod with a container that is Running, you also see information about when the container entered the `Running` state.

### Terminated

A Container in the `Terminated` state began execution and then ran to completion or failed for some reason.

When you use `kubectl` to query a Pod with a container that is `Terminated`, you see a reason, an exit code, and the start and finish time for that container's period of execution.

If a container has a `preStop` hook configured, that runs before the container enters the `Terminated` state.

## Container restart policy

The `spec` of a Pod has a `restartPolicy` field with possible value Always, OnFailure, and Never. The default value is Always.

The `restartPolicy` applies to all the containers in the Pod. `restartPolicy` only refers to restarts of the containers by kubelet on the same node. After containers in a Pod exit, the kubelet restarts them with an exponential back-off delay(10s,20s,40s,...), that is capped at five minutes. Once a container has executed for 10 minutes without any problems, the kubelet resets the restart backoff timer for that container.

## Pod conditions

A Pod has a PodStatus, which has an array of PodConditions through which Pod has or has not passed:

- `PodScheduled`: the Pod has been scheduled to a node.
- `ContainersReady`: all containers in the Pod are ready.
- `Initialized`: all init containers have started successfully.
- `Ready`: the Pod is able to serve requests and should be added to the load balancing pools of all matching Services.

## Container probes

A Probe is a diagnostic performed periodically by the kubelet on a container. To perform a diagnostic, the kubelet calls a Handler implemented by the container. There are three types of handlers:

- ExecAction: Executes a specified command inside the container. The diagnostic is considered successful if the command exits with a status code of 0.

- TCPSocketAction: Perform a TCP check against the Pod's IP address on a specified port. The diagnostic is considered successful if the port is open.

- HTTPGetAction: Perform an HTTP `GET` request against the Pod's IP address on a specified port and path. The diagnostic is considered succussful if the response has a status code greater than or equal to 200 and less than 400.

Each probe has one of three results:

- `Success`: The container passed the diagnostic.
- `Failure`: The container failed the diagnostic.
- `Unknown`: The diagnostic failed, so no action should be taken.

The kubelet can optionally perform and react to three kinds of probes on running containers:

- `livenessProbe`: Indicates whether the container is running.
- `readinessProbe`: Indicates whether the container is ready to respond to requests.
- `startupProbe`: Indicates whether the application within the container is started.

If probe is not set, default state is `success`

### 