package tfl_test

import (
	"bytes"
	"testing"

	"github.com/activatedio/protogen"
	"github.com/activatedio/protogen/tfl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessageValue_Render(t *testing.T) {

	a := assert.New(t)
	r := require.New(t)

	cases := []struct {
		name    string
		arrange func() tfl.MessageValue
		assert  func(got []byte, err error)
	}{
		{
			name:    "empty",
			arrange: tfl.NewMessageValue,
			assert: func(got []byte, err error) {
				r.NoError(err)
				a.Equal(`{
}`, string(got))
			},
		},
		{
			name: "full",
			arrange: func() tfl.MessageValue {
				unit := tfl.NewMessageValue()
				unit.AddFields(
					tfl.NewStringField("field1", "value1"),
					tfl.NewStringField("field2", "value2"),
					tfl.NewStringField("field3", "value3").EndSemicolon(),
					tfl.NewStringField("field4", "value4").EndComma(),
				)
				return unit
			},
			assert: func(got []byte, err error) {
				r.NoError(err)
				a.Equal(`{
  field1: "value1"
  field2: "value2"
  field3: "value3";
  field4: "value4",
}`, string(got))
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(_ *testing.T) {
			unit := tt.arrange()
			buf := &bytes.Buffer{}
			err := unit.Render(protogen.NewWriterOutput(buf))
			tt.assert(buf.Bytes(), err)
		})
	}
}
