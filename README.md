# ☁️ Liam's wasmCloud Playground 🎡

## Introduction

This repo is for building Proof-of-Concept (PoC) [wasmCloud](https://wasmcloud.com/) applications. The goal is to deepen my understanding of wasmCloud, its internals, its strengths, its weaknesses, and how the [developer experience](https://en.wikipedia.org/wiki/User_experience#Developer_experience) (DX) is when building applications with it.

## Getting Started

The below steps will guide you through setting up the wasmCloud stack and deploying both a Rust and Go Actors & capability providers. Finally we will explore integrating both actors, and integrating an access control solution.

### Pre-requisites

1. Install the pre-requisites:
    - [just](https://github.com/casey/just)
    - Docker & Docker Compose
    - [wash](https://wasmcloud.com/docs/installation) (wasmCloud CLI)
    - Go & [TinyGo](https://tinygo.org/getting-started/install/) if wanting to compile a Go component.

### Running the wasmCloud Stack

1. Run `just up` to start the docker compose stack running to the wasmCloud components and dashboard UI, accessible at [localhost:3030](http://localhost:3030).

### Deploy a Rust App (Component and Provider)
1. `cd` into the `hello-rust` directory and run `wash build`. This builds the component in `../build` directory, which is mounted inside the wasmCloud container as specified in the Docker `compose.yaml` file.
1. Still in the `hello-rust` directory, run `just deploy` (which runs `wash app deploy wadm.yaml` under the hood) to deploy the hello application. This includes a http capability provider and an actor listening on [localhost:8061](http://localhost:8061).
1. Run `curl localhost:8061` to see the actor's response.

### Deploy a Go App (Component and Provider)

1. `cd` into the `hello-go` directory and run `wash build`. This builds the component in `../build` directory, which is mounted inside the wasmCloud container as specified in the Docker `compose.yaml` file.
1. Still in the `hello-go` directory, run `just deploy` (which runs `wash app deploy wadm.yaml` under the hood) to deploy the hello application. This includes a http capability provider and an actor listening on [localhost:8062](http://localhost:8062).
1. Run `curl localhost:8062` to see the actor's response.

### Troubleshooting

- You may need to [install Rust](https://www.rust-lang.org/learn/get-started) and the `wasi32-wasi` target, which can be installed by running `rustup target add wasm32-wasi` after installing Rust.

- You may need to [install TinyGo](https://tinygo.org/getting-started/install/).

### OpenFGA (Access Control Solution) Setup

By default (running `just up`), OpenFGA components will be provisioned. Use `just up-no-openfga` to run only wasmCloud components using Docker Compose.

1. The OpenFGA playground will be available at [localhost:3000/playground](http://localhost:3000/playground).
