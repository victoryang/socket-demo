# Secrets

Kubernetes Secrets let you store and manage sensitive information such as passwords, OAuth tokens, and ssh keys. Storing confidential information in a Secret is safer and more flexible than putting it verbatim in a Pod definition or in a container image.

A secret is an object that contains a small amount of sensitive data such as a password, a token, or a key. Such information might otherwise be put in a Pod specification or in an image. Users can create Secrets and the system also creates some Secrets.

> **Caution:**
> Kubernetes Secrets are, by default, stored as unencrypted base64-encoded strings. By default they can be retrieved - as plain text - by anyone with API access, or anyone with access to Kubernetes' underlying data store, etcd. In order to safely use Secrets, it is recommended you (at a minimum):
> 1. Enable Encryption at Rest for Secrets
> 2. Enable or configure RBAC rules that restrict reading and writing the Secret. Be aware that secrets can be obtained implicitly by anyone with the permission to create a Pod.

## Overview of Secrets

To use a Secret, a Pod needs to reference the Secret. A Secret can be used with a Pod in three ways:

- As files in a volume mounted on one or more of its containers.
- As container environment variable.
- By the kubelet when pulling images for the Pod.

The name of a Secret object must be a valid DNS subdomain name. You can specify the `data` and/or the `stringData` field when creating a configuration file for a Secret. The `data` and the `stringData` fields are optional. The values for all keys in the `data` field have to be base64-encoding strings. If the conversion to base64 string is not desirable, you can choose to specify the `stringData` field instead, which accepts arbitrary strings as values.

The keys of `data` and `stringData` must consist of alphanumeric characters, `-`, `_` or `.`. All key-value pairs in the `stringData` field are internally merged into the `data` field. If a key appears in both the `data` and the `stringData` field, the value specified in the `stringData` field takes precedence.

## Types of Secret

When creating a Secret, you can specify its type using the `type` field of the `Secret` resource, or certain equivalent `kubectl` command line flags (if available). The Secrets type is used to facilitate programmatic handling of the Secret data.

Kubernetes provides several builtin types for common usage scenarios. These types vary in terms of the validations performed and the constraints Kubernetes imposes on them.

|||
|-|-|
|`Opaque`|arbitray user-defined data|
|`kubernetes.io/service-account-token`|service account token|
|`kubernetes.io/dockercfg`|serialized `~/.dockercfg` file|
|`kubernetes.io/dockerconfigjson`|serialized `~/.docker/config.json` file|
|`kubernetes.io/basic-auth`|credentials for basic authentication|
|`kubernetes.io/ssh-auth`|credentials for SSH authentication|
|`kubernetes.io/tls`|data for a TLS client or server|
|`bootstrap.kubernetes.io/token`|bootstrap token data|