package cloudsigma

const (
	EndpointBalance = "balance"
)

// Balance
type Balance struct {
	Args *Args
}

// NewBalance returns a Balance object.
func NewBalance() *Balance {
	o := Balance{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointBalance
	return &o
}

// List returns the args required for a Balance GET request.
func (o *Balance) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
