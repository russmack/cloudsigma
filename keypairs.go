package cloudsigma

const (
	EndpointKeypairs = "keypairs"
)

// Keypairs
type Keypairs struct {
	Args *Args
}

// NewKeypairs returns a Keypairs object.
func NewKeypairs() *Keypairs {
	o := Keypairs{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointKeypairs
	return &o
}

// NewList returns the args required for a Keypairs GET request.
func (o *Keypairs) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
