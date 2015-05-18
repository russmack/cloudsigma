package cloudsigma

const (
	EndpointBurstUsage = "burstusage"
)

// BurstUsage
type BurstUsage struct {
	Args *Args
}

// NewBurstUsage returns a BurstUsage object.
func NewBurstUsage() *BurstUsage {
	o := BurstUsage{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointBurstUsage
	return &o
}

// NewList returns the args required for a BurstUsage GET request.
func (o *BurstUsage) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
