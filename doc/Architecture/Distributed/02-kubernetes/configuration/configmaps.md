# ConfigMaps

A ConfigMap is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.

A ConfigMap allows you to decouple environment-specific configuration from your container images, so that your applications are easily portable.

> **Caution**: ConfigMap does not provide secrecy or encryption. If the data you want to store are confidential, use a Secret rather than a ConfigMap, or use additional (third party)tools to keep your data private.

## Motivation

Use a ConfigMap for setting configuration data separately from application code.

For example, imagine that you are developing an application that you can run on your own computer(for development) and in the cloud(to handle real traffic). You write the code to look in an environment variable named `DATABASE_HOST`. Locally, you set that variable to `localhost`. In the cloud, you set it to refer to a Kubernetes Service that exposes the database component to your cluster. This lets you fetch a container image running in the cloud and debug the exact same code locally if needed.

A ConfigMap is not designed to hold large chunks of data. The data stored in a ConfigMap cannot exceed 1 MiB. If you need to store settings that are larger than this limit, you may want to consider mounting a volume or use a separate database or file service.

## ConfigMap object

A ConfigMap is an API object that lets you store configuration for other objects to use. Unlike most Kubernetes objects that have a `spec`, a ConfigMap has `data` and `binaryData` fields. These fields accept key-value pairs as their values. Both the `data` field and the `binaryData` are optional. The `data` field is designed to contain UTF-8 byte sequences while the `binaryData` field is designed to contain binary data.

The name of a ConfigMap must be a valid DNS subdomain name.

Each key under the `data` or the `binaryData` field must consist of alphanumeric characters, `-`, `_` or `.`. The keys sotred in `data` must not overlap with the keys in the `binaryData` field.

Starting from v1.19, you can add an immutable field to a ConfigMap definition to create an immutable ConfigMap.

## ConfigMaps and Pods

You can write a Pod `spec` that refers to a ConfigMap and configures the container(s) in that Pod based on the data in the ConfigMap. The Pod and the ConfigMap must be in the same namespace.

Here's an example ConfigMap that has some keys with single values, and other keys where the value looks like a fragment of a configruation format.

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
    name: game-demo
data:
    # property-like keys: each key maps to a simple value
    player_initial_lives: "3"
    ui_properties_file_name: "user-interface.properties"

    # file-like keys
    game.properties: |
        enemy.types=aliens,monsters
        player.maximum-live=5
    user-interface.properties: |
        color.good=purple
        color.bad=yellow
        allow.textmode=true
```

There are four different ways that you can use a ConfigMap to configure a container inside a Pod:
1. Inside a container command and args
2. Environment variable for a container
3. Add a file in read-only volume, for the application to read
4. Write code to run inside the Pod that uses the Kubernetes API to read a ConfigMap

These different methods lend themselves to different ways of modeling the data being consumed. For the first three methods, the kubelet uses the data from the ConfigMap when it launches container(s) for a Pod.

The fourth method means you have to write code to read the ConfigMap and its data. However, because you're using the Kubernetes API directly, your application can subscribe to get updates whenever the ConfigMap changes, and react when that happes. By accessing the Kubernetes API directly, this technique also lets you access a ConfigMap in a different namespace.

Here's an example Pod that uses values from `game-demo` to configure a Pod:

```yaml
apiVerison: v1
kind: Pod
metadata:
    name: configmap-demo-pod
spec:
    containers:
        - name: demo
          image: alpine
          command: ["sleep", "3600"]
          env:
            # Define the environment variable
            - name: PLAYER_INITIAL_LIVES # Notice that the case is defferent here from the key name in the ConfigMap
              valueFrom:
                configMapKeyRef:
                    name: game-demo
                    key: player_initial_lives
            - name: UI_PROPERTEIES_FILE_NAME
              valueFrom:
                vonfigMapKeyRef:
                    name: game-demo
                    key: ui_properties_file_name
        volumeMounts:
        - name: config
          mountPath: "/config"
          readOnly: true
    volumes:
    # You set volumes at the Pod level, then mount them into containers inside that Pod
    - name: config
      configMap:
        # Provide the name of the ConfigMap you want to mount.
        name: game-demo
        # An array of keys from the ConfigMap to create as files
        items:
        - key: "game.properties"
          path: "game.properties"
        - key: "user-interface.properties"
          path: "user-interface.properties"
