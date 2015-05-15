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

// List returns the args required for a Subscriptions GET request.
func (o *Subscriptions) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
