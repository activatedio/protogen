package proto

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
				a.Equal(`syntax = "proto3";

package unit;

`, string(got))
			},
		},
		{
			name: "full",
			arrange: func() File {

				f := NewFile("unit")

				f.AddImports(
					NewImport("subpath1/path1"),
					NewImport("subpath2/path2"),
				).
					AddOptions(
						NewOption("option1", NewStringConstant("value1")),
						NewOption("option2", NewStringConstant("value2")),
					).
					AddMessages(
						NewMessage("Message1").AddFields(
							NewField("Field1", FieldParams{
								FieldType:     "bool",
								Number:        1001,
								Repeated:      false,
								InlineComment: `@gotags: yaml:"field1"`,
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
						NewMethod("Method1", MethodParams{
							RequestName:  "Request1",
							ResponseName: "Response1",
						}),
						NewMethod("Method2", MethodParams{
							RequestName:  "Request2",
							ResponseName: "Response2",
						}),
					),
					NewService("Service2").AddMethods(
						NewMethod("Method3", MethodParams{
							RequestName:  "Request3",
							ResponseName: "Response3",
						}),
						NewMethod("Method4", MethodParams{
							RequestName:  "Request4",
							ResponseName: "Response4",
						}),
					),
				)

				return f
			},
			assert: func(got []byte, err error) {
				r.NoError(err)
				a.Equal(`syntax = "proto3";

package unit;

import "subpath1/path1";
import "subpath2/path2";

option option1 = "value1";
option option2 = "value2";

message Message1 {
  bool Field1 = 1001; // @gotags: yaml:"field1"
  repeated string Field2 = 1002;
}

message Message2 {
  number Field3 = 1001;
  string Field4 = 1002;
}

service Service1 {
  rpc Method1 (Request1) returns (Response1) {
  }
  rpc Method2 (Request2) returns (Response2) {
  }
}

service Service2 {
  rpc Method3 (Request3) returns (Response3) {
  }
  rpc Method4 (Request4) returns (Response4) {
  }
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
