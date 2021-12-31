# Testing in QEMU

## Testing with "make check"

The "make check" testing family includes most of the C based tests in QEMU. For a quick help, run make check-help from the source tree.

The usual way to run these tests is:
```bash
make check
```

which includes QAPI schema tests, unit tests, QTests and some iotests. Different sub-types of "make check" tests will be explained below.

Before running tests, it is best to build QEMU programs first. Some tests expect the executables to exist and will fail with obscure messages if they cannot find them.

### Unit tests

Unit tests, which can be invoked with make check-unit, are simple C tests that typically link to individual QEMU object files and exercise them by calling exported functions.

If you are writing new code in QEMU, consider adding a unit test, especially for utility modules that are relatively stateless or have few dependencies. To add a new unit test:

1. Create a new source file. For example, tests/unit/foo-test.c
2. Write the test. Normally you would include the header file which exports the module API, then verify the interface behaves as expected from your test. The test code should be organised with the glib testing framework. Copying and modifying an existing test is usually a good idea.
3. Add the test to tests/unit/meson.build. the unit tests are listed in a dictionary called tests. The values are any additional sources and dependencies to be linked with the test. For a simple test whose source is in tests/unit/foo-test.c, it is enough to add an entry like:
```
{
    ...
    'foo-test':[],
    ...
}
```

Since unit tests don’t require environment variables, the simplest way to debug a unit test failure is often directly invoking it or even running it under gdb. However there can still be differences in behavior between make invocations and your manual run, due to $MALLOC_PERTURB_ environment variable (which affects memory reclamation and catches invalid pointers better) and gtester options. If necessary, you can run

```
make check-unit V=1
```

and copy the actual command line which executes the unit test, then run it from the command line.

### QTest

QTest is a device emulation testing framework. It can be very useful to test device models; it could also control certain aspect of QEMU (such as virtual clock stepping), with a special purpose "qtest" protocol.

QTest cases can be executed with
```
make check-qtest
```

### QAPI schema tests

The QAPI schema tests validate the QAPI parser used by QMP, by feeding predefined input to the parser and comparing the result with the reference output.

### check-block

## QEMU iotests

### Writing a new test case