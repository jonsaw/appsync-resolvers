package resolvers

// Handler responds to requests in resolvers
type Handler interface {
	Serve(invocation) (interface{}, error)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions
type HandlerFunc func(invocation) (interface{}, error)

// Serve calls from resolver
func (f HandlerFunc) Serve(in invocation) (interface{}, error) {
	return f(in)
}
