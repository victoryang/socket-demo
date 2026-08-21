# Gunyah

## 构建

### hypervisor

**base**

- config_base (config)
- module_base (hyp)
- arch_base (arch)
- interface_base (interfaces)

**set and list**

- interfaces
- objects
- external_objects
- guestapis
- types
- hypercalls
- registers

**event**

- event_sources
- module_with_events
- interfaces_with_events

**template_engines**

- first_class_templates (first_class_object)
- typed_templates (typed)
- typed_guestapi_templates (typed_guestapi)
- hypercalls_templates (hypercalls)
- registers_templates (registers)

**process_variant_conf**

- 解析配置文件字段

**configs sanity checking**

**Architecture setup**

- Add the arch-specific include directories for asm/ headers
- Add the arch generic include directory for asm-generic/ headers
- Set up for a freestanding ARMv8.2 EL2 target
- enable stack protection by default

**Toolchain setup**

- LLVM
- clang
- cpp-dsl
- CTU rule
    - cc-ctu-ast
    - cc-ctu-map
    - cc-ctu-all
- static analyzer
    - cc-analyze

https://clang.llvm.org/docs/analyzer/user-docs/CrossTranslationUnit.html

**Parse the module configuations**

- parse_module_conf
- parse_interface_conf

**Global generated headers depends**

include/

config.h

```
grep -r "configs " hyp | grep build.conf

grep -r "interface " hyp | grep build.conf

grep -r "base_module " hyp | grep build.conf
```

**Hypercalls generated files**

guestapi/include/guest_interface.h  guestapi/src/guest_interface.c

**Set up the simple code generator**

autogen/

```
grep -r "template simple" hyp | grep build.conf
```

**Set up the Clang static analyser**

**Generate types and events for first class objects**

objects/

```
grep -r "first_class_object" hyp | grep build.conf
```

**Setup the types generator**

```
grep -r "template" hyp | grep  typed_guestapi
```

**Setup the hypercalls generator**

guestapi/

```
grep -r "hypercalls " hyp | grep build.conf
```

**Setup the events generators**

have_events = True

```
grep -r "events" hyp | grep build.conf
```

**Generate register accessors**

hypregisters.h

```
grep -r "registers" hyp | grep build.conf
```

**Build version setup**

**Symbols version setup**


## 启动流程

### platform
head.S
```asm
const soc_qemu_hyp_arg0
# __entry_el2
adrl x1, soc_qemu_hyp_arg0
str x0, [x1]
...
b	aarch64_init
```

init_el2.S
```asm
# aarch64_init
// Disable debug, aborts and interrupts.
    msr	DAIFSet, 0xf
...
    abs64	x9, SCTLR_EL2_BOOT
    msr	SCTLR_EL2, x9
    isb
...
// Save logical CPU number in w20 (call-preserved)
    mov	w20, w4
...
// Initialize the KASLR and return base virtual address
    bl	aarch64_init_kaslr
...
// Initialize and enable the MMU
    mov	x2, 1
    bl	aarch64_init_address_space
...
    mov	w0, w20
    bl	trigger_boot_runtime_first_init_event
...
    mov	w0, w20
    b	boot_cold_init
```

## hyp

### core

***api***
hypercall 原型

***boot***
启动流程代码

***cpulocal***
提供当前cpu index信息

***cspace_trivial***

***cspace_twolevel***
cspace源码

***globals***

***idle***
为每个cpu创建idle线程

***ipi***
ipi

***irq***
hypervisor 硬件中断处理

***object_standard***
内核对象管理

***partition_standard***
内存逻辑分区，hypervisor、vm 分别拥有自己的内存分区

***power***
系统和cpu电源管理，支持 suspend,resume,cpu_off,cpu_on，idle等功能

***preemt***

***schedule_fprr***
fixed priority rr，线程调度器，可设置线程的 priority，affinity，time_slice，pin，running_state(running/block)等

***task_queue***

