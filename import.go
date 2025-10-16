package protogen

// Import represents an import in the file
type Import interface {
	Renderer
	SetWeak()
	SetPublic()
	SetPath(string)
}
