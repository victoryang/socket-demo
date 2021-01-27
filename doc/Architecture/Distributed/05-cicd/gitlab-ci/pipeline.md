# Pipeline

## Pipeline schedules

Pipelines are normally run based on certain conditions being met. For example, when a branch is pushed to repository.

### Working with scheduled pipelines

After configuration, Gitlab supports many functions for working with scheduled pipelines.

#### Running manually

#### Taking ownership

Pipelines are executed as a user, who owns a schedule. This influences what projects and other resources the pipeline has access to.

## Triggering a pipeline

- Authentication tokens
- Adding a new trigger
- Revoking a trigger
- Triggering a pipeline

## Pipeline settings

### Git strategy

- `git clone`, which is slower since it clones the repository from scratch for every job, ensuring that the local working copy is always pristine.
- `git fetch`, which is default in GitLab and faster as it re-uses the local working copy (falling back to clone if it doesn’t exist). This is recommended, especially for large repositories.

### Git shallow clone

It is possible to limit the number of changes that GitLab CI/CD fetches when cloning a repository. Setting a limit to git depth can speed up Pipelines execution.

In GitLab 12.0 and later, newly created projects automatically have a default git depth value of 50. The maximum allowed value is 1000.

## Pipeline architectures

- Basic
- Direct Acyclic Graph
- Child/Parent Pipelines

## Pipeline Efficiency

- Speed up your DevOps processes
- Reduces costs
- Shortens the development feedback loop

### Identify bottlenecks and common failures

The easiest indicators to check for inefficient pipelines are the runtimes of the jobs, stages, and the total runtime of the pipeline itself. The total pipeline duration is heavily influenced by the:

- Total number of stages and jobs.
- Dependencies between jobs.
- The "critical path", which represents the minimum and the maxmum pipeline duration.

Additional points to pay attention relate to `Gitlab Runners`:

- Availability of the runners and the resources they are provisioned with:
- Build dependencies and their installation time.
- Container image size.
- Network latency and slow connections.

Pipelines frequently failing unnecessarily also causes slowdowns in the development lifecycle. You should look for problematic patterns with failed jobs:

- Flaky unit tests which fail randomly, or produce unreliable test results.
- Test coverage drops and code quality correlated to that behavior.
- Failures that can be safely ignored, but that halt the pipeline instead.
- Tests that fail at the end of a long pipeline, but could be in an earlier stage, causing deplayed feedback.

### Pipeline analysis

Analyze the performance of your pipeline to find ways to improve efficiency. Analysis can help identity possible blockers in the CI/CD infrastructure. This includes analyzing:

- Job workloads.
- Bottlenecks in the execution times.
- The overall pipeline architecture.

Pipeline analysis can help identify issues with cost efficiency. For example, runners hosted with a paid cloud service may be provisioned with:

- More resources than needed for CI/CD pipelines, wasting money.
- Not enough resources, causing slow runtimes and wasting time.

### Pipeline Insights

The Pipeline success and duration charts give information about pipeline runtime and failed job counts.

Tests like unit tests, integration tests, end-to-end tests, code quality tests, and others ensure that problems are automatically found by the CI/CD pipeline. There could be many pipeline stages involved causing long runtimes.

You can improve runtimes by running jobs that test different things in parallel, in the same stage, reducing overall runtime. The downside is that you need more runners running simultaneously to support the parallel jobs.

### Directed Acyclic Graphs (DAG) visualization

### Pipeline Monitoring

#### Runner monitoring

- Disk and disk IO
- CPU usage
- Memory
- Runner process resources

### Storage usage

Review the storage use of the following to help analyze costs and efficiency:

- Job artifacts and their `expire_in` configuration. If kept for too long, storage usage grows and could slow pipelines down.
- Container registry usage
- Package registry usage

### Pipeline configuration

Make careful choices when configuring pipelines to speed up pipelines and reduce resource usage. This includes making use of GitLab CI/CD's built-in features that make pipelines run faster and more efficiently.

#### Reduce how often jobs run

Try to find which jobs don't need to run in all situations, and use pipeline configuration to stop them from running:

- Use the `interruptible` keyword to stop old pipelines when they are superseded by a newer pipeline.
- Use `rules` to skip tests that aren't needed. For example, skip backend tests when only the frontend code is changed.
- Run non-essential scheduled pipeline less frequently.

#### Fail fast

Ensure that errors are detected early in the CI/CD pipeline. A job that takes a very long time to complete keeps a pipeline from returning a failed status until the job completes.

Design pipelines so that jobs that can fail fast run earlier. For example, add an early stage and move the syntax, style linting, Git commit message verification, and similar jobs in there.

Decide if it's import for long jobs to run early, before faster feedback from faster jobs. The initial failure may make it clear that the rest of the pipeline shouldn't run, saving pipeline resources.

#### Directed Acyclic Graphs (DAG)

In a basic configuration, jobs always wait for all other jobs in earlier stages to complete before running. This is the simplest configuration, but it's also the slowest in most cases. Directed Acyclic Graphs and parent/child pipelines are more flexible and can be more efficient, but can also make pipelines harder to understand and analyze.

#### Caching

Another optimization method is to cache dependencies. If your dependencies change rarely, like NodeJS/node_modules, caching can make pipeline execution much faster.

You can use `cache:when` to cache download dependencies even when a job fails.

#### Docker Images

Downloading and initializing Docker images can be a large part of the overall runtime of jobs.

If a Docker image is slowing down job execution, analyze the base image size and network connection to the registry. If GitLab is running in the cloud, look for a cloud container registry offered by the vendor. In addition to that, you can make use of the GitLab container registry which can be accessed by the GitLab instance faster than other registries.

##### Optimize Docker images

Build optimized Docker images because large Docker images use up a lot of space and take a long time to download with slower connection speeds. If possible, avoid using one large image for all jobs. Use multiple smaller images, each for a specific task, that download and run faster.

Try to use custom Docker images with the software pre-installed. It’s usually much faster to download a larger pre-configured image than to use a common image and install software on it each time. Docker’s Best practices for writing Dockerfiles has more information about building efficient Docker images.

Methods to reduce Docker image size:

- Use a small base image, for example debian-slim.
- Do not install convenience tools like vim, curl, and so on, if they aren’t strictly needed.
- Create a dedicated development image.
- Disable man pages and docs installed by packages to save space.
- Reduce the RUN layers and combine software installation steps.
- Use multi-stage builds to merge multiple Dockerfiles that use the builder pattern into one Dockerfile, which can reduce image size.
- If using apt, add --no-install-recommends to avoid unnecessary packages.
- Clean up caches and files that are no longer needed at the end. For example rm -rf /var/lib/apt/lists/* for Debian and Ubuntu, or yum clean all for RHEL and CentOS.
- Use tools like dive or DockerSlim to analyze and shrink images.

#### Test, document, and learn

## 