package gosupervisor

// SupervisorRPC ...
type SupervisorRPC struct {
	URL string
}

// New ...
func New(url string) *SupervisorRPC {
	return &SupervisorRPC{URL: url}
}
