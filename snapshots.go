package cloudsigma

const (
	EndpointSnapshots = "snapshots"
)

// Snapshots
type Snapshots struct {
	Args *Args
}
type SnapshotRequest struct {
	Drive string `json:"drive"`
	Name  string `json:"name"`
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

// NewListDetailed returns the args required for a Snapshots GET request.
func (o *Snapshots) NewListDetailed() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/detail"
	return o.Args
}

// NewGet returns the args required for a Snapshots GET request.
func (o *Snapshots) NewGet(uuid string) *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/" + uuid + "/"
	return o.Args
}
