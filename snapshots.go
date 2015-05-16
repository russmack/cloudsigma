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

// List returns the args required for a Snapshots GET request.
func (o *Snapshots) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
