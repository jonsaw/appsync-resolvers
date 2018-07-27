package resolvers

import (
	"reflect"
)

// Repository stores all resolvers
type Repository struct {
	handler    Handler
	middleware []func(Handler) Handler
	resolvers  map[string]resolver
}

// Add stores a new resolver
func (r *Repository) Add(resolve string, handler interface{}) error {
	if r.resolvers == nil {
		r.resolvers = map[string]resolver{}
	}

	err := validators.run(reflect.TypeOf(handler))

	if err == nil {
		r.resolvers[resolve] = resolver{handler}
	}

	return err
}

// Handle responds to the AppSync request
func (r *Repository) Handle(in Invocation) (interface{}, error) {
	return r.handler.Serve(in)
}
