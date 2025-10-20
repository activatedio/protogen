package proto

import (
	"bytes"
	"testing"

	"github.com/activatedio/protogen"
	"github.com/activatedio/protogen/tfl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstant_Render(t *testing.T) {

	a := assert.New(t)
	r := require.New(t)

	cases := []struct {
		name     string
		unit     Constant
		expected string
	}{
		{
			name:     "string",
			unit:     NewStringConstant("test"),
			expected: `"test"`,
		},
		{
			name:     "float",
			unit:     NewFloatConstant(12.345),
			expected: `12.345`,
		},
		{
			name:     "int",
			unit:     NewIntConstant(12345),
			expected: `12345`,
		},
		{
			name:     "bool",
			unit:     NewBoolConstant(true),
			expected: `true`,
		},
		{
			name: "message",
			unit: NewMessageValueConstant(tfl.NewMessageValue().AddFields(
				tfl.NewStringField("f1", "v1"),
				tfl.NewStringField("f2", "v2"),
			)),
			expected: `{
  f1: "v1"
  f2: "v2"
}`,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(_ *testing.T) {
			buf := &bytes.Buffer{}
			err := tt.unit.Render(protogen.NewWriterOutput(buf))
			r.NoError(err)
			a.Equal(tt.expected, buf.String())
		})
	}
}
