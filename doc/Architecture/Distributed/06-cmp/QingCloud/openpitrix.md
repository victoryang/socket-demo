# OpenPitrix

https://github.com/openpitrix/openpitrix

## Overview

Openpitrix is a web-based open-source system to package, deploy and manage different types of applications including Kubernetes application, microservice application and serverless applications into multiple cloud environment such as AWS, Azure, Kubernetes, QingCloud, OpenStack, VMWare etc.

> Definition: Pitrix means the matrix of PaaS and IaaS which makes it easy to develop, deploy, manage applications including PaaS on various runtime environments, i.e., Pitrix = PaaS + IaaS + Matrix. It also means a matrix that contains endless(PI - the Greek letter 'π') applications.

## Features

- Multiple-cloud: Support multiple runtimes, such as AWS, Aliyun, Azure, Kubernetes, QingCloud, OpenStack, VMWare and so on.
- Multiple Apps types: Support a variety of application types including VM-based application, Kubernetes application, micronservice application and serverless application.
- Application Lifecycle Management: Developers can easily create and package applications, make flexible application versioning and publishing, other  than check, test and quick deploy applications through the application marketplace.
- Extendable and Pluggable: The types of runtime and application are highly extendable and pluggable, regardless of what new application type of runtime emerges.
- RBAC for organization: Provide multiple roles including regular users, ISV, developers and admin, admin can also create custom roles and department.

## Use cases

Typically there are serveral use cases for OpenPitrix.

- Deployed as one-stop-shop application management platform in an organization to support multiple cloud systems including hybrid cloud.
- Cloud management platform(CMP) can use OpenPitrix as a component to manage applications in multiple-cloud environment.
- Deployed as application management system in Kubernetes. OpenPitrix is different than Helm, OpenPitrix uses Helm under the hood though. In an organization, people usually want to categorize applications by status such as developing, testing, staging, production; or by departments of their organization, to name a few.

