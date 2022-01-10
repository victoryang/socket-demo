# Code

## kubernetes/cmd/kubelet

kubelet.go
```
// create cobra commandline args
app.NewKubeletCommand
```

app/server.go
- validate commandline args
- generate kubelet server by kubeletFlags and kubeletConfig

```
// construct a KubeletServer from kubeletFlags and kubeletConfig
kubeletServer := &options.KubeletServer{
    KubeletFlags:         *kubeletFlags,
    KubeletConfiguration: *kubeletConfig,
}
```

kubeletFlags vs kubeletConfig
```
// KubeletFlags contains configuration flags for the Kubelet.
// A configuration field should go in KubeletFlags instead of KubeletConfiguration if any of these are true:
// - its value will never, or cannot safely be changed during the lifetime of a node, or
// - its value cannot be safely shared between nodes at the same time (e.g. a hostname);
//   KubeletConfiguration is intended to be shared between nodes.
// In general, please try to avoid adding flags or configuration fields,
// we already have a confusingly large amount of them.
```

func Run
```
initOS
run
```

func run
create kubeDeps, 
```
kubeDeps.ContainerManager, err = cm.NewContainerManager(
			kubeDeps.Mounter,
			kubeDeps.CAdvisorInterface,
			cm.NodeConfig{
				RuntimeCgroupsName:    s.RuntimeCgroups,
				SystemCgroupsName:     s.SystemCgroups,
				KubeletCgroupsName:    s.KubeletCgroups,
				ContainerRuntime:      s.ContainerRuntime,
				CgroupsPerQOS:         s.CgroupsPerQOS,
				CgroupRoot:            s.CgroupRoot,
				CgroupDriver:          s.CgroupDriver,
				KubeletRootDir:        s.RootDirectory,
				ProtectKernelDefaults: s.ProtectKernelDefaults,
				NodeAllocatableConfig: cm.NodeAllocatableConfig{
					KubeReservedCgroupName:   s.KubeReservedCgroup,
					SystemReservedCgroupName: s.SystemReservedCgroup,
					EnforceNodeAllocatable:   sets.NewString(s.EnforceNodeAllocatable...),
					KubeReserved:             kubeReserved,
					SystemReserved:           systemReserved,
					ReservedSystemCPUs:       reservedSystemCPUs,
					HardEvictionThresholds:   hardEvictionThresholds,
				},
				QOSReserved:                             *experimentalQOSReserved,
				ExperimentalCPUManagerPolicy:            s.CPUManagerPolicy,
				ExperimentalCPUManagerPolicyOptions:     cpuManagerPolicyOptions,
				ExperimentalCPUManagerReconcilePeriod:   s.CPUManagerReconcilePeriod.Duration,
				ExperimentalMemoryManagerPolicy:         s.MemoryManagerPolicy,
				ExperimentalMemoryManagerReservedMemory: s.ReservedMemory,
				ExperimentalPodPidsLimit:                s.PodPidsLimit,
				EnforceCPULimits:                        s.CPUCFSQuota,
				CPUCFSQuotaPeriod:                       s.CPUCFSQuotaPeriod.Duration,
				ExperimentalTopologyManagerPolicy:       s.TopologyManagerPolicy,
				ExperimentalTopologyManagerScope:        s.TopologyManagerScope,
			},
			s.FailSwapOn,
			devicePluginEnabled,
			kubeDeps.Recorder)
```

init runtime service before RunKubelet
```
 kubelet.PreInitRuntimeService
```

RunKubelet is responsible for setting up and running a kubelet
```
RunKubelet(s, kubeDeps, s.RunOnce)
```

RunKubelet
```
CreateAndInitKubelet(){
    // NewMainKubelet instantiates a new Kubelet object along with all the required internal modules. 
    // No initialization of Kubelet and its modules should happen here
    k = NewMainKubelet()

    // sends an event that the kubelet has started up
    k.BirthCry()
    // starts garbage collection threads
    // container and image gc
    k.StartGarbageCollection()
}

// podconfig
kubeDeps.PodConfig, err = makePodSourceConfig(kubeCfg, kubeDeps, nodeName, nodeHasSynced)
```

