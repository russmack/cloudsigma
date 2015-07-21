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
// TODO: This endpoint seems to return the same data as normal list.
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

// Create returns the args required to create a snapshot.
func (o *Snapshots) NewCreate(snapshot SnapshotRequest) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.AddHeaderPair("Content-Type", "application/json")
	body := struct{ SnapshotRequest }{snapshot}
	o.Args.Body = body
	return o.Args
}

// Delete returns the args required to delete a snapshot.
func (o *Snapshots) NewDelete(objectId string) *Args {
	o.Args.Verb = "DELETE"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	return o.Args
}
