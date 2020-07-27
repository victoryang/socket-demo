# Harbor

[Harbor Doc](https://github.com/goharbor/harbor/blob/master/docs/README.md)

[Architecture Overview of Harbor](https://github.com/goharbor/harbor/wiki/Architecture-Overview-of-Harbor)

## Architecture

From now on, Harbor has been evolved to a complete OCI compliant cloud-native artifact registry.

### Data Access Layer

**k-v storage:** formed by Redis, provides data cache functions and supports temporarily persisting job metadata for the job service.

**data storage**: multiple storages supported for data persistence as backend storage of registry and chart museum. For checking more datails, please refer to the driver list document at [docker storage](https://docs.docker.com/reference/) and [ChartMuseum Github repository](https://github.com/chartmuseum/storage)

**Database**: stores the related metadata of Harbor models, like projects, users, roles, replication policies, tag retention policies, scanners, charts, and images. PostgreSQL is adopted.

### Fundamental Services

**Proxy:** reverse-proxy formed by the Nginx Server to provide API routing capabilities. Components of Harbor, such as core, registry, web portal, and token services, etc., are all behind this reversed proxy. The proxy forwards requests from browsers and Docker clients to various backend services.

**Core:** Harbor's core service, which mainly provides the following functions:

- API Server: A HTTP server accepting REST API requests and responding those requests rely on its submodules such as 'Authentication & Authorization', 'Middleware', and 'API Handlers'.
    - Authentication & Authorization
        - requests are protected by the authentication service which can be powered by a local database, AD/LDAP or OIDC.
        - RBAC mechanism is enabled for performing authorizations to the related actions, e.g: pull/push a image
        - Token service is designed for issuing a token for every docker push/pull command according to a user's role of a project. If there is no token in a request sent from a Docker client, the Registry will redirect the request to the token service.
    - Middleware: Preprocess some requests in advance to determine whether they match the required criteria and can be passed to the backend components for further processing or not. Some functions are implemented as kinds of middleware, such as 'quota management', 'signature check', 'vulernability severity check' and 'robot account parsing' etc.
    - API Handlers: Handle the corresponding REST API requests, mainly focus on parsing and validating requestparameters, completing business logic on top of the relevant API controller, and writing back the generated response.
- Config Manager: Covers the management of all the system configurations, like authentication type settings, email settings, and certificates, etc..
- Project Management: Manages the base data and corresponding metadata of the project, which is created to isolate the managed artifacts.
- Quota Manager: Manages the quota settings of projects and performs the quota validations when new pushed happened.
- Chart Controller: Proxy the chart related requests to backend  `chartmusem` and provides several extensions to improve chart management experiences.
- Retention Manager: Manages the tag retention policies and perform and monitor the tag retention processes.
- Content Trust: add extensions to the trust capability provided by backend Notary to support the smoothly content trust process. At present, only container images are supported to sign.
- Replication Controller: Manages the replication policies and registry adapters, triggers and monitors the concurrent replication processes. Many regsitry adapters are implemented
- Scan Manager: Manages the multiple configured scanners adapted by different providers and also provides scan summaries and reports for the specified artifacts.
- Notification Manager(webhook): A mechanism configured in Harbor so that artifact status changes in Harbor can be populated to the Webhook endpoints configured in Harbor.
- OCI Artifact Manager
- Registry Driver

**Job Service:** General job execution queue service to let other componets/services submit requests of running asynchronous tasks concurrently with simple restful APIs
**Log collector:** Log collector, responsible for collectiong logs of other modules into a single place.
**GC Controller:** magages the online GC schedule settings and start and track the GC progress.
**Chart Museum:** a 3rd party chart repository server providing chart management and access APIs
**Docker Registry:** a 3rd party registry server, responsible for storing Docker images and processing Docker push/pull commands. As Harbor needs to enforce access control to images, the Registry will direct clients to a token service to obtain a valid token for each pull or push request.
**Notary:** a 3rd party content trust server, responsible for securely publishing and verifying content.

