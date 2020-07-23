# goflux
[![codebeat badge](https://codebeat.co/badges/4ba2c5c6-da53-40c9-b1d3-506e2b5d852b)](https://codebeat.co/projects/github-com-nerzal-goflux-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Nerzal/goflux)](https://goreportcard.com/report/github.com/Nerzal/goflux)
[![Go Doc](https://godoc.org/github.com/Nerzal/goflux?status.svg)](https://godoc.org/github.com/Nerzal/goflux)
[![Build Status](https://github.com/Nerzal/goflux/workflows/Tests/badge.svg)](https://github.com/Nerzal/goflux/actions?query=branch%3Amaster+event%3Apush)
[![GitHub release](https://img.shields.io/github/tag/Nerzal/goflux.svg)](https://GitHub.com/Nerzal/goflux/releases/)
[![codecov](https://codecov.io/gh/Nerzal/goflux/branch/master/graph/badge.svg)](https://codecov.io/gh/Nerzal/goflux)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FNerzal%2Fgoflux.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FNerzal%2Fgoflux?ref=badge_shield)


## Note
Nothing is final right now and the API aswell as the models might change until V1.0 release

The created kubernetes files are currently narrowed down to my personal usecase. I'm very open for Changerequests to make this cli useful for a broader audience. 

## Usage Example

```sh
NAME:
   goflux - Used to automatically generate flux files

USAGE:
   goflux [global options] command [command options] [arguments...]

COMMANDS:
   init       Initialize new project
   namespace  Create a namespace file
   service    Create a service file
   backend    Create files for a backend service
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

To initialize a new project do
> goflux init --component myService

This will generate a basic folder structure.

To create all files needed for a backend deployment do 
> goflux backend --component myService --namespace myNamespace

### The following commands work without goflux init

**Create a Service File**
> goflux service --component myService --namespace myNamespace


**Create a Namespace File**
> goflux namespace --namespace myNamespace
