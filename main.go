package resolvers

// New returns a new Repository with a list of resolver
func New() *Repository {
	r := &Repository{}
	r.buildChain()
	return r
}
