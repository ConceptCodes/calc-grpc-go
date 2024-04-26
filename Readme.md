# Calculator GRPC API

[![Go Report Card](https://goreportcard.com/badge/github.com/conceptcodes/calc-grpc-go)](https://goreportcard.com/report/github.com/conceptcodes/calc-grpc-go)
[![Build Status](https://travis-ci.com/conceptcodes/calc-grpc-go.svg?branch=main)](https://travis-ci.com/conceptcodes/calc-grpc-go)
[![Coverage Status](https://coveralls.io/repos/github/conceptcodes/calc-grpc-go/badge.svg?branch=main)](https://coveralls.io/github/conceptcodes/calc-grpc-go?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/conceptcodes/calc-grpc-go.svg)](https://pkg.go.dev/github.com/conceptcodes/calc-grpc-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


This is a very simple calculator API written in Golang. Why [GRPC](https://grpc.io/docs/what-is-grpc/introduction/) you may ask?
It's a modern, high-performance, open-source and universal [RPC](https://en.wikipedia.org/wiki/Remote_procedure_call) framework.

**Available Operations**
- Add
- Subtract
- Multiply
- Divide

## Prerequisites

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/doc/install) (version 1.21 or higher)

The following dependencies are required to build and run this project

**Homebrew**
```sh
  brew install protobuf protoc-gen-go protoc-gen-go-grpc
```

or 

**Golang**
```sh
  go install google.golang.org/grpc/cmd/protoc-gen-go@v1.1
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## Usage

1. Generate go code from the .proto file
  ```sh
  make generate
  ```
2. Run the server
  ```sh
  make start
  ```
3. Run the client
  ```sh
  make client
  ```

