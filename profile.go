package cloudsigma

const (
	EndpointProfile = "profile"
)

// Profile
type Profile struct {
	Args Args
}

// NewProfile returns a Profile object.
func NewProfile() *Profile {
	o := Profile{}
	o.Args = Args{Resource: EndpointProfile}
	return &o
}

// List returns the args required for a Profile GET request.
func (o *Profile) List() Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
