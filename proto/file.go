package proto

import (
	"fmt"
	"io"

	"github.com/activatedio/protogen"
)

// File defines an interface for managing and rendering a protocol buffer file.
type File interface {
	AddImports(i ...Import) File
	AddOptions(i ...Option) File
	AddMessages(m ...Message) File
	AddServices(s ...Service) File
	Write(w io.Writer) error
}

// file represents a container for a package, imports, messages, and services in a proto file.
type file struct {
	packageName string
	imports     []Import
	options     []Option
	messages    []Message
	services    []Service
}

// Write generates and writes the complete contents of the file, including package declaration, imports, messages, and services.
func (f *file) Write(w io.Writer) error {
	output := protogen.NewWriterOutput(w)

	if err := writeProtoHeader(output, f.packageName); err != nil {
		return err
	}

	if err := renderElements(output, toRenderers(f.imports)); err != nil {
		return err
	}

	if len(f.imports) > 0 {
		if err := output.WriteLines(""); err != nil {
			return err
		}
	}

	if err := renderElements(output, toRenderers(f.options)); err != nil {
		return err
	}

	if len(f.options) > 0 {
		if err := output.WriteLines(""); err != nil {
			return err
		}
	}

	if err := renderElements(output, toRenderers(f.messages)); err != nil {
		return err
	}

	if err := renderElements(output, toRenderers(f.services)); err != nil {
		return err
	}

	return nil
}

// writeProtoHeader writes the proto syntax and package declaration to the output.
func writeProtoHeader(output protogen.Output, packageName string) error {
	return output.WriteLines(
		"syntax = \"proto3\";",
		"",
		fmt.Sprintf("package %s;", packageName),
		"",
	)
}

// renderElements iterates through a slice of Renderer interfaces and renders each element.
// It returns an error if any of the Render calls fail.
func renderElements(output protogen.Output, elements []protogen.Renderer) error {
	for _, element := range elements {
		if err := element.Render(output); err != nil {
			return err
		}
	}
	return nil
}

// AddImports appends one or more Import instances to the file's imports and returns the updated File instance.
func (f *file) AddImports(i ...Import) File {
	is := map[string]bool{}
	for _, _i := range f.imports {
		is[_i.GetPath()] = true
	}
	for _, _i := range i {
		if !is[_i.GetPath()] {
			is[_i.GetPath()] = true
			f.imports = append(f.imports, _i)
		}
	}
	return f
}

// AddOptions appends one or more Option instances to the file's options list and returns the updated File instance.
func (f *file) AddOptions(i ...Option) File {
	f.options = append(f.options, i...)
	return f
}

// AddMessages appends one or more Message instances to the file's message list and returns the updated File.
func (f *file) AddMessages(m ...Message) File {
	f.messages = append(f.messages, m...)
	return f
}

// AddServices appends one or more Service objects to the file and returns the updated File instance.
func (f *file) AddServices(s ...Service) File {
	f.services = append(f.services, s...)
	return f
}

// NewFile creates a new File instance with the specified package name.
func NewFile(packageName string) File {
	return &file{
		packageName: packageName,
	}
}
