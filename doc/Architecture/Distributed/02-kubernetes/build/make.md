# Make

Makefile --> build/root/Makefile

## build/root/Makefile

### make all

```bash
# Build code.
#
# Args:
#   WHAT: Directory names to build.  If any of these directories has a 'main'
#     package, the build will produce executable files under $(OUT_DIR)/bin.
#     If not specified, "everything" will be built.
#   GOFLAGS: Extra flags to pass to 'go' when building.
#   GOLDFLAGS: Extra linking flags passed to 'go' when building.
#   GOGCFLAGS: Additional go compile flags passed to 'go' when building.
#
# Example:
#   make
#   make all
#   make all WHAT=cmd/kubelet GOFLAGS=-v
#   make all GOLDFLAGS=""
#     Note: Specify GOLDFLAGS as an empty string for building unstripped binaries, which allows
#           you to use code debugging tools like delve. When GOLDFLAGS is unspecified, it defaults
#           to "-s -w" which strips debug information. Other flags that can be used for GOLDFLAGS 
#           are documented at https://golang.org/cmd/link/

all: generated_files
	hack/make-rules/build.sh $(WHAT)
```

### make release

```bash
# Build a release
# Use the 'release-in-a-container' target to build the release when already in
# a container vs. creating a new container to build in using the 'release'
# target.  Useful for running in GCB.
#
# Example:
#   make release
#   make release-in-a-container

release:
	build/release.sh
```

make quick-release

```bash
# Build a release, but skip tests
#
# Args:
#   KUBE_RELEASE_RUN_TESTS: Whether to run tests. Set to 'y' to run tests anyways.
#   KUBE_FASTBUILD: Whether to cross-compile for other architectures. Set to 'false' to do so.
#   KUBE_DOCKER_REGISTRY: Registry of released images, default to k8s.gcr.io
#   KUBE_BASE_IMAGE_REGISTRY: Registry of base images for controlplane binaries, default to k8s.gcr.io/build-image
#
# Example:
#   make release-skip-tests
#   make quick-release

release-skip-tests quick-release:
	build/release.sh
```

### make cross

```bash
# Cross-compile for all platforms
# Use the 'cross-in-a-container' target to cross build when already in
# a container vs. creating a new container to build from (build-image)
# Useful for running in GCB.
#
# Example:
#   make cross
#   make cross-in-a-container

cross:
	hack/make-rules/cross.sh
```

### make kubectl kube-proxy

```bash
# Add rules for all directories in cmd/
#
# Example:
#   make kubectl kube-proxy

CMD_TARGET = $(filter-out %$(EXCLUDE_TARGET),$(notdir $(abspath $(wildcard cmd/*/))))
$(CMD_TARGET): generated_files
	hack/make-rules/build.sh cmd/$@
```

### make generated_files

```bash
# Produce auto-generated files needed for the build.
#
# Example:
#   make generated_files

generated_files gen_openapi:
	$(MAKE) -f Makefile.generated_files $@ CALLED_FROM_MAIN_MAKEFILE=1
```