package protogen

import (
	"io"
	"strings"
)

// Output defines an interface for writing lines of text, typically used for structured rendering with error handling.
type Output interface {
	WriteLines(...string) error
}

// Renderer defines an interface for rendering structured output into an Output implementation.
type Renderer interface {
	Render(o Output) error
}

// indentingOutput is a struct that implements the Output interface, adding configurable indentation to each written line.
// It delegates the actual writing to an underlying Output implementation provided in the 'delegate' field.
// The 'level' field specifies the number of spaces used for indentation.
type indentingOutput struct {
	delegate Output
	level    int
}

// WriteLines writes the provided lines with added indentation based on the current level and delegates them for output.
// Returns an error if the underlying delegate fails to write any of the lines.
func (i *indentingOutput) WriteLines(s ...string) error {

	var err error

	for _, l := range s {

		sb := strings.Builder{}
		sb.WriteString(strings.Repeat(" ", i.level))
		sb.WriteString(l)

		err = i.delegate.WriteLines(sb.String())
		if err != nil {
			return err
		}
	}
	return nil
}

// NewIndentingOutput creates an Output instance that applies a consistent level of indentation to its written lines.
func NewIndentingOutput(delegate Output, level int) Output {
	return &indentingOutput{
		delegate: delegate,
		level:    level,
	}
}

// writerOutput is an implementation of the Output interface that writes data to an io.Writer.
type writerOutput struct {
	w io.Writer
}

// WriteLines writes each string in the provided slice to the writer, followed by a newline, and returns an error if any occur.
func (w *writerOutput) WriteLines(s ...string) error {
	var err error

	for _, l := range s {
		_, err = w.w.Write([]byte(l))
		if err != nil {
			return err
		}
		_, err = w.w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

// NewWriterOutput creates a new Output instance that writes lines to the specified io.Writer.
func NewWriterOutput(w io.Writer) Output {
	return &writerOutput{w: w}
}