```

A ConfigMap doesn't deifferentiate between single line property values and multi-line file-like values. What matters is how Pods and other objects consume those values.

For this example, defining a volume and mounting it inside the `demo` container as `/config` creates two files, `/config/game.properties` and `/config/user-interface.properties`, even though there are four keys in the ConfigMap. This is because the Pod definition specifies an `item` array in the `volumes` section. If you omit the `items` array entirely, every key in the ConfigMap becomes a file with the same name as the key, and you get 4 files.

## Using ConfigMaps as files from a Pod

To consume a ConfigMap in a volume in a Pod:

1. Create a ConfigMap or use an existing one. Multiple Pods can reference the same ConfigMap.
2. Modify your Pod definition to add a volume under `.spec.volumes[]`. Name the volume anything, and have a `.spec.volumes[].configmap.name` field set to reference your ConfigMap object.
3. Add a `.spec.containers[].volumeMounts[]` to each container that needs the ConfigMap. Specify `.spec.containers[].volumeMounts[].mountPath` to an unused directory name where you would like the ConfigMap to appear.
4. Modify your image or command line so that the program looks for files in that directory. Each key in the ConfigMap `data` map becomes the filename under `mountPath`.

This is an example of a Pod that mounts a ConfigMap in a volume:

```yaml
apiVersion: v1
kind: Pod
metadata:
    name: mypod
spec:
    containers:
    - name: mypod
      image: redis
      volumeMounts:
      - name: foo
        mountPath: "/etc/foo"
        readonly: true
    volumes:
    - name: foo
      configMap:
        name: myconfigmap
```

Each ConfigMap you want to use needs to be referred to in the `.spec.volumes`.

If there are multiple containers in the Pod, then each container needs its own `volumeMounts` block, but only one `.spec.volumes` is needed per ConfigMap.

## Mounted ConfigMaps are updated automatically

When a ConfigMap currently consumed in a volume is updated, projected keys are eventually updated as well. The kubelet checks whether the mounted ConfigMap is fresh on every periodic sync. However, the kubelet uses its local cache for getting the current value of the ConfigMap. The tyoe if the cache is configurable using the `ConfigMapAndSecretChangeDetectionStrategy` field in the KubeletConfigruation struct. A ConfigMap can be either propagated by watch(default), ttl-based, or by redirecting all requests directly to the API server. As a result, the total delay from the Pod can be as long as the kubelet sync period + cache propagation delay, where the cache propagation delay depends on the chosen cache type (it equals to watch propagation delay, ttl of cache, or zero correspondingly).

ConfigMaps consumed as environment variables are not updated automatically and require a pod restart.

## Immutable ConfigMaps

The Kubernetes beta feature *Immutable Secrets and ConfigMaps* provides an option to set individual Secrets and ConfigNaos as immutable. For clusters that extensively use ConfigMaos (at least tens of thousands of unique ConfigMap to Pod mounts), preventing changes to their data has the following advantages:

- protects you from accident(or unwanted) updates that could cause applications outages
- improves performance of your cluster by significantly reducing load on kube-apiserver, by closing watches for ConfigMaps marked as immutable.

This feature is controlled by the `ImmutableEphemeralVolumes` feature gate. You can create an immutable ConfigMap by setting the `immutable` field to `true`. For example:

```yaml
apiVerion: v1
kind: Pod
metatdata:
    ...
data:
    ...
immutable: true
```

Once a ConfigMap is marked as immutable, it is not possible to revert this change nor to mutate the contents of the `data` or the `binaryData` field. You can only delete and recreate the ConfigMap. Because existing Pods maintain a mount point to the deleted ConfigMap, it is recommended to recreate these pods.