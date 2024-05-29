# Story Tracker

<!--toc:start-->

- [Story Tracker](#story-tracker)
  - [Getting Started](#getting-started)
    - [Install](#install)
    - [Environment](#environment)
    - [Dev Containers](#dev-containers)
  - [Building](#building)
  - [Testing](#testing)
  - [Migrations](#migrations)
  - [References](#references)

<!--toc:end-->

Sandbox application for learning the Go programming language.

## Getting Started

This section provides some guides on how to stand up your development environment.

### Install

This project is written in Go and requires that it be installed prior to development. There are a few ways to install
Go on your system. The easiest way is to download the installer from the
[official Go website](https://go.dev/doc/install). The recommended method is to use a tool like
[mise-en-place](https://mise.jdx.dev/) to manage your Go installations. Other methods include using a package manager
like [brew](https://brew.sh/) or [apt](https://ubuntu.com/server/docs/package-management).

In addition to Go, this project also requires the following tools:

- [sqlc](https://sqlc.dev/): Generate type-safe Go code from SQL
- [dbmate](https://github.com/amacneil/dbmate): Database migration tool
- [just](https://github.com/casey/just): Tool for running project specific commands
- [air](https://github.com/air-verse/air): Live reload tool for Go applications

### Environment

This project uses `.env` files to manage the configuration of the current environment. The `.env.template` file should
be copied and renamed to `.env.local` and updated with the appropriate values.

### Dev Containers

This project provides configurations to be able to stand up your development environment inside a Docker container.
This requires a running instance of docker and a tool that supports the development container
[specification](https://containers.dev/supporting). Some examples are [DevPod](https://devpod.sh/) and VS Code via the
[Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)
extension pack.

## Building

To build the application, run the following command:

```bash
just build
```

This will generate all go code from SQL and compile the application. The resulting binary will be placed in the `bin`
folder.

To run the application in watch mode, use the following command:

```bash
just watch
```

This will start the application using the `air` tool, which will automatically reload the application when changes are
detected.

## Testing

> [!NOTE]
> Unit tests are still a work in progress for this project.

To run the tests for the application, use the following command:

```bash
just test
```

## Migrations

This application uses the `dbmate` tool to manage database migrations. To create a new migration, run the following
command:

```bash
dbmate new <migration_name>
```

To execute all pending migrations, run the following command:

```bash
just migrate-up
```

To Rollback the last migration, run the following command:

```bash
just migrate-down
```

## References

- [Huma](https://huma.rocks/)
- [Fiber](https://gofiber.io/)
- [sqlc](https://sqlc.dev/)
- [dbmate](https://github.com/amacneil/dbmate)
- [air](https://github.com/air-verse/air)