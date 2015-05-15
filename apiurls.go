package cloudsigma

const (
	EndpointApiUrls = ""
)

// ApiUrls
type ApiUrls struct {
	Args *Args
}

// NewApiUrls returns an ApiUrls object.
func NewApiUrls() *ApiUrls {
	o := ApiUrls{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointApiUrls
	return &o
}

// List returns the args required for an ApiUrls GET request.
func (o *ApiUrls) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = false
	return o.Args
}
