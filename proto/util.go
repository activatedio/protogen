package proto

import "github.com/activatedio/protogen"

func toRenderers[T protogen.Renderer](ts []T) []protogen.Renderer {
	renderers := make([]protogen.Renderer, len(ts))
	for i, t := range ts {
		renderers[i] = t
	}
	return renderers
}
