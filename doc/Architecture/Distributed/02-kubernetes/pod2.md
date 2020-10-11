# Pod

## Pod Phase

[lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)

A Pod's ```status``` field is a PodStatus object, which has a phase field

The phase of a Pod is simple, high-level summary of where the Pod is in its lifecycle. The phase is not intended to be a comprehensive rollup of observation of Container or Pod state, nor is it intended to be a comprehensive state machine.

The number and meaning of Pod phase values are tightly guarded. Other than what is documented here, nothing should be assumed about Pods that have a given ```phase``` value.

Here are the possible values of ```phase```:

|Value|Description|
|-|-|
|Pending|The Pod has been accepted by kubernetes ,but one or more of container image has not been created.This includes time before being scheduled as well as time spent downloading images over the network, which could take a while|
|Running|The Pod has been bond to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting|
|Succeeded|All Containers in the Pod have terminated in success, and will not be restarted.|
|Failed|All containers in Pod have terminated, and at least one container has terminated in failure. That is, the container either exit with non-zero status or was terminated by system|
|Unknown|For some reason the state of the Pod could not be obtained, typically due to an error in communicating with the host of the Pod.|

## Init Containers

[init containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)

Specialized containers that run before app containers in a Pod. Init Containers can contain utilities or setup scripts not present in an app image.

You can speficy init containers in the Pod specification alongside the ```containers``` array(which describes app containers)

### Understanding init containers

A Pod can have multiple containers running apps with it, but it can also have one or more init containers, which are run before the app containers are started.

Init containers are exactly like regular containers, except:

- init containers always run to completion
- Each init container must complete successfully brefore the next one starts

If a Pod's init container fails, Kubernetes repeatedly restarts the Pod until the init container succeeds. However, if the Pod has a ```restartPolicy``` of Never, Kubernetes does not restart the Pod.

To specify an init container of a Pod, add the ```initContainers``` field into the Pod specification, as an array of objects of type Container, alongside the app containers array. The status of the init containers is returned in ```.status.initContainerStatuses``` field as an array of the container statuses

### Differences from regular containers

Init containers support all the fields and feature of app containers, including resource limits, volumes, and security settings. However, the resource requests and limits for an init container are handled differently.

Also, init containers do not support ```lifecycle```, ```livenessProbe```, ```readinessProbe```, or ```startupProbe``` because they must run to completion before the Pod can be ready.

If you specify multiple init containers for a Pod, Kubelet runs each init container sequentially. Each init container must succeed brefore the next can run. When all of the init containers have run to completion, Kubelet initializes the application containers for the Pod and runs them as usual.

### Using init containers

Because init containers have separate images from app containers, they have some advantages for start-up related code:

- Init containers can contain utilities or custom code for setup that are not present in an app image. For example, there is no need to make an image ```FROM``` another image just to use a tool like `sed`, `awk`, `python`, or `dig` during setup

- The application image builder and deployer roles can work independently without the need to jointly build a single app image

- Init containers can run with a different view of the filesystem than app containers in the same Pod. Consequently, they can be given access to Secrets that app containers cannot access

- Because init containers run to completion before any app containers start, init containers offer a mechanism to block or delay app container startup until a set of preconditions are met. Once preconditions are met, all of the app containers in a Pod can start in parallel.

- Init containers can securely run utilities or custom code that would otherwise make an app container image less secure. By keeping unnecessary tools separate you can limit the attack surface of your app container image



## Pod Preset

[Pod Preset](https://kubernetes.io/docs/concepts/workloads/pods/podpreset/)

### Understanding Pod presets

A PodPreset is an API resource for injecting additional runtime requirements into a Pod at creation time. You use label selector to specify the Pods to which a given Podpreset applies.

Using a PodPreset allows pod template authors to not have to explicitly provide all information for every pod. This way, authors of pod templates consuming a specific service do not need to know all the details about that service.