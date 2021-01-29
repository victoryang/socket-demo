# Jenkins X

## What is Jenkins X

Jenkins X aims to automate and accelerate Continuous Integration and Continuous Delivery for developers on the cloud and/or Kubernetes so they can focus on building awesome software and not waste time figuring out how to use the cloud or manually create and maintain pipelines. 

Jenkins X 3 focus on a few main areas:

### Infrastructure

Moving management of infrastructure outside of Jenkins X, favouring solutions like Terraform. This reduces the surface area of Jenkins X and leverages expert OSS projects and communities around managing infrastructure and cloud resources.

### Secret Management

Adding an abstraction layer above secret management solutions so users can choose where the source of secrets can be stored, preferably outside of the Kubernetes cluster. This is a good practice for disaster recovery scenarios.

### Developer experience

Jenkins X 3.x includes a revived focus on developer experience. The introduction of Jenkins X plugins for Octant has addressed a long standing request from the open source community. Jenkins X 3.x will be focussing on new visualisations to help developers, operators and cross functioning teams.

### Maintainability

Created a new jx CLI which includes an extensible plugin model where each main subcommand off the jx base is it's own releasable git repository. This has significantly improved the Jenkins X codebase which helps with maintainability and contributions.

### Removing complexity and magic

Removing complexity out of Jenkins X and reusing other solutions wherever possible. Jenkins X 2.x was tightlu coupled to helm 2 for example. There were jx CLI steps that wrapped helm commands when installing applications into the cluster which injected secrets from an internal Vault and ultimately made it very confuing for users and maintainers.

Jenkins X 3.x prefers to avoid wrapping other CLIs unless consistent higher level UX is being provided say around managing secrets and underlying commands being executed are clearly printed in users terminals.