***thread_standard***
线程创建、激活、失效、switch、kill、退出等的线程状态管理

***timer***
时钟相关

***vdevice***
虚拟设备处理，vdev内存范围管理和访问控制

***vectors***
中断向量表

***wait_queue_broadcast***
等待队列处理

### mem

***addrspace***
内核地址空间，一个addRspace可绑定多个thread


***allocator_boot***
boot阶段的内存分配器，启动阶段映射data heap的头 2MB 作为内核boot阶段的内核对象分配器

***allocator_list***
partition 分配器

***hyp_aspace***
hypervisor 地址空间

***memdb***
内存管理数据库，用于快速查询内存信息

***memextent***
内存范围对象，对一段内存范围的构建，查询，所有权管理，映射/去映射等操作的抽象结合

***memextent_sparse***


***page_table***
页表管理接口


### ipc

***door_bell***
基于virq注入的通知机制

***msgqueue***
基于 Hypervisor 地址空间内msgqueue的通信机制


### vm

***arm_pv_time***

***arm_vm_pmu***

***arm_vm_smmuv2***

***arm_vm_sve_simple***

***arm_vm_timer***

***psci***
psci common handle
- cpu handle
    - cpu on
    - cpu off
    - cpu suspend
    - cpu resume
    - affinity info
- system handle
    - system off
    - system reset
    - psci feature
- vpm handle
    - vpm attach
    - vpm bind/unbind virq
    - system_wakeup
    - threshold
- vcpu handle
    - resume
    - start
    - wakeup
    - run_check
    - power_on
    - power_off
    - stop
    - online
    - offline

***psci_osi***
psci osi mode

***psci_pc***
psci platfom-coordinated mode

***rootvm***
rootvm 初始化相关

rootvm_init
```C
// 创建 boot vcpu 参数
    thread_create_t params = {
		.scheduler_affinity	  = boot_cpu_idx,
		.scheduler_affinity_valid = true,
		.scheduler_priority	  = ROOTVM_PRIORITY,
		.scheduler_priority_valid = true,
	};

// Allocate and setup the root thread
	thread_ptr_result_t thd_ret =
		partition_allocate_thread(root_partition, params);
	if (thd_ret.e != OK) {
		panic("Error allocating root thread");
	}
	thread_t *root_thread = (thread_t *)thd_ret.r;

vcpu_configure(root_thread, vcpu_options);

// Attach root cspace to root thread
cspace_attach_thread(root_cspace, root_thread)

//trigger_rootvm_init_event
scheduler_fprr_handle_rootvm_init(qcbor_enc_ctxt);
vgic_handle_rootvm_init(root_partition, root_thread, root_cspace,
                        hyp_env, qcbor_enc_ctxt);
addrspace_handle_rootvm_init(root_thread, root_cspace, qcbor_enc_ctxt);
power_handle_rootvm_init(root_cspace, qcbor_enc_ctxt);
rootvm_package_handle_rootvm_init(root_partition, root_thread,
                                    root_cspace, hyp_env, qcbor_enc_ctxt);
timer_handle_rootvm_init(hyp_env);
trace_standard_handle_rootvm_init(root_partition, root_cspace,
                                    qcbor_enc_ctxt);
vcpu_handle_rootvm_init(root_thread);
vfp_handle_rootvm_init(qcbor_enc_ctxt);
soc_qemu_handle_rootvm_init(root_partition, root_cspace, hyp_env,
                            qcbor_enc_ctxt);

// 运行 vcpu
vcpu_poweron(root_thread, vmaddr_result_ok(hyp_env.entry_ipa),register_result_ok(hyp_env.env_ipa), power_flags);

trigger_rootvm_started_event(root_thread);
```

***rootvm_package***
rootvm package (runtime and app) 初始化过程

