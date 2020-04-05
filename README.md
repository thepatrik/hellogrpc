# hellogrpc

[![Build Status](https://travis-ci.org/thepatrik/hellogrpc.svg?branch=master)](https://travis-ci.org/thepatrik/hellogrpc) [![Go Report Card](https://goreportcard.com/badge/github.com/thepatrik/hellogrpc)](https://goreportcard.com/report/github.com/thepatrik/hellogrpc) [![GoDoc](https://godoc.org/github.com/thepatrik/hellogrpc?status.svg)](https://godoc.org/github.com/thepatrik/hellogrpc)

Command line utility to demonstrate a simple gRPC server and client setup in golang. The gRPC server implements a function to mirror (reverse) a text, as well as a basic [health check](https://github.com/grpc/grpc-go/blob/master/health/grpc_health_v1/health.pb.go).

### Setup

Install with (requires [golang](https://golang.org/doc/install)):

```console
$ go install github.com/thepatrik/hellogrpc
```

### Usage

View usage help:

```console
$ hellogrpc --help
hellogrpc is a command line interface for testing gRPC in golang

Usage:
  hellogrpc [command]

Available Commands:
  help        Help about any command
  mirror      Mirror text
  serve       Run gRPC server

Flags:
  -h, --help   help for hellogrpc

Use "hellogrpc [command] --help" for more information about a command.
```

Run a gRPC server (defaults to port 9090):

```console
$ hellogrpc serve
```

Run a gRPC client in another process:

```console
$ hellogrpc mirror
input:              The quick brown 狐 jumped over the lazy 犬
output:             犬 yzal eht revo depmuj 狐 nworb kciuq ehT
```
