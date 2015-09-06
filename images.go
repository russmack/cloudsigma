package cloudsigma

const (
	EndpointImages = "drives"
)

// Images
type Images struct {
	Args *Args
}
type ImageRequest struct {
	Drive string `json:"drive"`
	Name  string `json:"name"`
}

// NewImages returns a Images object.
func NewImages() *Images {
	o := Images{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointImages
	return &o
}

// NewList returns the args required for a Images GET request.
/* TODO: ...
func (o *Images) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
*/

// NewListDetailed returns the args required for a Images GET request.
// TODO: This endpoint seems to return the same data as normal list.
/* TODO: ...
func (o *Images) NewListDetailed() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/detail"
	return o.Args
}
*/

// NewGet returns the args required for a Images GET request.
func (o *Images) NewDownload(uuid string) *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/" + uuid + "/download"
	return o.Args
}

// Create returns the args required to create a snapshot.
/* TODO: ...
func (o *Images) NewCreate(snapshot SnapshotRequest) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.AddHeaderPair("Content-Type", "application/json")
	body := struct{ SnapshotRequest }{snapshot}
	o.Args.Body = body
	return o.Args
}
*/
