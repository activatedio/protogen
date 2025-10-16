package protogen

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFile_Write(t *testing.T) {

	a := assert.New(t)
	r := require.New(t)

	tests := []struct {
		name    string
		arrange func() File
		assert  func(got []byte, err error)
	}{
		{
			name: "simple",
			arrange: func() File {
				return NewFile("unit")
			},
			assert: func(got []byte, err error) {
				r.NoError(err)
				a.Equal("package unit;\n", string(got))
			},
		},
		{
			name: "full",
			arrange: func() File {

				f := NewFile("unit")

				f.AddMessages(
					NewMessage("Message1").AddFields(
						NewField("Field1", FieldParams{
							FieldType: "bool",
							Number:    1001,
							Repeated:  false,
						}),
						NewField("Field2", FieldParams{
							FieldType: "string",
							Number:    1002,
							Repeated:  true,
						}),
					),
					NewMessage("Message2").AddFields(
						NewField("Field3", FieldParams{
							FieldType: "number",
							Number:    1001,
							Repeated:  false,
						}),
						NewField("Field4", FieldParams{
							FieldType: "string",
							Number:    1002,
							Repeated:  false,
						}),
					),
				).AddServices(
					NewService("Service1").AddMethods(
						NewMethod("Method1"),
						NewMethod("Method2"),
					),
					NewService("Service2").AddMethods(
						NewMethod("Method3"),
						NewMethod("Method4"),
					),
				)

				return f
			},
			assert: func(got []byte, err error) {
				r.NoError(err)
				a.Equal(`package unit;

message Message1 {
  bool Field1 = 1001;
  repeated string Field2 = 1002;
}

message Message2 {
  number Field3 = 1001;
  string Field4 = 1002;
}

service Service1 {
  rpc Method1 () returns ();
  
  rpc Method2 () returns ();
  
}

service Service2 {
  rpc Method3 () returns ();
  
  rpc Method4 () returns ();
  
}

`, string(got))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := tt.arrange()
			buf := &bytes.Buffer{}
			err := f.Write(buf)
			tt.assert(buf.Bytes(), err)
		})
	}

}
