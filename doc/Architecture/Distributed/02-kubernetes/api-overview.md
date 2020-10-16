# API Overview

[reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/)

[aliyun](https://zhuanlan.zhihu.com/p/94949824)

[kubernetes api chinese doc](http://docs.kubernetes.org.cn/31.html)

[go client example](https://github.com/kubernetes/client-go/tree/master/examples)

## API Version

- Alpha
- Beta
- Stable

## API Groups

API Groups make it easier to extend the Kubernetes API. The API group is specified in a REST path and in the `apiVerion` field of a serialized object.

- The *core* group is found at REST path `/api/v1`. The core group is not specified as part of the `apiVersion` field, for example, `apiVersion: v1`.
- The named groups are at REST path `/apis/$GROUP_NAME/$VERSION` and use `apiVersion: $GROUP_NAME/$VERSION` (for example, `apiVersion: batch/v1`).

## Enabling or disabling API groups

Certain resources and API groups are enabled by default. You can enable or disable them by setting --runtime-config on the API server. The --runtime-config flag accepts comma separated <key>=<value> pairs describing the runtime configuration of the API server. For example:

- to disable batch/v1, set --runtime-config=batch/v1=false
- to enable batch/v2alpha1, set --runtime-config=batch/v2alpha1

