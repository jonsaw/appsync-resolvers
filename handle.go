package resolvers

// Handler responds to requests in resolvers
type Handler interface {
	Serve(Invocation) (interface{}, error)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions
type HandlerFunc func(Invocation) (interface{}, error)

// Serve calls from resolver
func (f HandlerFunc) Serve(in Invocation) (interface{}, error) {
	return f(in)
}
