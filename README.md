> ## Protogen
>
> This is a library to enable generation of proto files from go language inputs
>

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/activatedio/protogen/ci.yaml?branch=main&style=flat-square)](https://github.com/activatedio/protogen/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/activatedio/protogen?style=flat-square)](https://goreportcard.com/report/github.com/activatedio/protogen)

# Protogen

Simple project to generate proto files from go language inputs

## Usage

``` go

f := NewFile("unit")

f.AddMessages(
    NewMessage("Message1").AddFields(
        NewField("Field1", FieldParams{
            FieldType: "bool",
            Number:    1001,
            Repeated:  false,
        }),
        NewField("Field2", FieldParams{
            FieldType: "string",
            Number:    1002,
            Repeated:  true,
        }),
    ),
    NewMessage("Message2").AddFields(
        NewField("Field3", FieldParams{
            FieldType: "number",
            Number:    1001,
            Repeated:  false,
        }),
        NewField("Field4", FieldParams{
            FieldType: "string",
            Number:    1002,
            Repeated:  false,
        }),
    ),
).AddServices(
    NewService("Service1").AddMethods(
        NewMethod("Method1"),
        NewMethod("Method2"),
    ),
    NewService("Service2").AddMethods(
        NewMethod("Method3"),
        NewMethod("Method4"),
    ),
)

buf := &bytes.Buffer{}
err := f.Write(buf)


```
