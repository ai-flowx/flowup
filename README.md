# installx

[![Build Status](https://github.com/ai-flowx/installx/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/ai-flowx/installx/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/ai-flowx/installx/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/ai-flowx/installx)
[![Go Report Card](https://goreportcard.com/badge/github.com/ai-flowx/installx)](https://goreportcard.com/report/github.com/ai-flowx/installx)
[![License](https://img.shields.io/github/license/ai-flowx/installx.svg)](https://github.com/ai-flowx/installx/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/ai-flowx/installx.svg)](https://github.com/ai-flowx/installx/tags)



## Introduction

*installx* is the installer of [flowx](https://github.com/ai-flowx/flowx) written in Go.



## Prerequisites

- Go >= 1.22.0



## Build

```bash
version=latest make build
```



## Usage

```
flowx installer

Usage:
  installx [flags]
  installx [command]

Available Commands:
  check       Check for updates to toolchains and installx
  help        Help about any command
  show        Show the installed toolchains
  update      Update toolchains and installx

Flags:
      --config string   config file (default "$HOME/.flowx/installx.yml")
  -h, --help            help for installx
  -v, --version         version for installx

Use "installx [command] --help" for more information about a command.
```



## Settings

*installx* parameters can be set in the directory [config](https://github.com/ai-flowx/installx/blob/main/config).

An example of configuration in [config.yml](https://github.com/ai-flowx/installx/blob/main/config/config.yml):

```yaml
apiVersion: v1
kind: installx
metadata:
  name: installx
spec:
  artifact:
    url: http://127.0.0.1:8080
    user: user
    pass: pass
```



## License

Project License can be found [here](LICENSE).



## Reference

- [bubbletea](https://github.com/charmbracelet/bubbletea/)
