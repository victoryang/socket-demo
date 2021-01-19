# Traefik

## Getting Started

### Configuration Introduction

Configuration in Traefik can refer to two different things:

- The fully dynamic routing configuration (referred to as the *dynamic configuration*)
- The startup configuration (referred to as the *static configuration*)

Elements in the *static configuration* set up connections to providers and define the entrypoints Traefik will listen to (these elements don't change often)

The *dynamic configuration* contains everything that defines how the requests are handled by your system. This configuration can changed and is seamlessly hot-reloaded, without any request interruption or connection loss.

### The Dynamic Configuration 

Traefik gets its dynamic configuration from providers: whether an orchestrator, a service registry, or a plain old configuration file.

### The Static Configuration

There are three different, mutually exclusive (e.g. you can use only one at the same time), ways to define static configuration options in Traefik:

1. In a configuration file
2. In the command-line argument
3. As environment variables

