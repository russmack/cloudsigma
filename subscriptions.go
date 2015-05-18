package cloudsigma

const (
	EndpointSubscriptions = "subscriptions"
)

// Subscriptions
type Subscriptions struct {
	Args *Args
}

// NewSubscriptions returns a Subscriptions object.
func NewSubscriptions() *Subscriptions {
	o := Subscriptions{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointSubscriptions
	return &o
}

// NewList returns the args required for a Subscriptions GET request.
func (o *Subscriptions) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
