package protogen

import "fmt"

// Service represents an interface that extends Renderer and allows adding RPC methods.
type Service interface {
	Renderer
	AddMethods(m ...Method) Service
}

// service is a struct that represents an RPC service with a name and a collection of methods.
type service struct {
	name    string
	methods []Method
}

// AddMethods appends one or more Method instances to the service and returns the updated Service instance.
func (s *service) AddMethods(m ...Method) Service {
	s.methods = append(s.methods, m...)
	return s
}

// Render generates a structured representation of the service and writes it to the given Output, returning any encountered error.
func (s *service) Render(o Output) error {

	var err error

	err = o.WriteLines(fmt.Sprintf("service %s {", s.name))

	if err != nil {
		return err
	}

	for _, m := range s.methods {
		i := NewIndentingOutput(o, 2)
		err = m.Render(i)
		if err != nil {
			return err
		}
	}

	return o.WriteLines("}", "")

}

// NewService creates a new Service instance with the specified name.
func NewService(name string) Service {
	return &service{
		name: name,
	}
}
