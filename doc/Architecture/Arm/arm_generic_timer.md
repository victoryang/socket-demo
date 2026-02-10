# Generic Timer

https://developer.arm.com/documentation/102379/0104/What-is-the-Generic-Timer-

https://zhuanlan.zhihu.com/p/696401986

## What is the Generic Timer

The Generic Timer provides a standardized timer framework for Arm cores. The Generic Timer includes a System Counter and set of per-core timers, as shown in the following diagram:

![](images/images_01_System-counter-block-diagram.png)

The System Counter is an always-on device, which provides a fixed frequency incrementing system count. The system count value is broadcast to all the cores in the system, giving the cores a common view of passage of time. The system count value is between 56 bits and 64 bits in width. From Armv8.6-A and Armv9.1-A, the frequency of the count is fixed at 1GHz. Pre-Armv8.6-A, the count frequency was a system design choice, typically in the range of 1MHz to 50MHz.

#### Note

The Generic Timer only measure the passage of time. It does not report the time or date. Usually, an SoC would also contain a Real-Time Clock (RTC) for time and date.

Each core has a set of timers. These timers are comparators, which compare against the broadcast system count that is provided by the System Counter. Software can configure timers to generate interrupts or events in set points in the future. Software can also use the system count to add timestamp, because the system count gives a common reference point for all cores.

In this guide, we will explain the operation and configuration of both the timers and the System Counter.

## The processor timers

The table shows the processor timers:

|||
|-|-|
|Timer name|When is the timer present?|
|EL1 physical timer|Always|
|EL1 virtual timer|Always|

### Count and frequency

The `CNTPCT_EL0` system register reports the current system count value.

Reads of `CNTPCT_EL0` can be made speculatively. This means that they can be read out of order regarding the program flow. This could be important in some cases, for example comparing timestamps. When the ordering of the counter read is important, an `ISB` can be used, as the following code shows:

```
loop:           // Polling for some communication to indicate a requirement to read
                // the timer
  LDR X1, [X2]
  CBZ x1, loop
  ISB           // Without this, the CNTPCT could be read before the memory location in
                // [X2] has had the value 0 written to it
  MRS X1, CNTPCT_EL0
```

`CNTFRQ_EL0` reports the frequency of the system count. However, this register is not populated by hardware. The register is write-able at the highest implemented Exception level and readable at all Exception levels. Firmware, typically running at El3, populates this register as part of early system initialization. Higher-level software, like an operating system, can then use the register to get the frequency.

### Timer registers

### Accessing the timers

For some timers, it is possible to configure which Exception levels can access the timer:

- EL1 Physical and Virtual Timers:
  - EL0 access to these timers is controlled by `CNTKCTL_EL1`.
- EL2 Physical and Virtual Timers:
  - When `HCR_EL2.{TGE,E2H}=={1,1}`,EL0 access to these timers is controlled by `CNTKCTL_EL2`. These timers were added as part of the support for the Armv8.1-A Virtualization Host Extension
- EL3 Physical Timer:
  - Access from s.EL1 can be trapped to EL3 using `SCR_EL3.ST`
  - From Armv8.4-A, when `SCR_EL3.EEL2==1`(Secure El2 enabled),then time is inaccessible from S.EL1 or S.EL2. Attempts to access the time from S.EL1 or S.EL2 will result in an UNDEF.

### Configuring a timer

There are two ways to configure a timer, either using the comparator (CVAL) register, or using the timer (TVAL) register.

The comparator register, `CVAL`, is a 64-bit register. Software writes a value to this register and the timer trigger when the count reaches, or exceeds, that value, as you can see here:

```
Timer Condition Met: CVAL <= System Count
```

The timer register, `TVAL`, is a 32-bit register. When software writes `TVAL`, the processor reads the current system count internally, adds the written value, and then populates `CVAL`:

```
CVAL = TVAL + System Counter 
Timer Condition Met: CVAL <= System Count
```

You can see this populating of `CVAL` in software. If you read the current system count, write 1000 to `TVAL`, and then read `CVAL`, you will see that `CVAL` is approximately 1000+system count. The count is approximate, because time will move on during the instruction sequence.

Reading `TVAL` back will show it decrementing down to 0, while the system count increments. `TVAL` reports a signed value, and will continue to decrement after the timer fires, which allows software to determine how long ago the timer fired. `TVAL` and `CVAL` gives software two different models for how to use the timer. If software needs a timer event in X ticks of the clock, software can write X to `TVAL`. Alternatively, if software wants an event when the system count reaches Y, software can write Y to `CVAL`.

