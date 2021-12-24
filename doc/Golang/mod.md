# Go Modules Reference

https://go.dev/ref/mod

## Introduction

Modules are how Go manages dependencies.

## Modules, packages, and versions

A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.

A module is identified by a module path, which is declared in a go.mod file, together with information about the module's dependencies. The module root directory is the directory that contains the go.mod file. The main module is the module containing the directory where the go command is invoked.

Each package within a module is a collection of source files in the same directory that are compiled together. A package path is the module path joined with the subdirectory containing the package(relative to the module root). For example, the module "golang.org/x/net" contains a package in the directory "html". That package's path is "golang/x/net/html".

### Module paths

A module path is the canonical name for a module, declared with the module directive in the module's go.mod file. A module's path is the prefix for package paths within the module.

A module path should describe both what the module does and where to find it. Typically, a module path consists of a repository root path, a directory within the repository (usually empty), and a major version suffix (only for major version 2 or higher).

- The repository root path is the portion of the module path that corresponds to the root directory of the version control repository where the module is developed. Most modules are defined in their repository's root directory, so this is usually the entire path. For example, golang.org/x/net is the repository root path for the module of the same name. See [Finding a repository for a module path](https://go.dev/ref/mod#vcs-find) for information on how the go command locates a repository using HTTP requests derived from a module path.
- If the module is not defined in the repository's root directory, the module subdirectory is the part of the module path that names the directory, not including the major version suffix. This also serves as a prefix for semantic version tags. For example, the module golang.org/x/tools/gopls is in the gopls subdirectory of the repository with root path golang.org/x/tools, so it has the module subdirectory gopls. See [Mapping versions to commits] and Module directories within a repository.
- If the module is released at major version 2 or higher, the module ath must end with a major version suffix like /v2. This may or may not be part of the repository golang.org/x/repo.

If a module might be depended on by other modules, these rules must be followed so that the go command can find and download the module. These are also several lexical restrictions on characters allowed in module paths.

### Versions

A version identifies an immutable snapshot of a module, which may be either a release or a pre-release. Each version starts with the letter v, followed by a semantic version. See semantic Versioning 2.0.0 for details on how versions are formatted, interpreted, and compared.

To summarize, a semantic version consists of three non-negative integers(the major,minor, and patch versions, from left to right) seperated by dots. The patch version may be followed by an optional pre-release string starting with a hypthen. The pre-release string or patch version may be followed by a build metadata string starting with a plus. For example, v0.0.0, v1.12.134, v8.0.5-pre and v2.0.9+meta are valid versions.

A version is considered unstable if its major version is 0 or it has a pre-release suffix. Unstable versions are not subject to compatibility requirements. For example, v0.2.0 may not be compatible with v0.1.0, and v1.5.0-beta may not be compatible with v1.5.0.

Go may access modules in version control systems using tags, branches, or revisions that don’t follow these conventions. However, within the main module, the go command will automatically convert revision names that don’t follow this standard into canonical versions. The go command will also remove build metadata suffixes (except for +incompatible) as part of this process. This may result in a pseudo-version, a pre-release version that encodes a revision identifier (such as a Git commit hash) and a timestamp from a version control system. For example, the command go get -d golang.org/x/net@daa7c041 will convert the commit hash daa7c041 into the pseudo-version v0.0.0-20191109021931-daa7c04131f5. Canonical versions are required outside the main module, and the go command will report an error if a non-canonical version like master appears in a go.mod file.

### Pseudo-versions

A pseudo-version is a specially formatted pre-release version the encodes information about a specific revision control repository. For example, v0.0.0-20191109021931-daa7c04131f5 is a pseudo-version.

Pseudo-versions may refer to revisions for which no semantic version tags are available. They may be used to test commits before creating version tags, for example, on a development branch.

Each pseudo-version has three parts:

- A base version prefix (vX.0.0 or vX.Y.Z-0), which is either derived from a semantic version tag tht precedes the revision or vX.0.0 if there is no such tag.
- A timestamp (yyyymmddhhmmss), which is the UTC time the revision was created. In Git, this is the commit time, not the author time.
- A revision identifier (abcdefabcdef), which is a 12-character prefix of the commit hash, or in Subversion, a zero-padded revision number.

### Major version suffixes

Starting with major version 2, module paths must have a major version suffix like /v2 that matches the major version. For example, if a module has the path example.com/mod at v1.0.0, it must have the path example.com/mod/v2 at version v2.0.0.

