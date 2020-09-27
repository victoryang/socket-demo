# Issue

## MacVlan API

### docker api error log

```
"Calling POST /containers/create?name=test_mac"
"form data: {\"Cmd\":[\"ash\"],\"Entrypoint\":null,\"HostConfig\":{\"ConsoleSize\":[0,0],\"LogConfig\":{},\"NetworkMode\":\"mac1\",\"Privileged\":true,\"RestartPolicy\":{}},\"Image\":\"alpine:latest\"}"


"container mounted via layerStore: &{/var/lib/docker/overlay2/57b481e553d535f4c2c355034f664271fa399da474674239c33138f6a1ce1749/merged 0x5647e5c24e80 0x5647e5c24e80}"


"Calling GET /version"


"Calling POST /containers/e61fac16e0520436103cce56e766868dd3ab503f66e09d51155dc676e6b1ff67/start"


"container mounted via layerStore: &{/var/lib/docker/overlay2/57b481e553d535f4c2c355034f664271fa399da474674239c33138f6a1ce1749/merged 0x5647e5c24e80 0x5647e5c24e80}"


"Assigning addresses for endpoint test_mac's interface on network mac1"
"RequestAddress(LocalDefault/192.168.16.0/20, <nil>, map[])"
"Request address PoolID:192.168.16.0/20 App: ipam/default/data, ID: LocalDefault/192.168.16.0/20, DBIndex: 0x0, Bits: 4096, Unselected: 4093, Sequence: (0xc0000000, 1)->(0x0, 126)->(0x1, 1)->end Curr:0 Serial:false PrefAddress:<nil> "
"Assigning addresses for endpoint test_mac's interface on network mac1"
"787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 (efb98b4).addSvcRecords(test_mac, 192.168.16.2, <nil>, true) updateSvcRecord sid:787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155"
"787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 (efb98b4).addSvcRecords(e61fac16e052, 192.168.16.2, <nil>, false) updateSvcRecord sid:787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155"
"Macvlan Endpoint Joined with IPv4_Addr: 192.168.16.2, Gateway: 192.168.16.1, MacVlan_Mode: bridge, Parent: enp0s3"
"787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 (efb98b4).addSvcRecords(test_mac, 192.168.16.2, <nil>, true) updateSvcRecord sid:787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155"
"787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 (efb98b4).addSvcRecords(e61fac16e052, 192.168.16.2, <nil>, false) updateSvcRecord sid:787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155"

"Programming external connectivity on endpoint test_mac (787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155)"
"EnableService e61fac16e0520436103cce56e766868dd3ab503f66e09d51155dc676e6b1ff67 START"
"EnableService e61fac16e0520436103cce56e766868dd3ab503f66e09d51155dc676e6b1ff67 DONE"

"createSpec: cgroupsPath: system.slice:docker:e61fac16e0520436103cce56e766868dd3ab503f66e09d51155dc676e6b1ff67"
"bundle dir created" bundle=/var/run/docker/containerd/e61fac16e0520436103cce56e766868dd3ab503f66e09d51155dc676e6b1ff67 module=libcontainerd namespace=moby root=/var/lib/docker/overlay2/57b481e553d535f4c2c355034f664271fa399da474674239c33138f6a1ce1749/merged
"sandbox set key processing took 348.928598ms for container e61fac16e0520436103cce56e766868dd3ab503f66e09d51155dc676e6b1ff67"
event module=libcontainerd namespace=moby topic=/tasks/create
event module=libcontainerd namespace=moby topic=/tasks/start



event module=libcontainerd namespace=moby topic=/tasks/exit
"Revoking external connectivity on endpoint test_mac (787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155)"
event module=libcontainerd namespace=moby topic=/tasks/delete
"ignoring event" module=libcontainerd namespace=moby topic=/tasks/delete type="*events.TaskDelete"
"787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 (efb98b4).deleteSvcRecords(test_mac, 192.168.16.2, <nil>, true) updateSvcRecord sid:787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 "
"787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 (efb98b4).deleteSvcRecords(e61fac16e052, 192.168.16.2, <nil>, false) updateSvcRecord sid:787322942eac07177e8bb008586f6c708b2a672f1d2ef23595a18d3be3e1f155 "
"Releasing addresses for endpoint test_mac's interface on network mac1"
"ReleaseAddress(LocalDefault/192.168.16.0/20, 192.168.16.2)"
"Released address PoolID:LocalDefault/192.168.16.0/20, Address:192.168.16.2 Sequence:App: ipam/default/data, ID: LocalDefault/192.168.16.0/20, DBIndex: 0x0, Bits: 4096, Unselected: 4092, Sequence: (0xe0000000, 1)->(0x0, 126)->(0x1, 1)->end Curr:3"
```

