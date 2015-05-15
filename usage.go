package cloudsigma

const (
	EndpointUsage = "usage"
)

// Usage
type Usage struct {
	Args *Args
}

// NewUsage returns a Usage object.
func NewUsage() *Usage {
	o := Usage{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointUsage
	return &o
}

// List returns the args required for a Usage GET request.
func (o *Usage) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
