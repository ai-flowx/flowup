# flowup

[![Build Status](https://github.com/ai-flowx/flowup/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/ai-flowx/flowup/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/ai-flowx/flowup/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/ai-flowx/flowup)
[![Go Report Card](https://goreportcard.com/badge/github.com/ai-flowx/flowup)](https://goreportcard.com/report/github.com/ai-flowx/flowup)
[![License](https://img.shields.io/github/license/ai-flowx/flowup.svg)](https://github.com/ai-flowx/flowup/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/ai-flowx/flowup.svg)](https://github.com/ai-flowx/flowup/tags)



## Introduction

*flowup* is the installer of [flowx](https://github.com/ai-flowx/flowx) written in Go.



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
  flowup [flags]
  flowup [command]

Available Commands:
  check       Check for updates to toolchains and flowup
  help        Help about any command
  show        Show the installed toolchains
  update      Update toolchains and flowup

Flags:
      --config string   config file (default "$HOME/.flowx/flowup.yml")
  -h, --help            help for flowup
  -v, --version         version for flowup

Use "flowup [command] --help" for more information about a command.
```



## Settings

*flowup* parameters can be set in the directory [config](https://github.com/ai-flowx/flowup/blob/main/config).

An example of configuration in [config.yml](https://github.com/ai-flowx/flowup/blob/main/config/config.yml):

```yaml
apiVersion: v1
kind: flowup
metadata:
  name: flowup
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