startKubelet
```
podCfg := kubeDeps.PodConfig
// Updates returns a channel of updates to the configuration, properly denormalized.
k.Run(podCfg.Updates())
```

## pkg/kubelet/kubelet.go

Bootstrap defines what kubelet can do, which can be invoked by cmd/kubelet
```
type Bootstrap interface {
    GetConfiguration() kubeletconfiginternal.KubeletConfiguration
    BirthCry()
    StartGarbageCollection()
    ListenAndServe(address net.IP, port uint, tlsOptions *server.TLSOptions, auth server.AuthInterface, enableDebuggingHandlers, enableContentionProfiling bool)
    ListenAndServeReadOnly(address net.IP, port uint)
    ListenAndServePodResources()
    Run(<-chan kubetypes.PodUpdate)
    RunOnce(<-chan kubetypes.PodUpdate) ([]RunPodResult, error)
}
```
### kubelet server

- containerRefManager
- oomWatcher
- klet.secretManager
- klet.configMapManager
- klet.livenewssManager
- klet.podManager
- klet.resourceAnalyzer
- imageManager
- klet.statusManager
- klet.volumeManager
- eviction.NewManager

```
nodeIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
nodeLister = corelisters.NewNodeLister(nodeIndexer)

// register pod source in three ways: 
// file config, url config and apiserver
kubeDeps.PodConfig, err = makePodSourceConfig(kubeCfg, kubeDeps, nodeName, nodeHasSynced)

containerGCPolicy := kubecontainer.GCPolicy

daemonEndpoints := &v1.NodeDaemonEndpoints

imageGCPolicy := images.ImageGCPolicy

evictionConfig := eviction.Config

kubeInformers := informers.NewSharedInformerFactory(kubeDeps.KubeClient, 0)
serviceLister = kubeInformers.Core().V1().Services().Lister()
serviceHasSynced = kubeInformers.Core().V1().Services().Informer().HasSynced
kubeInformers.Start(wait.NeverStop)

nodeRef := &v1.ObjectReference

oomWatcher, err := oomwatcher.NewWatcher(kubeDeps.Recorder)

secretManager = secret.NewWatchingSecretManager(kubeDeps.KubeClient, klet.resyncInterval)

configMapManager = configmap.NewWatchingConfigMapManager(kubeDeps.KubeClient, klet.resyncInterval)

imageBackOff := flowcontrol.NewBackOff(backOffPeriod, MaxContainerBackOff)

klet.livenessManager = proberesults.NewManager()
klet.readinessManager = proberesults.NewManager()
klet.startupManager = proberesults.NewManager()

// cache stores the PodStatus for the pods.
klet.podCache = kubecontainer.NewCache()

mirrorPodClient := kubepod.NewBasicMirrorClient(klet.kubeClient, string(nodeName), nodeLister)
klet.podManager = kubepod.NewBasicPodManager(mirrorPodClient, secretManager, configMapManager)

klet.statusManager = status.NewManager(klet.kubeClient, klet.podManager, klet)

klet.resourceAnalyzer = serverstats.NewResourceAnalyzer(klet, kubeCfg.VolumeStatsAggPeriod.Duration, kubeDeps.Recorder)

klet.runtimeClassManager = runtimeclass.NewManager(kubeDeps.KubeClient)

klet.reasonCache = NewReasonCache()

klet.workQueue = queue.NewBasicWorkQueue(klet.clock)
klet.podWorkers = newPodWorkers

runtime, err := kuberuntime.NewKubeGenericRuntimeManager
klet.containerRuntime = runtime
klet.streamingRuntime = runtime
klet.runner = runtime
runtimeCache, err := kubecontainer.NewRuntimeCache
klet.runtimeCache = runtimeCache

klet.StatsProvider = stats.NewCRIStatsProvider

klet.pleg = pleg.NewGenericPLEG(klet.containerRuntime, plegChannelCapacity, plegRelistPeriod, klet.podCache, clock.RealClock{})

klet.runtimeState = newRuntimeState(maxWaitForContainerRuntime)
klet.runtimeState.addHealthCheck("PLEG", klet.pleg.Healthy)
klet.updatePodCIDR

containerGC, err := kubecontainer.NewContainerGC

klet.containerDeletor = newPodContainerDeletor

imageManager, err := images.NewImageGCManager

klet.probeManager = prober.NewManager

tokenManager := token.NewManager

klet.volumePluginMgr, err = NewInitializedVolumePluginMgr

klet.pluginManager = pluginmanager.NewPluginManager

klet.volumeManager = volumemanager.NewVolumeManager

evictionManager, evictionAdmitHandler := eviction.NewManager

klet.admitHandlers.AddPodAdmitHandler(evictionAdmitHandler)

klet.admitHandlers.AddPodAdmitHandler(sysctlsAllowlist)

klet.admitHandlers.AddPodAdmitHandler(klet.containerManager.GetAllocateResourcesPodAdmitHandler())

klet.admitHandlers.AddPodAdmitHandler(lifecycle.NewPredicateAdmitHandler(klet.getNodeAnyWay, criticalPodAdmissionHandler, klet.containerManager.UpdatePluginResources))

klet.softAdmitHandlers.AddPodAdmitHandler(lifecycle.NewAppArmorAdmitHandler(klet.appArmorValidator))

klet.softAdmitHandlers.AddPodAdmitHandler(lifecycle.NewNoNewPrivsAdmitHandler(klet.containerRuntime))
klet.softAdmitHandlers.AddPodAdmitHandler(lifecycle.NewProcMountAdmitHandler(klet.containerRuntime))

klet.nodeLeaseController = lease.NewController

shutdownManager, shutdownAdmitHandler := nodeshutdown.NewManager
klet.admitHandlers.AddPodAdmitHandler(shutdownAdmitHandler)
```

