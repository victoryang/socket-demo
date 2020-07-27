# Harbor Admin

[official website](https://goharbor.io/docs/2.0.0/administration/)

## Configure Authentication

- Database Authentication
- LDAP/AD Authentication
- OIDC Provider Authentication

## Managing Users

Harbor manages images through projects. You provide access to these images to users by including the users in projects and assigning one of the following roles to them.

### Limited Guest

A Limited Guest does not have full read privileges for a project. They can pull images but cannot push, and they cannot see logs or the other members of a project. For example, you can create limited guests for users from different organizations who share access to a project.

### Guest

Guest has read-only privilege for a specified project. They can pull and retag images, but cannot push.

### Developer

Developer has read and write privileges for a project.

### Master

Master has elevated permissions beyond those of 'Developer' including the ability to scan images, view replications jobs, and delete images and helm charts.

### ProjectAdmin

When creating a new project, you will be assigned the "ProjectAmin" role to the project. Besides read-write privileges, the "ProjectAdmin" also has some management privilges, such as adding and removing members, starting a vulnerability scan.

### System-level roles

#### Harbor system administrator

Has the most privileges. In addition to the privileges mentioned above, "Harbor system administrator" can also list all projects, set an ordinary user as administrator, delete users and set vulnerability scan policy for all images. The public project "library" is also owned by the administrator.

#### Anonymous

When a user is not logged in, the user is considered as an "Anonymous" user. An anonymous user has no access to private projects and has read-only access to public projects.

## Configure Global Settings

## Configure Project Quotas

When setting project quotas, it is useful to know how Harbor calculates storage use, especially in relation to image pushing, retagging, and garbage collection.

- Harbor computes image size when blobs and manifests are pushed from the Docker client.

- Shared blobs are only computed once per project. In Docker, blob sharing is defined globally. In Harbor, blob sharing is defined at the project level. As a consequence, overall storage usage can be greater than the actual disk capacity.

- Retagging images reserves and releases resources:
    - If you retag an image within a project, the storage usage does not change because there are no new blobs or manifests.
    - If you retag an image from one project to another, the storage usage will increase.

- During garbage collection, Harbor frees the storage used by untagged blobs in the project.

- Helm chart size is not calculated.

## Configuring Replication

Replication allows users to replicate resources, namely images and charts, between Harbor and non-Harbor registries, in both pull or push mode.

When the Harbor system administrator has set a replication rule, all resources that match the defined filter patterns are replicated to the destination registry when the triggering condition is met. Each resource that is replicated starts a replication task. If the namespace does not exist in the destination registry, a new namespace is created automatically. If it already exists and the user account that is configured in the replication policy does not have write privileges in it, the process fails. Member information is not replicated.

There might be some delay during replication based on the condition of the network. If a replication taks fails, it is re-scheduled for a few minutes later and retried serveral times.

## Vulnerability Scanning

Harbor provides static analysis of vulnerabilities in images through the open source projects Trivy and Clair. To be able to use Trivy, Clair or both you must have enabled Trivym, Clair or both when you installed your Harbor instance.

You can also connect Harbor to your own instance of Trivy/Clair or to other additional vulnerability scanners through Harbor’s embedded interrogation service. These scanners can be configured in the Harbor UI at any time after installation. 

## Garbage Collection

