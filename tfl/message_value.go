package tfl

import (
	"github.com/activatedio/protogen"
)

// MessageValue is an interface for rendering structured messages and managing fields within a message.
// It embeds protogen.Renderer and includes methods for adding fields to the message structure.
type MessageValue interface {
	protogen.Renderer
	AddFields(...Field) MessageValue
}

// messageValue is a struct used to build and render structured messages using a slice of Field implementations.
type messageValue struct {
	fields []Field
}

// AddFields appends the provided fields to the message and returns the updated MessageValue.
func (m *messageValue) AddFields(f ...Field) MessageValue {
	m.fields = append(m.fields, f...)
	return m
}

// Render writes the structured representation of messageValue to the provided Output, applying field rendering and indentation.
func (m *messageValue) Render(o protogen.Output) error {

	err := o.Write("{\n")

	if err != nil {
		return err
	}

	io := protogen.NewIndentingOutput(o, 2)

	for _, f := range m.fields {
		err = o.StartLine()
		if err != nil {
			return err
		}
		err = f.Render(io)
		if err != nil {
			return err
		}
	}

	err = o.StartLine()
	if err != nil {
		return err
	}

	return o.Write("}")
}

// NewMessageValue creates and returns a new instance of MessageValue with an initialized internal state.
func NewMessageValue() MessageValue {
	return &messageValue{}
}
