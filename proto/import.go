package proto

import (
	"fmt"

	"github.com/activatedio/protogen"
)

// Import represents an import in the file
type Import interface {
	protogen.Renderer
}

type importStatement struct {
	path string
}

func (i *importStatement) Render(o protogen.Output) error {
	return o.WriteLines(fmt.Sprintf("import \"%s\";", i.path))
}

// NewImport creates a new Import instance with the specified path.
func NewImport(path string) Import {
	return &importStatement{
		path: path,
	}
}
