package cloudsigma

const (
	EndpointLocations = "locations"
)

type Locations struct {
	Args *Args
}

// NewCloudStatus returns a CloudStatus object.
func NewLocations() *Locations {
	o := Locations{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointLocations
	return &o
}

// NewList returns the args required for a Locations GET request.
func (o *Locations) NewList() *Args {
	o.Args.Verb = "GET"
	//o.Args.RequiresAuth = true
	return o.Args
}
