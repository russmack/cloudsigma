package cloudsigma

const (
	EndpointIps = "ips"
)

// IPs
type Ips struct {
	Args *Args
}
type IpRequest struct {
}

// NewIps returns a Ips object.
func NewIps() *Ips {
	o := Ips{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointIps
	return &o
}

// NewList returns the args required for a Ips GET request.
func (o *Ips) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// NewListDetailed returns the args required for a Ips GET request.
func (o *Ips) NewListDetailed() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/detail"
	return o.Args
}

// NewGet returns the args required for a Ips GET request.
func (o *Ips) NewGet(uuid string) *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/" + uuid + "/"
	return o.Args
}