#### member structure

##### podCache

pkg/kubelet/container/cache.go

Cache stores the PodStatus for the pods. It represents *all* the visible pods/containers in the container runtime. All cache entries are at least as new or newer than the global timestamp, while individual entries may be slightly newer than the global timestamp. If a pod has no states known by the runtime, Cache returns an empty PodStatus object with ID populated.

Cache provides two methods to retrieve the PodStatus: the non-blocking Get() and the blocking GetNewerThan() method. The component responsible for populating the cache is expected to call Delete() to explicitly free the cache entries.

##### podmanager

pkg/kubelet/pod/pod_manager.go

Manager stores and manages access to pods, maintaining the mappings between static pods and mirror pods.

The kubelet discovers pod update from 3 sources: file, http, and apiserver. Pods from non-apiserver resources are called static pods, and API server is not aware of the existence of static pods. In order to monitor the status of such pods, the kubelet creates a mirror pod for each static pod via the API server.

A mirror pod has the same pod full name (name and namespace) as its static couterpart (albeit different metadata such as UID,etc). By leveraging the fact that the kubelet reports the pod status using the pod full name, the status of the mirror pod always reflects the actual status of the static pod. When a static pod gets deleted, the associated orphaned mirror pod will also be removed.

### kubelet.Run
kubelet.Run
```
 kl.initializeModules

 volumeManager.Run

 kl.nodeLeaseController.Run

 kl.statusManager.Start()

 kl.runtimeClassManager.Start

// Start the pod lifecycle event generator.
kl.pleg.Start()
kl.syncLoop(updates, kl)
```

#### kubelet.syncLoop

```
// syncLoop is the main loop for processing changes. It watches for changes from
// three channels (file, apiserver, and http) and creates a union of them. For
// any new change seen, will run a sync against desired state and running state.

syncTicker := time.NewTicker(time.Second)

kl.syncLoopIteration(updates, handler, syncTicker.C, housekeepingTicker.C, plegCh)
```

#### kubelet.syncLoopIteration
```
// syncLoopIteration reads from various channels and dispatches pods to the
// given handler.
//
// Arguments:
// 1.  configCh:       a channel to read config events from
// 2.  handler:        the SyncHandler to dispatch pods to
// 3.  syncCh:         a channel to read periodic sync events from
// 4.  housekeepingCh: a channel to read housekeeping events from
// 5.  plegCh:         a channel to read PLEG updates from
```

