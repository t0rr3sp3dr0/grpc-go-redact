# grpc-go-redact
## Why?

GRPC structs generated in Golang automaticly implement the `.String()` func printing out all of the fields. This does blocks writing custom overrides to redact secrets directly in the method ensuring secrets are never printed. 

This package runs after GRPC generation to automaticly overrite the `.String()` method with redaction support.

## Install

*  `go get github.com/samkreter/grpc-go-redact` or download the
  binaries from releases page.

## Usage

Example:

Run `grpc-go-redact` with generated file `test.pb.go`.

```
grpc-go-redact -input=./test.pb.go
```
The custom .String() method will updated directly in the `test.pb.go` file.
