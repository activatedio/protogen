package protogen

import "fmt"

// Method defines an interface that extends Renderer, representing an RPC method implementation.
type Method interface {
	Renderer
}

// method represents an implementation of an RPC method definition with its name, request, and response types.
type method struct {
	name         string
	requestName  string
	responseName string
}

// Render generates an RPC method definition string and writes it to the provided Output instance.
func (m *method) Render(o Output) error {

	var res []string

	res = append(res, fmt.Sprintf("rpc %s (%s) returns (%s);", m.name, m.requestName, m.responseName), "")

	return o.WriteLines(res...)
}

// NewMethod creates a new Method instance with the specified name.
func NewMethod(name string) Method {
	return &method{
		name: name,
	}
}