### docker run log
```
"Calling POST /v1.40/containers/create?name=test_mac3"
"form data: {\"AttachStderr\":false,\"AttachStdin\":false,\"AttachStdout\":false,\"Cmd\":[\"ash\"],\"Domainname\":\"\",\"Entrypoint\":null,\"Env\":[],\"HostConfig\":{\"AutoRemove\":false,\"Binds\":null,\"BlkioDeviceReadBps\":null,\"BlkioDeviceReadIOps\":null,\"BlkioDeviceWriteBps\":null,\"BlkioDeviceWriteIOps\":null,\"BlkioWeight\":0,\"BlkioWeightDevice\":[],\"CapAdd\":null,\"CapDrop\":null,\"Capabilities\":null,\"Cgroup\":\"\",\"CgroupParent\":\"\",\"ConsoleSize\":[0,0],\"ContainerIDFile\":\"\",\"CpuCount\":0,\"CpuPercent\":0,\"CpuPeriod\":0,\"CpuQuota\":0,\"CpuRealtimePeriod\":0,\"CpuRealtimeRuntime\":0,\"CpuShares\":0,\"CpusetCpus\":\"\",\"CpusetMems\":\"\",\"DeviceCgroupRules\":null,\"DeviceRequests\":null,\"Devices\":[],\"Dns\":[],\"DnsOptions\":[],\"DnsSearch\":[],\"ExtraHosts\":null,\"GroupAdd\":null,\"IOMaximumBandwidth\":0,\"IOMaximumIOps\":0,\"IpcMode\":\"\",\"Isolation\":\"\",\"KernelMemory\":0,\"KernelMemoryTCP\":0,\"Links\":null,\"LogConfig\":{\"Config\":{},\"Type\":\"\"},\"MaskedPaths\":null,\"Memory\":0,\"MemoryReservation\":0,\"MemorySwap\":0,\"MemorySwappiness\":-1,\"NanoCpus\":0,\"NetworkMode\":\"mac1\",\"OomKillDisable\":false,\"OomScoreAdj\":0,\"PidMode\":\"\",\"PidsLimit\":0,\"PortBindings\":{},\"Privileged\":false,\"PublishAllPorts\":false,\"ReadonlyPaths\":null,\"ReadonlyRootfs\":false,\"RestartPolicy\":{\"MaximumRetryCount\":0,\"Name\":\"no\"},\"SecurityOpt\":null,\"ShmSize\":0,\"UTSMode\":\"\",\"Ulimits\":null,\"UsernsMode\":\"\",\"VolumeDriver\":\"\",\"VolumesFrom\":null},\"Hostname\":\"\",\"Image\":\"alpine:latest\",\"Labels\":{},\"NetworkingConfig\":{\"EndpointsConfig\":{}},\"OnBuild\":null,\"OpenStdin\":true,\"StdinOnce\":false,\"Tty\":true,\"User\":\"\",\"Volumes\":{},\"WorkingDir\":\"\"}"


"container mounted via layerStore: &{/var/lib/docker/overlay2/979d04f6b86ee19276632bee27f509d1c89573e472c4df86d54060669b2654af/merged 0x5647e5c24e80 0x5647e5c24e80}"


"Calling POST /v1.40/containers/ec5ea22a356c587fa0df50d653c02de9020bd3530bbd1eb329b05f089ebb8b3d/wait?condition=next-exit"


"Calling POST /v1.40/containers/ec5ea22a356c587fa0df50d653c02de9020bd3530bbd1eb329b05f089ebb8b3d/start"

"container mounted via layerStore: &{/var/lib/docker/overlay2/979d04f6b86ee19276632bee27f509d1c89573e472c4df86d54060669b2654af/merged 0x5647e5c24e80 0x5647e5c24e80}"


"Assigning addresses for endpoint test_mac3's interface on network mac1"
"RequestAddress(LocalDefault/192.168.16.0/20, <nil>, map[])"
"Request address PoolID:192.168.16.0/20 App: ipam/default/data, ID: LocalDefault/192.168.16.0/20, DBIndex: 0x0, Bits: 4096, Unselected: 4093, Sequence: (0xc0000000, 1)->(0x0, 126)->(0x1, 1)->end Curr:3 Serial:false PrefAddress:<nil> "
"Assigning addresses for endpoint test_mac3's interface on network mac1"
"84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649 (efb98b4).addSvcRecords(test_mac3, 192.168.16.2, <nil>, true) updateSvcRecord sid:84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649"
"84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649 (efb98b4).addSvcRecords(ec5ea22a356c, 192.168.16.2, <nil>, false) updateSvcRecord sid:84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649"
"Macvlan Endpoint Joined with IPv4_Addr: 192.168.16.2, Gateway: 192.168.16.1, MacVlan_Mode: bridge, Parent: enp0s3"
"84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649 (efb98b4).addSvcRecords(test_mac3, 192.168.16.2, <nil>, true) updateSvcRecord sid:84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649"
"84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649 (efb98b4).addSvcRecords(ec5ea22a356c, 192.168.16.2, <nil>, false) updateSvcRecord sid:84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649"

"Programming external connectivity on endpoint test_mac3 (84064be3540d73ad03e15832e895e7a8291916028c3eb3dee114239cf527f649)"
"EnableService ec5ea22a356c587fa0df50d653c02de9020bd3530bbd1eb329b05f089ebb8b3d START"
"EnableService ec5ea22a356c587fa0df50d653c02de9020bd3530bbd1eb329b05f089ebb8b3d DONE"

"createSpec: cgroupsPath: system.slice:docker:ec5ea22a356c587fa0df50d653c02de9020bd3530bbd1eb329b05f089ebb8b3d"
"bundle dir created" bundle=/var/run/docker/containerd/ec5ea22a356c587fa0df50d653c02de9020bd3530bbd1eb329b05f089ebb8b3d module=libcontainerd namespace=moby root=/var/lib/docker/overlay2/979d04f6b86ee19276632bee27f509d1c89573e472c4df86d54060669b2654af/merged
"sandbox set key processing took 503.087651ms for container ec5ea22a356c587fa0df50d653c02de9020bd3530bbd1eb329b05f089ebb8b3d"
event module=libcontainerd namespace=moby topic=/tasks/create
event module=libcontainerd namespace=moby topic=/tasks/start
```