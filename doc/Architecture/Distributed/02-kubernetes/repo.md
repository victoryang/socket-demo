# Repository

[kubernetes on github](https://github.com/kubernetes/kubernetes)

## Hierarchy

### Staging

This directory is the staging area for packages that have been split to their own repository. The content here will be periodically published to respective top-level k8s.io repositories.

Kubernetes code uses the repositories in this directory via simlinks in the vendor/k8s.io directory into this staging area. For example, when kubernates code imports a package from the k8s.io/client-go, that import is resolved to `staging/src/k8s.io/client-go` relative to the project root.

```
// pkg/example/some_code.go
package example

import (
  "k8s.io/client-go/dynamic" // resolves to staging/src/k8s.io/client-go/dynamic
)
```

