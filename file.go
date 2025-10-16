package protogen

import (
	"fmt"
	"io"
)

// File defines an interface for managing and rendering a protocol buffer file.
type File interface {
	AddImports(i ...Import) File
	AddMessages(m ...Message) File
	AddServices(s ...Service) File
	Write(w io.Writer) error
}

// file represents a container for a package, imports, messages, and services in a proto file.
type file struct {
	packageName string
	imports     []Import
	messages    []Message
	services    []Service
}

// Write generates and writes the complete contents of the file, including package declaration, imports, messages, and services.
func (f *file) Write(w io.Writer) error {

	o := NewWriterOutput(w)
	var err error

	err = o.WriteLines(fmt.Sprintf("package %s;", f.packageName), "")

	if err != nil {
		return err
	}

	for _, i := range f.imports {
		err = i.Render(o)
		if err != nil {
			return err
		}
	}

	for _, m := range f.messages {
		err = m.Render(o)
		if err != nil {
			return err
		}
	}

	for _, s := range f.services {
		err = s.Render(o)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddImports appends one or more Import instances to the file's imports and returns the updated File instance.
func (f *file) AddImports(i ...Import) File {
	f.imports = append(f.imports, i...)
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
