package proto

import (
	"fmt"
	"strings"

	"github.com/activatedio/protogen"
)

// FieldParams defines parameters for a field in a proto message, including its type, number, and whether it is repeated.
type FieldParams struct {
	FieldType     string
	Number        int32
	Repeated      bool
	InlineComment string
}

// Field represents an interface that extends Renderer for defining a structured field in a message or schema.
type Field interface {
	protogen.Renderer
}

// field represents a field with a name, type, unique number, and a flag indicating if it is repeated.
type field struct {
	name          string
	fieldType     string
	number        int32
	repeated      bool
	inlineComment string
}

// Render formats the field as a string in protocol buffer syntax and writes it to the provided Output instance.
func (f *field) Render(o protogen.Output) error {
	sb := strings.Builder{}
	if f.repeated {
		sb.WriteString("repeated ")
	}
	sb.WriteString(f.fieldType)
	sb.WriteString(" ")
	sb.WriteString(f.name)
	sb.WriteString(" = ")
	sb.WriteString(fmt.Sprintf("%d", f.number))
	sb.WriteString(";")
	if f.inlineComment != "" {
		sb.WriteString(" // ")
		sb.WriteString(f.inlineComment)
	}
	return o.WriteLines(sb.String())
}

// NewField creates a new Field with the specified name and parameters.
// name is the name of the field, and params defines its type, number, and repetition.
func NewField(name string, params FieldParams) Field {
	return &field{
		name:          name,
		fieldType:     params.FieldType,
		number:        params.Number,
		repeated:      params.Repeated,
		inlineComment: params.InlineComment,
	}
}
