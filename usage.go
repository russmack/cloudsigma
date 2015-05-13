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

// NewGet returns the args required for a Usage GET request.
func (o *Usage) NewGet() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
