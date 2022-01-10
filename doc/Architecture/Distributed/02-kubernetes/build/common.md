# Common

hack/lib/init.sh
build/common.sh
build/lib/release.sh

## hack/lib/init.sh

```bash
# Until all GOPATH references are removed from all build scripts as well,
# explicitly disable module mode to avoid picking up user-set GO111MODULE preferences.
# As individual scripts (like hack/update-vendor.sh) make use of go modules,
# they can explicitly set GO111MODULE=on
export GO111MODULE=off

source "${KUBE_ROOT}/hack/lib/version.sh"
source "${KUBE_ROOT}/hack/lib/golang.sh"
source "${KUBE_ROOT}/hack/lib/etcd.sh"
```

## build/common.sh

```bash
# Constants
readonly KUBE_BUILD_IMAGE_REPO=kube-build
readonly KUBE_BUILD_IMAGE_CROSS_TAG="$(cat "${KUBE_ROOT}/build/build-image/cross/VERSION")"

readonly KUBE_DOCKER_REGISTRY="${KUBE_DOCKER_REGISTRY:-k8s.gcr.io}"
readonly KUBE_BASE_IMAGE_REGISTRY="${KUBE_BASE_IMAGE_REGISTRY:-k8s.gcr.io/build-image}"

# This version number is used to cause everyone to rebuild their data containers
# and build image.  This is especially useful for automated build systems like
# Jenkins.
#
# Increment/change this number if you change the build image (anything under
# build/build-image) or change the set of volumes in the data container.
readonly KUBE_BUILD_IMAGE_VERSION_BASE="$(cat "${KUBE_ROOT}/build/build-image/VERSION")"
readonly KUBE_BUILD_IMAGE_VERSION="${KUBE_BUILD_IMAGE_VERSION_BASE}-${KUBE_BUILD_IMAGE_CROSS_TAG}"

# Make it possible to override the `kube-cross` image, and tag independent of `KUBE_BASE_IMAGE_REGISTRY`
readonly KUBE_CROSS_IMAGE="${KUBE_CROSS_IMAGE:-"${KUBE_BASE_IMAGE_REGISTRY}/kube-cross"}"
readonly KUBE_CROSS_VERSION="${KUBE_CROSS_VERSION:-"${KUBE_BUILD_IMAGE_CROSS_TAG}"}"
```

### output dir
```bash
# Here we map the output directories across both the local and remote _output
# directories:
#
# *_OUTPUT_ROOT    - the base of all output in that environment.
# *_OUTPUT_SUBPATH - location where golang stuff is built/cached.  Also
#                    persisted across docker runs with a volume mount.
# *_OUTPUT_BINPATH - location where final binaries are placed.  If the remote
#                    is really remote, this is the stuff that has to be copied
#                    back.
# OUT_DIR can come in from the Makefile, so honor it.
readonly LOCAL_OUTPUT_ROOT="${KUBE_ROOT}/${OUT_DIR:-_output}"
readonly LOCAL_OUTPUT_SUBPATH="${LOCAL_OUTPUT_ROOT}/dockerized"
readonly LOCAL_OUTPUT_BINPATH="${LOCAL_OUTPUT_SUBPATH}/bin"
readonly LOCAL_OUTPUT_GOPATH="${LOCAL_OUTPUT_SUBPATH}/go"
readonly LOCAL_OUTPUT_IMAGE_STAGING="${LOCAL_OUTPUT_ROOT}/images"
```

### basic setup functions

kube::build::verify_prereqs
```bash
# Verify that the right utilities and such are installed for building Kube. Set
# up some dynamic constants.
# Args:
#   $1 - boolean of whether to require functioning docker (default true)
#
# Vars set:
#   KUBE_ROOT_HASH
#   KUBE_BUILD_IMAGE_TAG_BASE
#   KUBE_BUILD_IMAGE_TAG
#   KUBE_BUILD_IMAGE
#   KUBE_BUILD_CONTAINER_NAME_BASE
#   KUBE_BUILD_CONTAINER_NAME
#   KUBE_DATA_CONTAINER_NAME_BASE
#   KUBE_DATA_CONTAINER_NAME
#   KUBE_RSYNC_CONTAINER_NAME_BASE
#   KUBE_RSYNC_CONTAINER_NAME
#   DOCKER_MOUNT_ARGS
#   LOCAL_OUTPUT_BUILD_CONTEXT
```

### build

```bash
# Set up the context directory for the kube-build image and build it.
function kube::build::build_image()

# Build a docker image from a Dockerfile.
# $1 is the name of the image to build
# $2 is the location of the "context" directory, with the Dockerfile at the root.
# $3 is the value to set the --pull flag for docker build; true by default
# $4 is the set of --build-args for docker.
function kube::build::docker_build()

# If the data container exists AND exited successfully, we can use it.
# Otherwise nuke it and start over.
function kube::build::ensure_data_container()

# Run a command in the kube-build image.  This assumes that the image has
# already been built.
function kube::build::run_build_command()

# Run a command in the kube-build image.  This assumes that the image has
# already been built.
#
# Arguments are in the form of
#  <container name> <extra docker args> -- <command>
function kube::build::run_build_command_ex()


```