syncHandler
```go
type syncHandler interface{
    HandlePodAdditions
    HandlePodUpdates
    HandlePodRemove
    HandlePodReconcile
    HandlePodSync
    HandlePodCleanups
}
```

- podmanager
- probemanager
- podworker
- podKiller

#### kl.syncPod

syncPod is the transaction script for the sync of a single pod (setting up) a pod. The reverse (teardown) is handled in syncTerminatingPod and syncTerminatedPod. If syncPod exits without error, then the pod runtime state is in sync with the desired configuration state (pod is running). If syncPod exits with a transient error, the next invocation of syncPod is expected to make progress towards reaching the runtime state.

Arguments:
o - the SyncPodOptions for this invocation

The workflow is:

* If the pod is being creatd, record pod worker start latency 
* Call generateAPIPodStatus to prepare an v1.PodStatus for the pod
* If the pod is being seen as running for the first time, record pod start latency
* Update the status of the pod in the status manager
* Kill the pod if it should not be running due to soft admission
* Create a mirror pod if the pod is a static pod, and does not already have a mirror pod
* Create the data directories for the pod if they do not exist
* Wait for volumes to attach/mount
* Fetch the pull secrets for the pod
* Call the container runtime's SyncPod callback
* Update the traffic shaping for the pod's ingress and egress limits

If any step of this workflow errors, the error is returned, and is repeated on the next syncPod call.

```
firstSeenTime = kubetypes.ConvertToTimestamp(firstSeenTimeStr).Get()

apiPodStatus := kl.generateAPIPodStatus(pod, podStatus)

existingStatus, ok := kl.statusManager.GetPodStatus(pod.UID)

kl.statusManager.SetPodStatus(pod, apiPodStatus)

kl.killPod(pod, p, nil)

kl.podManager.CreateMirrorPod(pod)

kl.makePodDataDirs(pod)

kl.volumeManager.WaitForAttachAndMount(pod)

kl.getPullSecretsForPod(pod)

kl.containerRuntime.SyncPod
```

#### kubeGenericRuntimeManager.SyncPod

SyncPod syncs the running pod into the desired pod by executing folloing steps:
1. Compute sandbox and container changes.
2. Kill pod sandbox if necessary.
3. Kill any containers that should not be running.
4. Create sandbox if necessary.
5. Create ephemeral containers.
6. Create init containers.
7. Create normal containers.

```
podContainerChanges := m.computePodActions(pod, podStatus)

killResult := m.killPodWithSyncResult

m.pruneInitContainersBeforeStart(pod, podStatus)

podSandboxID, msg, err = m.createPodSandbox

start := func(typeName, metricLabel string, spec *startSpec) {
    m.startContainer(podSandboxID, podSandboxConfig, spec, pod, podStatus, pullSecrets, podIP, podIPs)
}

start("ephemeral container", metrics.EphemeralContainer, ephemeralContainerStartSpec(&pod.Spec.EphemeralContainers[idx]))

start("init container", metrics.InitContainer, containerStartSpec(container))

start("container", metrics.Container, containerStartSpec(&pod.Spec.Containers[idx]))
```

kubeGenericRuntimeManager.startContainer

startContainer starts a container and returns a message indicates why it is failed on error.
It starts the container through the following steps:
* pull the image
* create the container
* start the container
* run the post start lifecycle hooks (if applicable)


#### probe
kl.probeManager.AddPod(pod)
start goroutines for probes
- startup
- readiness
- liveness

```
func (m *manager) AddPod(pod *v1.Pod) {
	m.workerLock.Lock()
	defer m.workerLock.Unlock()

	key := probeKey{podUID: pod.UID}
	for _, c := range pod.Spec.Containers {
		key.containerName = c.Name

		if c.StartupProbe != nil {
			key.probeType = startup
			if _, ok := m.workers[key]; ok {
				klog.ErrorS(nil, "Startup probe already exists for container",
					"pod", klog.KObj(pod), "containerName", c.Name)
				return
			}
			w := newWorker(m, startup, pod, c)
			m.workers[key] = w
			go w.run()
		}

		if c.ReadinessProbe != nil {
			key.probeType = readiness
			if _, ok := m.workers[key]; ok {
				klog.ErrorS(nil, "Readiness probe already exists for container",
					"pod", klog.KObj(pod), "containerName", c.Name)
				return
			}
			w := newWorker(m, readiness, pod, c)
			m.workers[key] = w
			go w.run()
		}

		if c.LivenessProbe != nil {
			key.probeType = liveness
			if _, ok := m.workers[key]; ok {
				klog.ErrorS(nil, "Liveness probe already exists for container",
					"pod", klog.KObj(pod), "containerName", c.Name)
				return
			}
			w := newWorker(m, liveness, pod, c)
			m.workers[key] = w
			go w.run()
		}
	}
}
```

