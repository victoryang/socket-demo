# Rust

[CS140E operating system](https://github.com/dddrrreee/cs140e-20win)

## Installation

### Official Doc
https://www.rust-lang.org/learn/get-started

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

### Update

```bash
rustup update
```

## Cargo

### Cargo guide

https://doc.rust-lang.org/cargo/index.html

#### Advantages

- Introduces two metadata files with various bits of package information
- Fetches and builds your package’s dependencies
- Invokes `rustc` or another build tool with the correct parameters to build your package
- Introduces conventions to make working with Rust packages easier

#### Create a New Package

To start a new package with Cargo, use `Cargo new`:

```bash
# create a executable binary program
cargo new hello_world --bin

# a library
cargo new hello_world --lib

# release
cargo build --release
```

This also initializes a new git repository by default.

#### Dependencies

https://crates.io/

crates.io is the Rust community's central package registry that serves as a location to discover and download packages. `cargo` is configured to use it by default to find requested packages.

To depend on a library hosted on crates.io, add it to your `Cargo.toml`

`Cargo.lock` contains the exact information about which revision of all of these dependencies we used.

#### Package Layout

https://doc.rust-lang.org/cargo/guide/project-layout.html

```bash
.
├── Cargo.lock
├── Cargo.toml
├── src/
│   ├── lib.rs
│   ├── main.rs
│   └── bin/
│       ├── named-executable.rs
│       ├── another-executable.rs
│       └── multi-file-executable/
│           ├── main.rs
│           └── some_module.rs
├── benches/
│   ├── large-input.rs
│   └── multi-file-bench/
│       ├── main.rs
│       └── bench_module.rs
├── examples/
│   ├── simple.rs
│   └── multi-file-example/
│       ├── main.rs
│       └── ex_module.rs
└── tests/
    ├── some-integration-tests.rs
    └── multi-file-test/
        ├── main.rs
        └── test_module.rs
```

#### Cargo.toml vs Cargo.lock

- `Cargo.toml` is about describing your dependencies in a broad sense, and is written by you
- `Cargo.lock` contains exact information about your dependencies. It is maintained by Cargo and should not be manually edited

**dependency by git**

```
[dependencies]
rand = { git = "https://github.com/rust-lang-nursery/rand.git", rev = "9f35b8e" }
```

#### Cargo Home

```bash
# vendoring all dependencies of a project
`cargo vendor`
```

#### build cache

https://doc.rust-lang.org/cargo/guide/build-cache.html

## Rust

https://doc.rust-lang.org/book/