rootvm_package_handle_rootvm_init
```C
// map rootVM image to hypervisor's adrspace
pgtable_hyp_map(root_partition, map_range_r.r.base, map_size,
			      (paddr_t)&image_pkg_start,
			      PGTABLE_HYP_MEMTYPE_WRITEBACK, PGTABLE_ACCESS_R,
			      VMSA_SHAREABILITY_INNER_SHAREABLE);

// map the root_thread memory as RW by default.
memextent_construct(
		root_partition, root_cspace, PLATFORM_ROOTVM_LMA_BASE,
		PLATFORM_ROOTVM_LMA_SIZE, PGTABLE_ACCESS_RWX,
		MEMEXTENT_MEMTYPE_ANY, MEMEXTENT_TYPE_SPARSE, false, &me_cap);

// get info from image
rootvm_package_process_image(
		hyp_env, pkg_hdr, map_range_r, addrspace, me, map_size, ipa);
	vmaddr_t runtime_ipa = info.runtime_ipa;
	vmaddr_t app_ipa     = info.app_ipa;
	size_t	 offset	     = info.offset;

// Map all the remaining root VM memory as RW
memextent_map(me, addrspace, ipa, map_attrs,
            addrspace_map_flags_default()) != OK
```

***smccc***
smc 原型

***vcpu***
- vcpu thread 的创建、激活和去激活、退出、状态保存和查询、虚拟中断的绑定和解绑
- vcpu操作，包括 vcpu_configure/vcpu_poweron/vcpu_poweroff/vcpu_suspend/vcpu_resume/vcpu_warm_reset/vcpu_halted
- 异常 trap 的处理逻辑
- 虚拟fp的处理

***vcpu_power***
处理 vcpu 的 power_on/power_off/stop/killed 等逻辑

***vcpu_run***
vcpu运行前检查和运行的hypercall处理逻辑，vcpu与virq的绑定、解绑

***vgic***
- 虚拟gic模拟
    - vic 对象管理
    - vgicd
    - vgicr
    - attach vcpu
    - mpidr handle
- handle rootvm initialization
    - 

***vic***

***viommu***

***virtio***
virtio 前后端协议部分

***virtio_backend***
virtio 后端处理

***virtio_mmio***
基于mmio总线 virtio协议部分

***virtio_pci***
基于pci总线 virtio协议部分

***virtio_virtq***
virtio协议virtqueue部分

***vpci***
虚拟pci

***vpm_base***
virtual pm

***vrtc_pl031***


## rm

### arch


### config


### platform


### src

***device_manager***
hlos 和 svm 的设备管理器

***dt***
设备树处理

***elf***
elf 处理

***event***
异步通知机制

***exit***
退出事件处理

***heap_mgnt***
可读写内存区域管理器

***hyp***
访问hypervisor api 接口封装

***irq_manager***
中断管理器，通过hypervisor接口管理所有中断注入到各个虚拟机

***memparcel***
内存管理器，通过hypervisor接口管理所有内存分配到各个虚拟机

***preempt***

***rpc***
虚拟机间通信机制

***vm_config***
vm配置管理

***vm_console***

***vm_creation***
创建vm及相关资源

***vm_dt***
vm设备树处理

***vm_firmware***
vm启动固件相关

***vm_ipa***
vm中间地址

***vm_memory***
vm内存管理

***vm_mgnt***
vm管理

***vm_passthrough_config***
vm直通配置



## gunyah 深度解析

[gunyah 深度解析](https://blog.csdn.net/weixin_38498942/article/details/157979711?spm=1001.2101.3001.6650.2&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogOpenSearchComplete%7ECtr-2-157979711-blog-157979024.235%5Ev43%5Epc_blog_bottom_relevance_base5&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogOpenSearchComplete%7ECtr-2-157979711-blog-157979024.235%5Ev43%5Epc_blog_bottom_relevance_base5&utm_relevant_index=2)

[avf](https://www.csdn.net/article/2024-02-29/136363923)


## 架构设计

### 内存管理

#### 设计目标

- 提供