#### deletePod
```
HandlePodRemoves

kl.deletePod(pod)

kl.podWorkers.UpdatePod(UpdatePodOptions{
    Pod:        pod,
    UpdateType: kubetypes.SyncPodKill,
})

m.killPodWithSyncResult

m.killContainersWithSyncResult

m.killContainer(pod, container.ID, container.Name, "", reasonUnknown, gracePeriodOverride)
```

### PLEG

```
type GenericPLEG struct {
	// The period for relisting.
	relistPeriod time.Duration
	// The container runtime.
	runtime kubecontainer.Runtime
	// The channel from which the subscriber listens events.
	eventChannel chan *PodLifecycleEvent
	// The internal cache for pod/container information.
	podRecords podRecords
	// Time of the last relisting.
	relistTime atomic.Value
	// Cache for storing the runtime states required for syncing pods.
	cache kubecontainer.Cache
	// For testability.
	clock clock.Clock
	// Pods that failed to have their status retrieved during a relist. These pods will be
	// retried during the next relisting.
	podsToReinspect map[types.UID]*kubecontainer.Pod
}
```

kl.pleg.Start 
```
// relist every 1 seconds
go wait.Until(g.relist, g.relistPeriod, wait.NeverStop)
```

k1.pleg.Healthy
```
elapsed > relistThreshold
```

relist queries the container runtime for list of pods/containers, compare with the internal pods/containers, and generates events accordingly.

GenericPLEG.relist
```
podList, err := g.runtime.GetPods(true)
pods := kubecontainer.Pods(podList)
g.podRecords.setCurrent(pods)

// Compare the old and the current pods, and generate events.
eventsByPodID := map[types.UID][]*PodLifecycleEvent{}
for pid := range g.podRecords {
    oldPod := g.podRecords.getOld(pid)
    pod := g.podRecords.getCurrent(pid)
    // Get all containers in the old and the new pod.
    allContainers := getContainersFromPods(oldPod, pod)
    for _, container := range allContainers {
        events := computeEvents(oldPod, pod, &container.ID)
        for _, e := range events {
            updateEvents(eventsByPodID, e)
        }
    }
}

computeEvents
    - generateEvents

// updateCache will inspect the pod and update the cache.
g.updateCache(pod, pid)
```

handler.HandlePodSyncs([]*v1.Pod{pod})

### Runtime

cmd/kubelet/app/server.go
kubelet.run
```
kubelet.PreInitRuntimeService

kubeDeps.ContainerManager, err = cm.NewContainerManager
```

pkg/kubelet/kubelet.go
kubelet.PreInitRuntimeService
start a runtime server which implys cri interface

- dockershim
- remote
    - cri server

```
// DockerContainerRuntime
runDockershim(
    kubeCfg,
    kubeDeps,
    crOptions,
    runtimeCgroups,
    remoteRuntimeEndpoint,
    remoteImageEndpoint,
    nonMasqueradeCIDR,
)

// RemoteRuntimeService
kubeDeps.RemoteRuntimeService, err = remote.NewRemoteRuntimeService(remoteRuntimeEndpoint, kubeCfg.RuntimeRequestTimeout.Duration)
```

- DockerContainerRuntime
    - pkg/kubelet/kubelet_dockershim.go
- RemoteRuntimeService
    - pkg/kubelet/cri/remote/remote.go

```
createAndInitKubelet(
    ...
    &kubeServer.ContainerRuntimeOptions,
    kubeServer.ContainerRuntime,
    ...
)
```

