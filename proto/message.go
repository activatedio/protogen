package proto

import (
	"fmt"

	"github.com/activatedio/protogen"
)

// Message represents a composable and renderable structure that can hold and manage multiple Field elements.
// It extends the Renderer interface and provides functionality to add fields to a message.
type Message interface {
	protogen.Renderer
	AddFields(...Field) Message
}

// message represents a struct that defines a named message with a collection of structured fields.
type message struct {
	name   string
	fields []Field
}

// AddFields adds one or more Field elements to the message and returns the updated Message instance.
func (m *message) AddFields(f ...Field) Message {
	m.fields = append(m.fields, f...)
	return m
}

// Render generates a formatted representation of the message and writes it to the provided Output.
// It writes each field with proper indentation, utilizing the Output interface for structured rendering.
// Returns an error if any part of the rendering or writing process fails.
func (m *message) Render(o protogen.Output) error {

	var err error

	err = o.WriteLines(fmt.Sprintf("message %s {", m.name))

	if err != nil {
		return err
	}

	for _, f := range m.fields {

		io := protogen.NewIndentingOutput(o, 2)
		err = f.Render(io)
		if err != nil {
			return err
		}
	}

	return o.WriteLines("}", "")
}

// NewMessage creates a new Message instance with the specified name. It initializes the message with no fields.
func NewMessage(name string) Message {
	return &message{
		name: name,
	}
}
