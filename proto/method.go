package proto

import (
	"fmt"

	"github.com/activatedio/protogen"
)

// Method defines an interface for rendering structured RPC methods and supports adding options to the method configuration.
type Method interface {
	protogen.Renderer
	AddOptions(o ...Option) Method
}

// method represents a gRPC method definition with its name, request type, response type, and related options.
type method struct {
	name            string
	requestName     string
	responseName    string
	clientStreaming bool
	serverStreaming bool
	options         []Option
}

// AddOptions appends one or more Option instances to the method's options and returns the updated Method.
func (m *method) AddOptions(o ...Option) Method {
	m.options = append(m.options, o...)
	return m
}

// Render generates the RPC method definition with its name, request type, and response type in the provided output.
// It also processes and renders each associated option, handling errors from writing operations accordingly.
func (m *method) Render(o protogen.Output) error {

	req := m.requestName
	if m.clientStreaming {
		req = "stream " + req
	}
	resp := m.responseName
	if m.serverStreaming {
		resp = "stream " + resp
	}

	err := o.WriteLines(fmt.Sprintf("rpc %s (%s) returns (%s) {", m.name, req, resp))
	if err != nil {
		return err
	}

	io := protogen.NewIndentingOutput(o, 2)

	for _, opt := range m.options {
		err = opt.Render(io)
		if err != nil {
			return err
		}
	}

	return o.WriteLines("}")
}

// MethodParams defines the request and response names for a method in a proto service.
// ClientStreaming and ServerStreaming select streaming on either side; both
// true emits a bidirectional stream.
type MethodParams struct {
	RequestName     string
	ResponseName    string
	ClientStreaming bool
	ServerStreaming bool
}

// NewMethod creates a new Method instance with the provided name and MethodParams,
// which define the request and response types for the RPC method.
func NewMethod(name string, params MethodParams) Method {
	return &method{
		name:            name,
		requestName:     params.RequestName,
		responseName:    params.ResponseName,
		clientStreaming: params.ClientStreaming,
		serverStreaming: params.ServerStreaming,
	}
}