```
runtime, err := kuberuntime.NewKubeGenericRuntimeManager

klet.containerRuntime = runtime
klet.streamingRuntime = runtime
klet.runner = runtime

runtimeCache, err := kubecontainer.NewRuntimeCache(klet.containerRuntime)
klet.runtimeCache = runtimeCache

klet.dockerLegacyService = kubeDeps.dockerLegacyService
klet.runtimeService = kubeDeps.RemoteRuntimeService

klet.runtimeClassManager = runtimeclass.NewManager(kubeDeps.KubeClient)
```

pkg/kubelet/kuberuntime/kuberuntime_manager.go
NewKubeGenericRuntimeManager
```
kubeRuntimeManager := &kubeGenericRuntimeManager {
    ...
    runtimeService:         newInstrumentedRuntimeService(runtimeService),
    imageService:           newInstrumentedImageManagerService(imageService),
    ...
}
```

#### CRI

```
// CRIService includes all methods necessary for a CRI server.
type CRIService interface {
	runtimeapi.RuntimeServiceServer
	runtimeapi.ImageServiceServer
	Start() error
}
```

- RuntimeService
    - PodSandbox
    - Containers
    - Streaming API
        - Exec
        - Attach
        - PortForward
    - Verion + Status
- ImageService

kubeRuntimeManager := &kubeGenericRuntimeManager {
    ...
    runtimeService: newInstrumentedRuntimeSerice(runtimeService),
    ...
}

### probeManager

#### statusManager

```
// pkg/kubelet/kubelet.go
klet.statusManager = status.NewManager(klet.kubeClient, klet.podManager, klet)

// pkg/kubelet/status/status_manager.go
// Manager is the Source of truth for kubelet pod status, and should be kept up-to-date with
// the latest v1.PodStatus. It also syncs updates back to the API server.
type Manager interface {
	PodStatusProvider

	// Start the API server status sync loop.
	Start()

	// SetPodStatus caches updates the cached status for the given pod, and triggers a status update.
	SetPodStatus(pod *v1.Pod, status v1.PodStatus)

	// SetContainerReadiness updates the cached container status with the given readiness, and
	// triggers a status update.
	SetContainerReadiness(podUID types.UID, containerID kubecontainer.ContainerID, ready bool)

	// SetContainerStartup updates the cached container status with the given startup, and
	// triggers a status update.
	SetContainerStartup(podUID types.UID, containerID kubecontainer.ContainerID, started bool)

	// TerminatePod resets the container status for the provided pod to terminated and triggers
	// a status update.
	TerminatePod(pod *v1.Pod)

	// RemoveOrphanedStatuses scans the status cache and removes any entries for pods not included in
	// the provided podUIDs.
	RemoveOrphanedStatuses(podUIDs map[types.UID]bool)
}
```

// kl.Run()
kl.statusManager.Start()
```
go wait.Forever(func() {
    for {
        select {
        case syncRequest := <-m.podStatusChannel:
            klog.V(5).InfoS("Status Manager: syncing pod with status from podStatusChannel",
                "podUID", syncRequest.podUID,
                "statusVersion", syncRequest.status.version,
                "status", syncRequest.status.status)
            m.syncPod(syncRequest.podUID, syncRequest.status)
        case <-syncTicker:
            klog.V(5).InfoS("Status Manager: syncing batch")
            // remove any entries in the status channel since the batch will handle them
            for i := len(m.podStatusChannel); i > 0; i-- {
                <-m.podStatusChannel
            }
            m.syncBatch()
        }
    }
}, 0)

// m.syncPod
// syncPod syncs the given status with the API server. The caller must not hold the lock.
```

#### probeManager

```
// pkg/kubelet/kubelet.go
klet.livenessManager = proberesults.NewManager()
klet.readinessManager = proberesults.NewManager()
klet.startupManager = proberesults.NewManager()

klet.probeManager = prober.NewManager(
    klet.statusManager,
    klet.livenessManager,
    klet.readinessManager,
    klet.startupManager,
    klet.runner,
    kubeDeps.Recorder)
```

```
// pkg/kubelet/prober/prober_manager.go

```