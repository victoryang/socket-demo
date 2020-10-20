# Pod

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