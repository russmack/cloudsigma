package cloudsigma

const (
	EndpointLicenses = "licenses"
)

// Licenses
type Licenses struct {
	Args *Args
}

// NewLicenses returns a Licenses object.
func NewLicenses() *Licenses {
	o := Licenses{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointLicenses
	return &o
}

// List returns the args required for a Licenses GET request.
func (o *Licenses) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