Major version suffixes implement the import compatibility rule:
    If an old package and a new package have the same import path, the new package must be backwards compatible with old package.

### Resolving a package to a module

When the go command loads a package using a package path, it needs to determine which module provides the package.

The go command starts by searching the build list for modules with paths that are prefixes of the package path. For example, if the package example.com/a/b is imported, and the module example.com/a is in the build list, the go command will check whether example.com/a contains the package, in the directory b. At least one file with the .go extension must be present in a directory for it to be considered a package. Build constraints are not applied for this purpose. If exactly one module in the build list provides the package, that module is used. If no modules provide the package or if two or more modules provide the package, the go command reports an error. The -mod=mod flag instructs the go command to attempt to find new modules providing missing packages and to update go.mod and go.sum. The go get and go mod tidy commands do this automatically.

When the go command looks up a new module for a package path, it checks the GOPROXY environment variable, which is a comma-separated list of proxy URLs or the keywords direct or off. A proxy URL indicates the go command should contact a module proxy using the GOPROXY protocol. direct indicates that the go command should communicate with a version control system. off indicates that no communication should be attempted. The GOPRIVATE and GONOPROXY environment variables can also be used to control this behavior.

### go.mod files

A module is defined by a UTF-8 encoded text file named go.mod in its root directory. The go.mod file is line-oriented. Each line holds a single directive, made up of a keyword followed by arguments.

The go.mod file is designed to be human readable and machine writable. The go command provides several subcommands that change go.mod files. For example, go get can upgrade or downgrade specific dependencies. Commands that load the module graph will automatically update go.mod when needed. go mod edit can perform low-level edits. The golang.org/x/mod/modfile package can be used by Go programs to make the same changes programmatically.

A go.mod file is required for the main module, and for any replacement module specified with a local file path. However, a module that lacks an explicit go.mod file may still be required as a dependency, or used as a replacement specified with a module path and version; see Compatibility with non-module repositories.

### Lexical elements

### Module paths and versions

Most identifies and strings in a go.mod file are either module path or versions.

### Grammar

### module directive

### Deprecation

### go directive

### require directive

### exclude directive

### replace directive

### retract directive

### Automatic updates

Most commands report an error if go.mod is missing information or doesn’t accurately reflect reality. The go get and go mod tidy commands may be used to fix most of these problems. Additionally, the -mod=mod flag may be used with most module-aware commands (go build, go test, and so on) to instruct the go command to fix problems in go.mod and go.sum automatically.

## Minimal version selection (MVS)

### Replacement
### Exclusion
### Upgrades
### Downgrade

## Module graph pruning

## Compatible with non-module repositories

### +incompatible versions

### Minimal module compatibility

## Module-aware commands

Most go commands may run in Module-aware mode or GOPATH mode. In module-aware mode, the go command uses go.mod files to find versioned dependencies, and it typically loads packages out of the module cache, downloading modules if they are missing. In GOPATH mode, the go command ignores modules; it looks in vendor directories and in GOPATH to find dependencies.

As of Go 1.16, module-aware mode is enabled by default, regardless of whether a go.mod file is present. In lower versions, module-aware mode was enabled when a go.mod file was present in the current directory or any parent directory.

Module-aware mode may be controlled with the GO111MODULE environment variable, which can be set to on,off,or auto.

- The -mod flag controls whether go.mod may be automatically updated and whether the vendor directory is used.
    - -mod=mod tells the go command to ignore the vendor directory and to automatically update go.mod, for example, when an imported package is not provided by any known module.
    - -mod=readonly tells the go command to ignore the vendor directory and to report an error if go.mod needs to be updated.
    - -mod=vendor tells the go command to use the vendor directory. In this mode,the go command will not use the network or the module cache.
    - By default, if the go version in go.mod is 1.14 or higher and a vendor directory is present, the go command acts as if -mod=vendor were used. Otherwise, the go command acts as if -mod=readonly were used.

### Vendoring

## Module proxies

### GOPROXY protocol

### Communicating with proxies

### Serving modules directly from a proxy

## Version control systems

### Finding a repository for a module path

### Mapping versions to commits

### Mapping pseudo-versions to commits

### Mapping branches and commits to versions

### Module directories within a repository

### Controlling version control tools with GOVCS