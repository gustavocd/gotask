# gotask ![](https://github.com/gustavocd/gotask/workflows/GoTask%20Tests/badge.svg)

Given a set of background tasks and a set of foreground tasks, a device is
**optimally configured** when the device is loaded with a background task and a
foreground task whose **resource consumption** is equal to or as close as possible
to the device's **resource capacity** without surpassing it.

## Quick Start

### Running dev environment

```bash
make dev
```

### Running tests

```bash
make test
```

### See more available commands

```bash
make help
```

## Build binary for production

This command will create a new binary called `gotask` inside `bin` folder, 
and another inside folder `linux_amd64` for linux servers.

```bash
make build
# You should see Version (git hash) and Build Time for the binary
./bin/gotask -version
```
