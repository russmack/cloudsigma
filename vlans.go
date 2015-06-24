package cloudsigma

const (
	EndpointVlans = "vlans"
)

// Vlans
type Vlans struct {
	Args *Args
}

// NewVlans returns a Vlans object.
func NewVlans() *Vlans {
	o := Vlans{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointVlans
	return &o
}

// NewVlansDetailed returns a Vlans object with a VlansDetailed endpoint.
func NewVlansDetailed() *Vlans {
	o := Vlans{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointVlans + "/detail"
	return &o
}

// NewList returns the args required for a Vlans GET request.
func (o *Vlans) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
