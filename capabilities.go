package cloudsigma

const (
	EndpointCapabilities = "capabilities"
)

type Capabilities struct {
	Args *Args
}

// NewCloudStatus returns a CloudStatus object.
func NewCapabilities() *Capabilities {
	o := Capabilities{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointCapabilities
	return &o
}

// List returns the args required for a Capabilities GET request.
func (o *Capabilities) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
