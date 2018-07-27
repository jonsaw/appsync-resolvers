package resolvers

import "fmt"

type dispatch struct {
	repository *Repository
}

func (d dispatch) Serve(in Invocation) (interface{}, error) {
	handler, found := d.repository.resolvers[in.Resolve]

	if found {
		return handler.call(in.payload())
	}

	return nil, fmt.Errorf("No resolver found: %s", in.Resolve)
}
