# Controller

cmd/kube-controller-manager/app/controllermanager.go
```

startServiceAccountTokenController

NewControllerManagerCommand {
    KnownControllers() {
        NewControllerInitializers
    }
}
```

// NewControllerInitializers is a public map of named controller groups (you can start more than one in an init func)
// paired to their InitFunc.  This allows for structured downstream composition and subdivision.

```go
controllers := map[string]InitFunc{}
    controllers["endpoint"] = startEndpointController
    controllers["endpointslice"] = startEndpointSliceController
    controllers["endpointslicemirroring"] = startEndpointSliceMirroringController
    controllers["replicationcontroller"] = startReplicationController
    controllers["podgc"] = startPodGCController
    controllers["resourcequota"] = startResourceQuotaController
    controllers["namespace"] = startNamespaceController
    controllers["serviceaccount"] = startServiceAccountController
    controllers["garbagecollector"] = startGarbageCollectorController
    controllers["daemonset"] = startDaemonSetController
    controllers["job"] = startJobController
    controllers["deployment"] = startDeploymentController
    controllers["replicaset"] = startReplicaSetController
    controllers["horizontalpodautoscaling"] = startHPAController
    controllers["disruption"] = startDisruptionController
    controllers["statefulset"] = startStatefulSetController
    controllers["cronjob"] = startCronJobController
    controllers["csrsigning"] = startCSRSigningController
    controllers["csrapproving"] = startCSRApprovingController
    controllers["csrcleaner"] = startCSRCleanerController
    controllers["ttl"] = startTTLController
    controllers["bootstrapsigner"] = startBootstrapSignerController
    controllers["tokencleaner"] = startTokenCleanerController
    controllers["nodeipam"] = startNodeIpamController
    controllers["nodelifecycle"] = startNodeLifecycleController
    if loopMode == IncludeCloudLoops {
        controllers["service"] = startServiceController
        controllers["route"] = startRouteController
        controllers["cloud-node-lifecycle"] = startCloudNodeLifecycleController
        // TODO: volume controller into the IncludeCloudLoops only set.
    }
    controllers["persistentvolume-binder"] = startPersistentVolumeBinderController
    controllers["attachdetach"] = startAttachDetachController
    controllers["persistentvolume-expander"] = startVolumeExpandController
    controllers["clusterrole-aggregation"] = startClusterRoleAggregrationController
    controllers["pvc-protection"] = startPVCProtectionController
    controllers["pv-protection"] = startPVProtectionController
    controllers["ttl-after-finished"] = startTTLAfterFinishedController
    controllers["root-ca-cert-publisher"] = startRootCACertPublisher
    controllers["ephemeral-volume"] = startEphemeralVolumeController
    if utilfeature.DefaultFeatureGate.Enabled(genericfeatures.APIServerIdentity) &&
        utilfeature.DefaultFeatureGate.Enabled(genericfeatures.StorageVersionAPI) {
        controllers["storage-version-gc"] = startStorageVersionGCController
    }
```

StartControllers
```
```

## startNodeLifecycleController

pkg/controller/nodelifecycle/node_lifecycle_controller.go

```
// NewNodeLifecycleController returns a new taint controller.
NewNodeLifecycleController (
    ctx.InformerFactory.Coordination().V1().Leases(),
    ctx.InformerFactory.Core().V1().Pods(),
    ctx.InformerFactory.Core().V1().Nodes(),
    ctx.InformerFactory.Apps().V1().DaemonSets(),
    // node lifecycle controller uses existing cluster role from node-controller
    ctx.ClientBuilder.ClientOrDie("node-controller"),
    ctx.ComponentConfig.KubeCloudShared.NodeMonitorPeriod.Duration,
    ctx.ComponentConfig.NodeLifecycleController.NodeStartupGracePeriod.Duration,
    ctx.ComponentConfig.NodeLifecycleController.NodeMonitorGracePeriod.Duration,
    ctx.ComponentConfig.NodeLifecycleController.PodEvictionTimeout.Duration,
    ctx.ComponentConfig.NodeLifecycleController.NodeEvictionRate,
    ctx.ComponentConfig.NodeLifecycleController.SecondaryNodeEvictionRate,
    ctx.ComponentConfig.NodeLifecycleController.LargeClusterSizeThreshold,
    ctx.ComponentConfig.NodeLifecycleController.UnhealthyZoneThreshold,
    ctx.ComponentConfig.NodeLifecycleController.EnableTaintManager,
) {
    podInformer.Informer().AddEventHandler(
        cache.ResourceEventHandlerFuncs{
            AddFunc,
            UpdateFunc,
            DeleteFunc,
        }
    )

nc.getPodsAssignedToNode

nc.podLister = podInformer.Lister()

nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
    AddFunc,
    UpdateFunc,
    DeleteFunc,
})

nc.leaseLister = leaseInformer.Lister()

nc.nodeLister = nodeInformer.Lister()

nc.daemonSetStore = daemonSetInformer.Lister()
}
```

nc.podUpdated()
```
// processPod is processing events of assigning pods to nodes. In particular:
// 1. for NodeReady=true node, taint eviction for this pod will be cancelled
// 2. for NodeReady=false or unknown node, taint eviction of pod will happen and pod will be marked as not ready
// 3. if node doesn't exist in cache, it will be skipped and handled later by doEvictionPass
```

nc.taintManager.PodUpdated()
```
// PodUpdated is used to notify NoExecuteTaintManager about Pod changes
```

lifecycleController.Run()
```
// Run starts an asynchronous loop that monitors the status of cluster nodes.

go nc.taintManager.Run(stopCh)

go wait.Until(nc.doNodeProcessingPassWorker, time.Second, stopCh)

go wait.Until(nc.doPodProcessingWorker, time.Second, stopCh)

go wait.Until(nc.doNoExecuteTaintingPass, scheduler.NodeEvictionPeriod, stopCh)

nc.monitorNodeHealth()
```