package resolvers

// Use appends middleware to repository
func (r *Repository) Use(middleware func(Handler) Handler) {
	r.middleware = append(r.middleware, middleware)
	r.buildChain()
}

func (r *Repository) buildChain() {
	r.handler = dispatch{repository: r}
	for i := len(r.middleware) - 1; i >= 0; i-- {
		r.handler = r.middleware[i](r.handler)
	}
}
