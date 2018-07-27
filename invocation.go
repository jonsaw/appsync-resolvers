package resolvers

import "encoding/json"

type context struct {
	Arguments json.RawMessage `json:"arguments"`
	Source    json.RawMessage `json:"source"`
}

// Invocation data received from AppSync
type Invocation struct {
	Resolve string  `json:"resolve"`
	Context context `json:"context"`
}

func (in Invocation) isRoot() bool {
	return in.Context.Source == nil || string(in.Context.Source) == "null"
}

func (in Invocation) payload() json.RawMessage {
	if in.isRoot() {
		return in.Context.Arguments
	}

	return in.Context.Source
}
