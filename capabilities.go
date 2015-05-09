package cloudsigma

const (
	EndpointCapabilities = "capabilities"
)

type Capabilities struct {
	Args Args
}

// NewCloudStatus returns a CloudStatus object.
func NewCapabilities() *Capabilities {
	o := Capabilities{}
	o.Args = Args{Resource: EndpointCapabilities}
	return &o
}

// NewGet returns the args required for a Capabilities GET request.
func (o *Capabilities) NewGet() Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
