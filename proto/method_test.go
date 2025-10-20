package proto_test

import (
	"bytes"
	"testing"

	"github.com/activatedio/protogen"
	"github.com/activatedio/protogen/proto"
	"github.com/activatedio/protogen/tfl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMethod_Render(t *testing.T) {

	a := assert.New(t)
	r := require.New(t)

	cases := []struct {
		name     string
		arrange  func() proto.Method
		expected string
	}{
		{
			name: "simple",
			arrange: func() proto.Method {
				return proto.NewMethod("method1", proto.MethodParams{
					RequestName:  "request1",
					ResponseName: "response1",
				})
			},
			expected: "rpc method1 (request1) returns (response1) {\n}\n",
		},
		{
			name: "with options",
			arrange: func() proto.Method {
				return proto.NewMethod("method1", proto.MethodParams{
					RequestName:  "request1",
					ResponseName: "response1",
				}).AddOptions(
					proto.NewOption("option1", proto.NewStringConstant("value1")),
					proto.NewOption("option2",
						proto.NewMessageValueConstant(tfl.NewMessageValue().AddFields(
							tfl.NewStringField("field1", "value1").EndSemicolon(),
							tfl.NewStringField("field2", "value2").EndSemicolon(),
						)),
					),
				)
			},
			expected: `rpc method1 (request1) returns (response1) {
  option option1 = "value1";
  option option2 = {
    field1: "value1";
    field2: "value2";
  };
}
`,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(_ *testing.T) {
			buf := &bytes.Buffer{}
			unit := tt.arrange()
			err := unit.Render(protogen.NewWriterOutput(buf))
			r.NoError(err)
			a.Equal(tt.expected, buf.String())
		})
	}
}
