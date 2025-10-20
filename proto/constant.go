package proto

import (
	"errors"
	"fmt"
	"strings"

	"github.com/activatedio/protogen"
	"github.com/activatedio/protogen/tfl"
)

// Constant represents an interface for rendering structured constants into an Output implementation.
type Constant interface {
	protogen.Renderer
}

// constString is the first constant in the iota sequence.
// constBool is the second constant in the iota sequence.
// constFloat is the third constant in the iota sequence.
// constInt is the fourth constant in the iota sequence.
// constMessage is the fifth constant in the iota sequence.
const (
	constString = iota
	constBool
	constFloat
	constInt
	constMessage
)

// constant represents a flexible type that encapsulates various constant values such as strings, booleans, floats, integers, or messages.
type constant struct {
	constType    int
	stringValue  string
	boolValue    bool
	floatValue   float64
	intValue     int
	messageValue tfl.MessageValue
}

// Render serializes the constant value to the provided Output based on its type and returns an error if any occurs.
func (c *constant) Render(o protogen.Output) error {

	switch c.constType {
	case constString:
		return o.Write(fmt.Sprintf(`"%s"`, c.stringValue))
	case constBool:
		return o.Write(fmt.Sprintf("%t", c.boolValue))
	case constFloat:
		return o.Write(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", c.floatValue), "0"), "."))
	case constInt:
		return o.Write(fmt.Sprintf("%d", c.intValue))
	case constMessage:
		return c.messageValue.Render(o)
	default:
		return errors.New("unknown constant type")
	}
}

// NewStringConstant creates a new Constant of type string with the specified value.
func NewStringConstant(value string) Constant {
	return &constant{
		constType:   constString,
		stringValue: value,
	}
}

// NewBoolConstant creates a new Constant representing a boolean value.
func NewBoolConstant(value bool) Constant {
	return &constant{
		constType: constBool,
		boolValue: value,
	}
}

// NewFloatConstant creates a new float constant with the specified value and returns it as a Constant interface.
func NewFloatConstant(value float64) Constant {

	return &constant{
		constType:  constFloat,
		floatValue: value,
	}
}

// NewIntConstant creates a new constant of type integer with the specified value.
func NewIntConstant(value int) Constant {

	return &constant{
		constType: constInt,
		intValue:  value,
	}
}

// NewMessageValueConstant creates a Constant of type message using the provided tfl.MessageValue.
func NewMessageValueConstant(value tfl.MessageValue) Constant {

	return &constant{
		constType:    constMessage,
		messageValue: value,
	}
}
