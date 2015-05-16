package cloudsigma

const (
	EndpointAcls = "acls"
)

// Acls
type Acls struct {
	Args *Args
}

// NewAcls returns an Acls object.
func NewAcls() *Acls {
	o := Acls{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointAcls
	return &o
}

// List returns the args required for an Acls GET request.
func (o *Acls) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
