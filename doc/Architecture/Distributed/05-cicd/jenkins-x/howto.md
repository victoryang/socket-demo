# Jenkins X How it works

The GitOps repository templates contain the source code, script and docs to help you get your cloud resources created(e.g. a kubernetes cluster and maybe buckets and/or a secret manager).

Once you have created the GitOps repository from one of the available templates and follow the instructions to set up your infrastructure you install the git operator via the jx admin operator command:

```bash
jx admin operator
```

That command essentially installs the git operator chart, passing in the git URL, username and token to run the boot process.

## Git Operator

The git operator works by polling the git repository looking for changes and running a kubernetes job on each change. The Job resource is defined inside the git repository at versionStream/git-operator/job.yaml

You can view the boot Job log via the command:

```bash
jx admin log
```

or you can browse the log in the Octant UI in the operations tab.

## Boot Job

The boot job runs on startup and on any git commit to the GitOps repository you used to install the operator.

The boot job is defined in verionStream/git-operator/job.yaml in git and essentially:

- Runs the generate step
- Runs the apply step

## Generate step

This step is run in the following situations:

- On startup
- After each commit in a Pull Request
- Whenever a commits is made to the main branch which isn't a merge of a Pull Request merge

```