# OS boot

## Example

### littlekernel

- drop down to EL1
- enable caches so atomics and spinlocks work
- setup mmu according to mmu_initial_mappings
- load stack pointer
- clear bss
- load boot args
- call lk_main

### 