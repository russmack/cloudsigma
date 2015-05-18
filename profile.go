package cloudsigma

const (
	EndpointProfile = "profile"
)

// Profile
type Profile struct {
	Args *Args
}

// NewProfile returns a Profile object.
func NewProfile() *Profile {
	o := Profile{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointProfile
	return &o
}

// NewList returns the args required for a Profile GET request.
func (o *Profile) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