Remember that `TVAL` and `CVAL` are different ways to program the same timer. They are not two different timers.

### Interrupts

Timers can be configured to generate an interrupt. The interrupt from a core's timer can only be delivered to that core. This means that the timer of one core cannot be used to generate an interrupt that targets a different core.

The generation of interrupts is controlled through the `ctl` register, using these fields:

- `ENABLE` - Enables the timer.
- `IMASK` - Interrupt mask. Enables or disables interrupt generation.
- `ISTATUS` - When `ENABLE`==1, reports whether the timer is firing (Cval <= System Count).

To generate an interrupt, software must set `ENABLE` to 1 and clear `IMASK` to 0. When the timer fires (CVAL <= System Count), an interrupt signal is asserted to the interrupt controller. In Armv8-A systems, the interrupt controller is usually a Generic Interrupt Controller (GIC).

The interrupt ID(INTID) that is used for each timer is defined by the Server Base System Architecture (SBSA), shown here:

|Timer|SBSA recommended INTID|
|-|-|
|EL1 Physical Timer|30|
|EL1 Virtual Timer|27|
|Non-secure EL2 Physical Timer|26|
|Non-secure EL2 Virtual Timer|28|
|EL3 Physical Timer|29|
|Secure EL2 Physical Timer|20|
|Secure EL2 Virtual Timer|19|

#### Notes

These INTIDs are in the Private Peripheral Interrupt (PPI) range. These INTIDs are private to a specific core. This means that each core sees its EL1 physical timer as INTID 30. This is described in more detail in our Generic Interrupt Controller guide.

The interrupts generated by the timer behave in a level-sensitive manner. This means that, once the timer firing condition is reached, the timer will continue to signal an interrupt until one of the following situation occurs:

- `IMASK` is set to one, which masks the interrupt.
- `ENABLE` is clear to 0, which disables the timer.
- `TVAL` or `CVAL` is written, so that firing condition is no longer met.

When writing an interrupt handler for the timers, it is important that software clears the interrupt before deactivating the interrupt in the GIC. Otherwise the GIC will re-signal the same interrupt again.

The operation and configuration of the GIC is beyond the scope of this guide.

### Timer virtualization

Earlier, we introduced the different timers that are found in a processor. These timers can be divided into two groups: virtual timers and physical timers.

Physical timers, like the EL3 physical timer, CNTPS, compare against the count value provided by the System Counter. This value is referred to as the physical count and is reported by `CNTPCT_EL0`.

Virtual timers, like the EL1 Virtual Timer, CNTV, compare against a virtual count. The virtual count is calculated as:

`Virtual Count = Physical Count - <offset>`

The offset value is specified in the register CNTOFF_EL2, which is only accessible at EL2 or EL3.

If EL2 not implemented, the offset is fixed as 0. This means that the virtual and phsical count values are always the same.

The virtual count allows a hypervisor to show virtual time to a Virtual Machine (VM). For example, a hypervisor could use the offset to hide the passage of time when the VM was not scheduled.

This means that the virtual count can represent time experienced by the VM, rather than wall clock time.

### Event stream

The Generic Timer can also be used to generate an event stream as part of the Wait for Event mechanism. The `wfe` instruction puts the core into a low power state, with the core woken by an event.

Details about the WFE mechanism are beyond the scope of this guide.

There are several ways to generate an event, including:

- Executing the `sev` (Send Event) instruction on a different core
- Clearing the Global Exclusive Monitor of the core
- Using the Event stream from the core's the Generic Timer

The Generic Timer can be configured to generate a stream of events at a regular interval. One use for this configuration is to generate a timeout. `wfe` is typically used when waiting for a resource to become available, when the wait is not expected to be long. The event stream from the timers means that the maximum time that the core will stay in the low power state is bounded.

An event stream can be configured from the physical count, `CNTPCT_EL0`, or from the virtual count, `CNTPVT_EL0`:

`CNTKCTL_EL1` - Controls event stream generation from `CNTVCT_EL0` `CNTKCTL_EL2` - Controls event stream generation from `CNTPCT_EL0`

For each register, the controls are:

- `EVNTEN` Enables or disables the generation of events
- `EVNTI` Controls the rate of events
- `EVNTDIR` Controls when the event is generated

The control `EVENTI` specifies a bit position in the range 0 to 15. When the bit at the selected position changes, an event is generated. For example, if `EVENTI` is set to 3 then an event is generated when bit[3] of the count changes.

The control `EVNTDIR` controls whether the event is generated when the selected bit transitions from 1-to-0 from 0-to-1.