# GoML (Go Markup Language)
GoML is a YAML parser developed using the Go programming language. This parser can be used to 
parse YAML documents and utilize them within Go programs.

## Introduction
  YAML (YAML Ain't Markup Language) is a human-readable data serialization format that is commonly 
used for configuration files and data exchange between systems. GoML aims to provide a simple and 
efficient way to parse YAML documents in Go applications.

  GoML is designed to be easy to use, performant, and compatible with various YAML specifications. 
It allows developers to seamlessly integrate YAML parsing capabilities into their Go projects 
without having to deal with complex implementations.

## Installation

You can install GoML using the `go get` command:

```
go get github.com/unsubble/goml
```

After installing GoML, you can import it into your Go files:

```
import "github.com/unsubble/goml"
```

Then, you can use: 

```
  parser := parser.NewYAMLFileParser("FILE_PATH")
  parser.Parse()
```
To parse YAML documents within your Go program.

## License

GoML is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


