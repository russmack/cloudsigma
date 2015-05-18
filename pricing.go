package cloudsigma

const (
	EndpointPricing = "pricing"
)

// Pricing
type Pricing struct {
	Args *Args
}

// NewPricing returns a Pricing object.
func NewPricing() *Pricing {
	o := Pricing{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointPricing
	return &o
}

// NewList returns the args required for a Pricing GET request.
func (o *Pricing) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
