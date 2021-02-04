# Kaniko

kaniko is a tool to build container images from a Dockerfile, inside a container or Kubernetes cluster.

kaniko doesn't depend on a Docker daemon and executes each command within a Dockerfile completely in userspace. This enables building container images in environments that can't easily or security run a Docker daemon, such as a standard Kubernetes cluster.

kaniko is meant to be run as an image: `gcr.io/kaniko-project/executor`. We do not recommend running the kaniko executor binary in another images, as it might not work.

## How does kaniko work?

The kaniko executor image is responsible for building an image from a Dockerfile and pushing it to a registry. Within the executor image, we extract the filesystem of the base image (the FROM image in the Dockerfile). We then execute the commands in the Dockerfile, snapshotting the filesystem in userspace after each one. After each command, we append a layer of changed files to the base image (if there are any) and update image metadata.

## Known Issues

- kaniko does not support building Windows containers
- Running kaniko in any Docker image other than the official kaniko image is not supported.
    - This includes copying the kaniko executables from the official image into another image.
- kaniko does not support the v1 Registry API

## Using kaniko

To use kaniko to build and push an image for you, you will need:

- A build context, aka something to build
- A running instance of kaniko

### kaniko Build Contexts

kaniko's build context is very similar to the build context you would send your Docker daemon for an image build; it represents a directory containing a Dockerfile which kaniko will use to build your image. For example, a `COPY` command in your Dockerfile should refer to a file in the build context.

You will need to store your build context in a place that kaniko can access. Right now, kaniko supports these storage solutions:

*Note about Local Directory: this option refers to a directory within the kaniko container. If you wish to use this option, you will need to mount in your build context into the container as a directory*.

*Note about Local Tar: this option refers to a tar gz file within the kaniko container. If you wish to use this option, you will need to mount in your build context into the container as a file.*

*Note about Standard Input: the only Standard Input allowed by kaniko is in `.tar.gz` format*

If you using a GCS or S3 bucket, you will first need to create a compressed tar of your build context and upload it to your bucket. Once running, kaniko will then download and unpack the compressed tar of the build context before starting the image build.

To create a compressed tar, you can run:

```bash
tar -C <path to build context> -zcvf context.tar.gz .
```

Then, copy over the compressed tar into your bucket. For example, we can copy over the compressed tar to a GCS bucket with gsutil:

```bash
gsutil cp context.tar.gz gs://<bucket name>
```

When running kaniko, use the `--context` flag with the appropriate prefix to specify the location of your build context

### Using Standard Input

If running kaniko and using Standard Input build context, you will need to add the docker or kubernetes `-i, --interactive` flag. Once running, kaniko will then get the data from `STDIN` and create the build context as a compressed tar. It will then unpack the compressed tar of the build context before starting the image build. If no data is piped during the interactive run, you will need to send the EOF signal by yourself by press `Ctrl+D`.