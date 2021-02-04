# Lighthouse

[github](https://github.com/jenkins-x/lighthouse)

[prow](https://github.com/kubernetes/test-infra/tree/master/prow)

[jenkins-x/go-scm](https://github.com/jenkins-x/go-scm)

Lighthouse is a lightweight ChatOps based webhook handler which can trigger Jenkins X Pipelines, Tekton Pipelines or Jenkins Jobs based on webhooks from multiple git providers such as GitHub, GitHub Enterprise, BitBucket Server and GitLab.

## Installing

Lighthouse is bundled and released as Helm Chart. You find the install instruction in the Chart's README.

Depending on the pipeline engine you want to use, you can find more detailed instructions in one of the following documents:

- Lighthouse + Tekton
- Lighthouse + Jenkins

## Background

Lighthouse derived originally from `Prow` and started with a copy of its essential code.

Currently, Lighthouse supports the standard Prow plugins and handles push webhooks to branches to then trigger a pipeline execution on the agent of your choice.

Lighthouse uses the same `condif.yaml` and `plugin.yaml` for configuration than Prow.

## Comparison to Prow

Lighthouse reuses the Prow plugin source code and a bunch of plugins from Prow

Its got a few differences though:

- rather than being GitHub specific Lighthouse uses jenkins-x/go-scm so it can support any Git provider
- Lighthouse does not use a `ProwJob` CRD; instead, it has its own `LighthouseJob` CRD.

## Porting Prow commands

If there are any prow commands you want which we've not yet ported over, it is relatively easy to port Prow plugins.

We've reused the prow plugin code and configuration code; so it is mostly a case of switching imports of `k8s.io/test-infra/prow` to `github.com/jenkins-x/lighthouse/pkg/prow`, then modifying the GitHub client struct from, say, `github.PullRequest` to `scm.PullRequest`.

Most of the GiHub structs map 1-1 to the jenkins-x/go-scm equivalents (e.g. Issue, Commit, PullRequest). However, the go-scm API does tend to return slices to pointers to resources by default. There are some naming differences in different parts of the API as well. For example, compare the `githubClient` API for Prow lgtm versus the Lighthouse lgtm.

## Development

### Building

To build the code, fork and clone this git repository, then type:

```bash
make build
```

`make build` will build all relevant Lighthouse binaries natively for your OS which you then can run locally. For example, to run the webhook controller, you would type:

```bash
./bin/webhook
```

To see which other Make rules are available, run:
```bash
make help
```

### Environment variables

While Prow only supports GitHub as SCM provider, Lighthouse supports serveral Git SCM providers. Lighthouse achieves the abstraction over the SCM provider using the go-scm library. To configure your SCM, go-scm uses the following environment variables:

