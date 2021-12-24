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

## kubelet

pkg/kubelet/kubelet.go

Bootstrap
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

kubelet managers
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
klet.podCache = kubecontainer.NewCache()

mirrorPodClient := kubepod.NewBasicMirrorClient(klet.kubeClient, string(nodeName), nodeLister)
klet.podManager = kubepod.NewBasicPodManager(mirrorPodClient, secretManager, configMapManager)

klet.statusManager = status.NewManager(klet.kubeClient, klet.podManager, klet)

klet.resourceAnalyzer = serverstats.NewResourceAnalyzer(klet, kubeCfg.VolumeStatsAggPeriod.Duration, kubeDeps.Recorder)



```