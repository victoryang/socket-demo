# Runners

## Install

```bash
gitlab-runner install --user=gitlab-runner --working-directory=/home/gitlab-runner

gitlab-runner start
```

## Configuration



## Configuring runners in GitLab

In GitLab CI/CD, runners run the code defined in `.gitbal-ci.yml`. A runner is a lightweight, highly-scalable agent that picks up a CI job through the coordinator API of GitLab CI/CD, runs the job, and sends the result back to the GitLab instance.

Runners are created by an administrator and are visible in the GitLab UI. Runners can be specific to certain projects or available to all projects. This documentation is focused on using runners in GitLab. If you need to install and configure GitLab Runner, see the GitLab Runner documentation.

## Types of runners

In the GitLab UI there are three types of runners, based on who you want to have access:

- Shared runners are available to all groups and projects in a GitLab instance.
- Group runners are available to all projects and subgroups in a group.
- Specific runners are associated with specific projects. Typically, specific runners are used for one project at a time.

### Shared runners

