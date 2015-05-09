package cloudsigma

const (
	EndpointLocations = "locations"
)

type Locations struct {
	Args Args
}

// NewCloudStatus returns a CloudStatus object.
func NewLocations() *Locations {
	o := Locations{}
	o.Args = Args{Resource: EndpointLocations}
	return &o
}

// NewGet returns the args required for a Locations GET request.
func (o *Locations) NewGet() Args {
	o.Args.Verb = "GET"
	//o.Args.RequiresAuth = true
	return o.Args
}
