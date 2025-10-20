package tfl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/activatedio/protogen"
)

// fieldString is a constant used to represent a string field type in the context of field rendering operations.
const (
	fieldString = iota
)

// Field is an interface representing a renderable field with customizable end delimiters like semicolons or commas.
// Field extends protogen.Renderer, allowing rendering of structured output using the Render method.
// EndSemicolon returns the Field with a semicolon appended at the end.
// EndComma returns the Field with a comma appended at the end.
type Field interface {
	protogen.Renderer
	EndSemicolon() Field
	EndComma() Field
}

// field represents a structure for defining and rendering fields with specific types, values, and formatting rules.
type field struct {
	fieldType    int
	stringValue  string
	endSemicolon bool
	endComma     bool
	name         string
}

// EndSemicolon sets the field's end marker to a semicolon and returns the updated Field instance.
func (f *field) EndSemicolon() Field {
	f.endSemicolon = true
	return f
}

// EndComma ensures a comma is appended to the field when rendered, modifying the field's state.
func (f *field) EndComma() Field {
	f.endComma = true
	return f
}

// Render writes the field's name and value to the provided Output, formatted based on its type and end character settings.
// Returns an error if writing fails or the field type is unsupported.
func (f *field) Render(o protogen.Output) error {

	err := o.StartLine()
	if err != nil {
		return err
	}

	err = o.Write(fmt.Sprintf("%s: ", f.name))
	if err != nil {
		return err
	}

	sb := strings.Builder{}

	switch f.fieldType {
	case fieldString:
		sb.WriteString(fmt.Sprintf(`"%s"`, f.stringValue))
	default:
		return errors.New("unsupported field type")
	}

	switch {
	case f.endSemicolon:
		sb.WriteString(";")
	case f.endComma:
		sb.WriteString(",")
	}

	sb.WriteString("\n")

	return o.Write(sb.String())

}

// NewStringField creates a new field of type string with the specified name and value.
func NewStringField(name, value string) Field {
	return &field{
		fieldType:   fieldString,
		name:        name,
		stringValue: value,
	}
}
