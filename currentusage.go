package cloudsigma

const (
	EndpointCurrentUsage = "currentusage"
)

// CurrentUsage
type CurrentUsage struct {
	Args *Args
}

// NewCurrentUsage returns a CurrentUsage object.
func NewCurrentUsage() *CurrentUsage {
	o := CurrentUsage{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointCurrentUsage
	return &o
}

// NewList returns the args required for a CurrentUsage GET request.
func (o *CurrentUsage) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
