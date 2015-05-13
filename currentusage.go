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

// NewGet returns the args required for a CurrentUsage GET request.
func (o *CurrentUsage) NewGet() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
