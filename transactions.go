package cloudsigma

const (
	EndpointTransactions = "ledger"
)

// Transactions
type Transactions struct {
	Args *Args
}

// NewTransactions returns a Transactions object.
func NewTransactions() *Transactions {
	o := Transactions{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointTransactions
	return &o
}

// NewList returns the args required for a Transactions GET request.
func (o *Transactions) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
