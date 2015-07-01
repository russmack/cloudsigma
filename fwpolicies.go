package cloudsigma

const (
	EndpointFwPolicies = "fwpolicies"
)

// FwPolicies
type FwPolicies struct {
	Args *Args
}

// NewFwPolicies returns a FwPolicies object.
func NewFwPolicies() *FwPolicies {
	o := FwPolicies{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointFwPolicies
	return &o
}

// NewList returns the args required for a FwPolicies GET request.
func (o *FwPolicies) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
