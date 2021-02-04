# Knative Build

**Knative Build is deprecated in favor of Tekton Pipelines. There are no plans to produce future releases of this component**

A Knative build extends Kubernetes and utilizes existing Kubernetes primitives to provide you with the ability to run on-cluster container builds from source. For example, you can write a build that uses Kubernetes-native resources to obtain your source code from a repository, build a container image, then run that image.

While Knative builds are optimized for building, testing, and deploying source code, you are still responsible for developing the corresponding components that:

- Retrieve source code from repositories.
- Run multiple sequential jobs against a shared filesystem, for example:
    - Install dependencies
    - Run unit and integration tests.
- Build container images.
- Push container images to an image registry, or deploy them to a cluster.

The goal of Knative build is to provide a standard, portable, resuable, and performance optimized method for defining and running on-cluster container image builds. By providing the "boring but difficult" task of running builds on Kubernetes, Knative saves you from having to independently develop and reproduce these common Kubernetes-based development processes.

While today, a Knative build does not provide a complete standalone CI/CD solution, it does however, provide a lower-level building block that was purposefully designed to enable integration and utilization in larger systems.