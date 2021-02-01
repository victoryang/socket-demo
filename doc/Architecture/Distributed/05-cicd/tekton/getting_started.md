# Getting Started

## Prerequisites

- A Kubernetes cluster version 1.15 or higher for Tekton Pipeline v0.11.0 or higher, or a Kubernetes cluster version 1.11 or higher for Tekton releases before v0.11.0
- Enable RBAC in the cluster
- Grant current user the cluster-admin role.

## Persistent Volumes

To run a CI/CD workflow, you need to provide Tekton a Persistent volume for storage purpose. Tekton rquests a volume of 5Gi with the default storage class by default. Your Kubernetes cluster, such as one from Google Kubernetes Engine, may have persistent volume set up at the time of creation, thus no extra step is required; if not, you may have to create them manually.

## Your first CI/CD workflow with Tekton

With Tekton, each operation in your CI/CD workflow becomes a Step, which is executed with a container image you specify. Steps are then organized in Tasks, which run as a Kubernetes pod in your cluster. You can further organize Tasks into Pipelines, which can control the order of execution of serveral tasks.

To create a Task, create a Kubernetes object using the Tekton API with the kind Task. The following YAML file specifies a Task with one simple Step, which prints a Hello World! message using the official Ubuntu image:

```yaml
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: hello
spec:
  steps:
    - name: hello
      image: ubuntu
      command:
        - echo
      args:
        - "Hello World!"
```

Write the YAML above to a file named task-hello.yaml, and apply it to your Kubernetes cluster:

```bash
kubectl apply -f task-hello.yaml
```

To run this task with Tekton, you need to create a TaskRun, which is another Kubernetes object used to specify run time information for a Task.

To view this TaskRun object you can run the following Tekton CLI (tkn) command:

```bash
tkn task start hello --dry-run
```

After running the command above, the following TaskRun definition should be shown:

```yaml
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  generateName: hello-run-
spec:
  taskRef:
    name: hello
```

To use the TaskRun above to start the echo Task, you can either use tkn or kubectl.

Start with kubectl:

```bash
tkn task start hello
```

Start with kubectl:

```bash
# use tkn's --dry-run option to save the TaskRun to a file
tkn task start hello --dry-run > taskRun-hello.yaml
# create the TaskRun
kubectl create -f taskRun-hello.yaml
```

Tekton will now start running your Task. To see the logs of the last TaskRun, run the following tkn command:

```yaml
tkn taskrun logs --last -f
```

It may take a few moments before your Task completes. When it executes, it should show the folloing output:

```bash
[hello] Hello World!
```

## Extending your first CI/CD Workflow with a second Task and a Pipeline

As you learned previously, with Tekton, each operation in your CI/CD workflow becomes a `step`, which is executed with a container image you specify. `Steps` are then organized in `Tasks`, which run as a Kubernetes pod in your cluster. You can further organize `Tasks` into `Pipelines`, which can control the order of execution of serveral `Tasks`.

To create a second `Task`, create a Kubernetes object using the Tekton API with the kind `Task`. The following YAML file specifies a `Task` with one simple `Step`, which prints a `Goodbye World!` message using the official Ubuntu image:

```yaml
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: goodbye
spec:
  steps:
    - name: goodbye
      image: ubuntu
      script: |
        #!/bin/bash
        echo "Goodbye World!"
```

Write the YAML above to a file named `task-goodbye.yaml`, and apply it to your Kubernetes cluster:

```bash
kubectl apply -f task-goodbye.yaml
```

To run this task with Tekton, you need to create a `TaskRun`, which is another Kubernetes object used to specify run time information for a `Task`.

To view this `TaskRun` object you can run the following Tekton CLI (tkn) command

```bash
tkn task start goodbye --dry-run
```

After running the command above, the following TaskRun definition should be shown:

```yaml
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  generateName: goodbye-run-
spec:
  taskRef:
    name: goodbye
```

To use the `TaskRun` above to start the `echo` `Task`, you can either use `tkn` or `kubectl`. Start with `tkn`:

```bash
tkn task start goodbye
```

Start with `kubectl`:

```bash
# use tkn's --dry-run option to save the TaskRun to a file
tkn task start goodbye --dry-run > taskRun-goodbye.yaml
# create the TaskRun
kubectl create -f taskRun-goodbye.yaml
```

Tekton will now start running your Task. To see the logs of the TaskRun, run the following tkn command:

```bash
tkn taskrun logs --last -f
```

It may take a few moments before your Task completes. When it executes, it should show the following output:

```bash
[goodbye] Goodbye World!
```

To create a `Pipeline`, create a Kubernetes objkect using the Tekton API with the kind `Pipeline`. The following YAML file specifies a `Pipeline`.

```yaml
apiVersion: tekton.dev/v1beta1
kind: Pipeline
metadata:
  name: hello-goodbye
spec:
  tasks:
  - name: hello
    taskRef:
      name: hello
  - name: goodbye
    runAfter:
     - hello
    taskRef:
      name: goodbye
```

Write the YAML above to a file named `pipeline-hello-goodbye.yaml`, and apply it to your Kubernetes cluster:

```bash
kubectl apply -f pipeline-hello-goodbye.yaml
```

To run this pipeline with Tekton, you need to create a `pipelineRun`, which is another Kubernetes object used to specify run time information for a `Pipeline`.

To view this `pipelineRun` object you can run the following Tekton CLI (`tkn`) command:

```
tkn pipeline start hello-goodbye --dry-run
```

After running the command above, the following `TaskRun` definition should be shown:

```yaml
apiVersion: tekton.dev/v1beta1
kind: PipelineRun
metadata:
  generateName: hello-goodbye-run-
spec:
  pipelineRef:
    name: hello-goodbye
```

To use the `pipelineRun` above to start the `echo` `Pipeline`, you can either use `tkn` or `kubectl`.

Start with `tkn`:

```bash
tkn pipeline start hello-goodbye
```

Start with `kubectl`:

```bash
# use tkn's --dry-run option to save the pipelineRun to a file
tkn pipeline start hello-goodbye --dry-run > pipelineRun-hello-goodbye.yaml
# create the pipelineRun
kubectl create -f pipelineRun-hello-goodbye.yaml
```

Tekton will now start running your `Pipeline`. To see the logs of the `pipelineRun`, run the following `tkn` command:

```bash
tkn pipelinerun logs --last -f 
```

It may take a few moments before your `Pipeline` completes. When it executes, it should show the following output:

```bash
[hello : hello] Hello World!

[goodbye : goodbye] Goodbye World!
```