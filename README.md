# shup

[![Build Status](https://github.com/cligpt/shup/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/cligpt/shup/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/cligpt/shup/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/cligpt/shup)
[![Go Report Card](https://goreportcard.com/badge/github.com/cligpt/shup)](https://goreportcard.com/report/github.com/cligpt/shup)
[![License](https://img.shields.io/github/license/cligpt/shup.svg)](https://github.com/cligpt/shup/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/cligpt/shup.svg)](https://github.com/cligpt/shup/tags)



## Introduction

*shup* is the installer of [shai](https://github.com/cligpt/shai) written in Go.



## Prerequisites

- Go >= 1.22.0



## Build

```bash
version=latest make build
```



## Usage

```
shai installer

Usage:
  shup [flags]
  shup [command]

Available Commands:
  check       Check for updates to toolchains and shup
  help        Help about any command
  show        Show the active and installed toolchains
  update      Update toolchains and shup

Flags:
      --config string   config file (default "$HOME/.shai/shup.yml")
  -h, --help            help for shup
  -v, --version         version for shup

Use "shup [command] --help" for more information about a command.
```



## Settings

*shup* parameters can be set in the directory [config](https://github.com/cligpt/shup/blob/main/config).

An example of configuration in [config.yml](https://github.com/cligpt/shup/blob/main/config/config.yml):

```yaml
apiVersion: v1
kind: shup
metadata:
  name: shup
spec:
  drive:
    host: 127.0.0.1
    port: 65050
```



## License

Project License can be found [here](LICENSE).



## Reference

- [warp](https://www.warp.dev/)
