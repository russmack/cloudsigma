package cloudsigma

const (
	EndpointServers = "servers"
)

// Servers
type Servers struct {
	Args *Args
}

// NewServers returns a Servers object.
func NewServers() *Servers {
	o := Servers{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointServers
	return &o
}

// List returns the args required for a Servers GET request.
func (o *Servers) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
