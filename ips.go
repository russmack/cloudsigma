package cloudsigma

const (
	EndpointIps = "ips"
)

// Servers
type Ips struct {
	Args *Args
}

// NewIps returns a Ips object.
func NewIps() *Ips {
	o := Ips{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointIps
	return &o
}

// List returns the args required for a Ips GET request.
func (o *Ips) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
