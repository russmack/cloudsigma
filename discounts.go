package cloudsigma

const (
	EndpointDiscounts = "discount"
)

// Discounts
type Discounts struct {
	Args *Args
}

// NewDiscounts returns a Discounts object.
func NewDiscounts() *Discounts {
	o := Discounts{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointDiscounts
	return &o
}

// List returns the args required for a Discounts GET request.
func (o *Discounts) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
