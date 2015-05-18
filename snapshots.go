package cloudsigma

const (
	EndpointSnapshots = "snapshots"
)

// Snapshots
type Snapshots struct {
	Args *Args
}

// NewSnapshots returns a Snapshots object.
func NewSnapshots() *Snapshots {
	o := Snapshots{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointSnapshots
	return &o
}

// NewList returns the args required for a Snapshots GET request.
func (o *Snapshots) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
