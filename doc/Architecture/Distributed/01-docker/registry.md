# Registry

[docker registry 1.0](https://github.com/docker/docker-registry)

[docker registry 2.0](https://github.com/docker/distribution)

[OCI Distribution Specification](https://github.com/opencontainers/distribution-spec)

## Docker Registry 1.0

The Registry is a stateless, highly scalable server side application that stores and lets you distribute Docker images.

You should use the Registry if you want to:

- tightly control where your images are being stored
- fully own your images distribution pipeline
- integrate image storage and distribution tightly into your in-house development workflow

## Docker Registry 2.0

The Docker toolset to pack, ship, store, and deliver content.

|Component|Description|
|:--:|:--:|
|registry|An implemetation of the OCI Distribution Specification|
|libraries|A rich set of libraries for interacting with distribution components|
|documentation|Docker's full documentation set is available|


## Understanding the Registry

A registry is a storage and content delivery system, holding named Docker images, available in different tagged version.

> Example: the image ```distribution/registry```, with tags ```2.0``` and ```2.1```

Storage itself is delegated to drivers. The default storage driver is the local posix filesystem, which is suitable for development or small deployments. Additional cloud-based storage drivers like S3, Microsoft Azure, Ceph Rados, OpenStack Swift and Aliyun OSS are also supported

## Docker Image Specification

### Docker Image Specification v1.0.0

An *Image* is an ordered collection of root filesystem changes and the corresponding execution parameters for use within a container runtime. This specification outlines the format of these filesystem changes and corresponding parameters and describes how to create and use them for use with a container runtime and execution tool.

#### Terminology

This specification uses the following terms:

##### Layers

Images are composed of *layers*. *Image layer* is a general term which may be used to refer to one or both of the following:

1. The metadata for the layer, described in the JSON format.
2. The filesystem changes described by a layer.

To refer to the former you may use the term *Layer JSON* or *Layer Metadata*. To refer to the latter you may use the term *Image Filesystem Changeset* or *Image Diff*.

***Image JSON***

Each layer has an associated JSON structure which describes some basic information about the image such as date created, author, and the ID of its parent image as well as execution/runtime configuration like its entry point, default arguments, CPU/memory shares, networking, and volumes.

***Image Filesystem Changeset***

Each layer has an archive of the files which have been added, changed, or deleted relative to its parent layer. Using a layer-based or union filesystem such as AUFS, or by computing the diff from filesystem snapshots, the filesystem changeset can be used to present a series of image layers as if they were one cohesive filesystem.

***Image ID***

Each layer is given an ID upon its creation. It is represented as a hexadecimal encoding of 256 bits, e.g., `a9561eb1b190625c9adb5a9513e72c4dedafc1cb2d4c5236c9a6957ec7dfd5a9`. Image IDs should be sufficiently random so as to be globally unique. 32 bytes read from /dev/urandom is sufficient for all practical purposes. Alternatively, an image ID may be derived as a cryptographic hash of image contents as the result is considered indistinguishable from random. The choice is left up to implementors.

***Image Parent***

Most layer metadata structs contain a `parent` field which refers to the Image from which another directly descends. An image containers a separate JSON metadata file and set of changes relative to the filesystem of its parent image. *Image Ancestor* and *Image Descendant* are also common terms.

***Image Checksum***

Layer metadata structs contain a cryptographic hash of the contents of the layer's filesystem changeset. Though the set of changes exists as a simple Tar archive, two archives with identical filenames and content will have different SHA digests if the last-access or last-modified times of any entries differ. For this reason, image checksums are generated using the TarSum algorithm which produces a cryptographic hash of file contents and selected headers only.

***Tag***

A tag serves to map a descriptive, user-given name to any single image ID. An image name suffix(the name component after:) is often referred to as a tag as well, but they SHOULD be limited to the set of alphanumeric character`[a-zA-Z0-9]`, punctuation characters `[._-]`, and MUST NOT contain a `:` character.

***Respository***

A collection of tags grouped under a common prefix (the name component`:`). For example, in an image tagged with the name `my-app:3.1.4`, `my-app` is the `Repository` component of the name. Acceptable values for repository name are implementation specific, but they SHOULD be limited to the set of alphanumeric characters `[a-zA-Z0-9]`, and punctuation characters `[._-]`, however it MAY contain additional `/` and `:` characters for organizational puerposes, with the last `:` character being interpreted dividing the repository component of the name from the tag suffix component.


### Image Manifest Version 2, Schema 1

This document outlines the format of of the V2 image manifest. It is a provisional manifest to provide a compatibility with the V1 image format, as the requirements are defined for the V2 Schema 2 image.

Image manifests describe the various constituents of a docker image. Image manifests can be serialized to JSON format with the following media types:

|Manifest Type|Media Type|
|-|-|
|manifest|"application/vnd.docker.distribution.manifest.v1+json"|
|signed manifest|"application/vnd.docker.distribution.manifest.v1+prettyjws"|

#### Manifest Field Descriptions

Manifest provides the base accessible fields for working with V2 image format in the registry.

- `name` string
    the name of the image's repository
- `tag` string
    tag of the image
- `architecture` string
    host arhchitecture on which this image is intended to run. This is for information purposes and not currently used by the engine
- `fsKayers` array
    fsLayers is a list of filesystem layer blob sums contained in this image.
    An fsLayer is a struct consisting of the following fields - `blobSum` digest. Digest 
    ```blobSum is the digest of the referenced filesystem image layer. A digest can be a tarsum or sha256 hash.```

- `history` array
    history is a list of unstructured historical data for v1 compatibility. It contain ID of the image layer and ID of the layer's parent layers.
    history is a struct consisting of the following fields
    - `v1Compatibility` string

- `schemaVersion` int
    SchemaVersion is the image manifest schema that this image follows


## Docker Registry HTTP API V2

[api](https://github.com/docker/distribution/blob/v2.2.1/docs/spec/api.md)

### Introduction

The *Docker Registry HTTP API* is the protocol to facilitate distribution of images to the docker engine. It interacts with instances of the docker registry, which is a service to manage information about docker images and enable their distribution. The specification covers the operation of version 2 of this API, known as *Docker Registry HTTP API V2*.

While the V1 registry protocol is usable, there are several problems with the architecture that have led to this new version. The main driver of this specification these changes to the docker the image format, covered in docker. The new, self-contained image manifest simplifies image definition and improves security. This specification will build on that work, leveraging new properties of the manifest format to improve performance, reduce bandwidth usage and decrease the likelihood of backend corruption.

### Scope

This specification covers the URL layout and protocols of the interaction between docker registry and docker core. This will affect the docker core registry API and the rewrite of docker-registry. Docker registry implementations may implement other API endpoints, but they are not coverd by this specification.

This includes the following features:

- Namespace-oriented URI layout
- PUSH/PULL registry server for V2 image manifest format
- Resumable layer PUSH support
- V2 Client library implementation

While authentication and authorization support will influence this specification, details of the protocol will be left to a future specification. Relevant header definitions and error codes are present to provide an indication of what a client may encounter.

#### Use Cases

For the most part, the use cases of the former registry API apply to the new version. Differentiating use cases are covered below.

**Image Verification**

A docker engine instance would like to run verified image named "library/ubuntu", with the tag "latest". The engine contacts the registry, requesting the manifest for "library/ubuntu:latest". An untrusted registry returns a manifest. Before proceeding to download the individual layers, the engine verifies the manifest's signature, ensuring that the content was produced from a trusted source and no tampering has occured. After each layer is downloaded, the engine verifies the digest of the layer, ensuring that the content matches that speficified by the manifest.

**Resumable Push**

Company X's build servers lose connectivity to docker registry before completing an image layer transfer. After connectivity returns, the build server attempts to re-upload the image. The registry notifies the build server that the upload has already been partially attempted. The build server responds by only sending the remaining data to complete tha image file.

**Resumable Pull**

Company X is having more connectivity problems but this time in their deployment datacenter. When downloading an image, the connection is interrupted before completion. The client keeps the partial data and uses http `Range` requests to avoid downloading repeated data.

**Layer Upload De-duplication**

Company Y's build system creates two identical docker layers from build processes A and B. Build process A completes uploading the layer before B. When process B attemts to upload the layer, the registry indicates that its not necessary because the layer is already known.

If process A and B upload the same layer at the same time, both operations will proceed and the first to complete will be stored in the registry.

#### Changes

f 
- Specify the delete API for layers and manifests

e
- Added support for listing registry contents
- Added pagination to tags API
- Added common apporoach to support pagination

d
- Allow repository name components to be one character.
- Clarified that single component names are allowed

c
- Added section covering digest format
- Added more clarification that manifest cannot be deleted by tag

b
- Added capability of doing streaming upload to PATCH blob upload
- Updated PUT blob upload to no longer take final chunk, now requires entire data or no data
- Removed `416 Requested Range Not Satisfiable` response status from PUT blob upload

a
- Added support for immutable manifest references in manifest endpoints
- Deleting a manifest by tag has been deprecated
- Specified `Docker-Content-Digest` header for appropriate entities
- Added error code for unsupported operations

#### Overview

This section covers client flows and details of the API endpoints. The URI layout of the new API is structured to support a rich authentication and authorization model by leveraging namespaces. All endpoints will be prefixed by the API version and the repository name:

```/v2/<name>/```


**Errors**

Actionable failure conditions, covered in detail in their relevant sections, are reported as part of 4xx responses, in a json response body.

**API VERSION CHECK**

```GET /v2/```

**Content Digests**

 It uniquely identifies content by taking a collision-resistant hash of the bytes. Such an identifier can be independently calculated and verified by selection of a common algorithm. If such an identifier can be communicated in a secure manner, one can retrieve the content from an insecure source, calculate it independently and be certain that the correct content was obtained. Put simply, the identifier is a property of the content.

 To disambiguate from other concepts, we call this identifier a digest. A digest is a serialized hash result, consisting of a algorithm and hex portion. The algorithm identifies the methodology used to calculate the digest. The hex portion is the hex-encoded result of the hash.

***Digest Header***

To provide verification of http content, any response may include a Docker- Content-Digest header.This will include the digest of the target entity returned in the response.
For blobs, this is the entire blob content. For manifests, this is the manifest body without the signature content, also known as the JWS payload.

**Pulling An Image**

An "image" is a combination of a JSON manifest and individual layer files.

The process of pulling an image centers around retrieving these two components.

The first step in pulling an image is to retrieve the manifest. For reference, the relevant manifest fields for the registry are the following:

|field|description|
|-|-|
|name|The name of the image|
|tag|The tag for this version of the image|
|fsLayers|A list of layer descriptors(including tarsum)|
|signature|A JWS used to verify the manifest content|

When the manifest is in hand, the client must verify the signature to ensure the names and layers are valid. Once confirmed, the client will then use the tarsums to download the individual layers. Layers are stored in as blobs in the V2 registry API, keyed by their tarsum digest.

***Pulling an Image Manifest***

The image manifest can be fetched with the following url:

```GET /v2/<name>/manifests/<reference>```

The `name` and `reference` parameters identify the image and are required.

A `404 Not Found` response will be returned if the image is unknown to the registry. If the image exists and the response is successful, the image manifest will be returned, with the following format:

```
   "name": <name>,
   "tag": <tag>,
   "fsLayers": [
      {
         "blobSum": <tarsum>
      },
      ...
    ]
   ],
   "history": <v1 images>,
   "signature": <JWS>
```

The client should verify the returned manifest signature for authenticity before fetching layers.

***Pulling a Layer***

Layers are stored in the blob portion of the registry, keyed by tarsum digest. Pulling a layer is carried out by a standard http request.

```GET /v2/<name>/blobs/<tarsum>```

This endpoint should support aggressive HTTP caching for image layers. Support for Etags, modification dates and other cache control headers should be included. To allow for incremental downloads, Range requests should be supported, as well.

**Pushing An Image**

Pushing an image works in the opposite order as a pull. After assembling the image manifest, the client must first push the individual layers. When the layers are fully pushed into the registry, the client should upload the signed manifest.

***Pushing a Layer***

All layer uploads use two steps to manage the upload process. The first step starts the upload in the registry service, returning a url to carry out the second step. The second step uses the upload url to transfer the actual data. Uploads are started with a POST request which returns a url that can be used to push data and check upload status.

The ```Location``` header will be used to communicate the upload location after each request. While it won't change in the this specification, clients should use the most recent value returned by the API.

***1 Starting An Upload***

```POST /v2/<name>/blobs/uploads/```

To begin the process, a POST request should be issued in the following format:

`POST /v2/<name>/blobs/uploads/`

The parameters of this request are the image namespace under which the layer will be linked. Responses to this request are covered below.

***2 Existing Layers***

The existence of a layer can be checked via a `HEAD` request to the blob store API. The request should be formatted as follow:

`HEAD /v2/<name>/blobs/<digest>`

If the layer with the tarsum specified in `digest` is available, a 200 OK response will be received, with no actual body content (this is according to http specification). The response will look as follows:

```
200 OK
Content-Length: <length of blob>
Docker-Content-Digest: <digest>
```

When this response is received, the client can assume that the layer is already available in the registry under the given name and should take no further action to upload the layer. Note that the binary digests may differ for the existing registry layer, but the tarsums will be guaranteed to match.

***3 Uploading the Layer***

If the POST request is successful, a `202 Accepted` response will be returned with upload URL in the `Location` header:

```
202 Accepted
Location: /v2/<name>/blobs/uploads/<uuid>
Range: bytes=0-<offset>
Content-Length: 0
Docker-Upload-UUID: <uuid>
```

The rest of the upload process can be carried out with the returned url, called the `Upload URL` from the Location header. All responses to the upload url, whether sending data or getting status, will be in this format. Though the URI format (`/v2/<name>/blobs/uploads/<uuid>`) for the `Location` header is specified, clients should treat it as an opaque url and should never try to assemble the it.

***4 Upload Progress***

The progress and chunk coordination of the upload process will be coordinated through the `Range` header. 

[upload in another way](https://developers.google.com/youtube/v3/guides/using_resumable_upload_protocol)

***5 Monolithic Upload***

***6 Chunked Upload***

***7 Completed Upload***

***8 Canceling an Upload***

***9 Errors***

**Deleting a Layer**

A layer may be deleted from the registry via its `name` and `digest`. A delete may be issued with the following request format:

```DELETE /v2/<name>/blobs/<digest>```

**Listing Repositories**

Images are stored in collections, known as a repository, which is keyed by a `name`. A registry instance may contain several repositories. The list of available repositories is made available through the **catalog**

```GET /v2/_catalog```

**Listing Image Tags**

```GET /v2/<name>/tags/list```

**Deleting an Image**

An image may be deleted from the registry via its `name` and `reference`.

```DELETE /v2/<name>/manifests/<reference>```

For deletes, `reference` must be a digest or the delete will fail. If the image exists and has been successfully deleted, the following response will be issued:

```
202 Accepted
Content-Length: None
```

If the image had already been deleted or did not exist, a `404 Not Found` response will be issued instead.

