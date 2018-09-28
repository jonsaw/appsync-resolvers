package resolvers

import "encoding/json"

// ContextData data received from AppSync
type ContextData struct {
	Arguments json.RawMessage `json:"arguments"`
	Source    json.RawMessage `json:"source"`
}

// Invocation data received from AppSync
type Invocation struct {
	Resolve string      `json:"resolve"`
	Context ContextData `json:"context"`
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
