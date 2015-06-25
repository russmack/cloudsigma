package cloudsigma

const (
	EndpointVlans = "vlans"
)

// Vlans
type Vlans struct {
	Args *Args
}
type VlanRequest struct {
}

// NewVlans returns a Vlans object.
func NewVlans() *Vlans {
	o := Vlans{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointVlans
	return &o
}

// NewList returns the args required for a Vlans GET request.
func (o *Vlans) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// NewListDetailed returns the args required for a Vlans GET request.
func (o *Vlans) NewListDetailed() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/detail"
	return o.Args
}

// NewGet returns the args required for a Vlans GET request.
func (o *Vlans) NewGet(uuid string) *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/" + uuid + "/"
	return o.Args
}
