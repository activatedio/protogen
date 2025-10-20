package proto

import (
	"fmt"

	"github.com/activatedio/protogen"
)

// Option represents a renderable configuration element that extends the protogen.Renderer interface.
type Option interface {
	protogen.Renderer
}

// option represents an implementation of the Option interface for rendering an option and its associated value.
// It encapsulates a name and a constantValue that conforms to the Constant interface.
type option struct {
	name          string
	constantValue Constant
}

// Render generates the textual representation of an option and writes it to the provided Output instance.
// Returns an error if any stage of rendering or writing fails.
func (o *option) Render(out protogen.Output) error {
	err := out.StartLine()
	if err != nil {
		return err
	}
	err = out.Write(fmt.Sprintf("option %s = ", o.name))
	if err != nil {
		return err
	}
	err = o.constantValue.Render(out)
	if err != nil {
		return err
	}
	err = out.Write(";\n")
	if err != nil {
		return err
	}
	return nil
}

// NewOption creates a new Option with the specified name and associated Constant value.
func NewOption(name string, constantValue Constant) Option {
	return &option{
		name:          name,
		constantValue: constantValue,
	}
}
