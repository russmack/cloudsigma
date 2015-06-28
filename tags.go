package cloudsigma

const (
	EndpointTags = "acls"
)

// Tags
type Tags struct {
	Args *Args
}

// NewTags returns an Tags object.
func NewTags() *Tags {
	o := Tags{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointTags
	return &o
}

// NewList returns the args required for an Tags GET request.
func (o *Tags) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// NewGet returns the args required for a Tags GET request.
func (o *Tags) NewGet(uuid string) *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/" + uuid + "/"
	return o.Args
}
