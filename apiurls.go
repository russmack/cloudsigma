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

// NewList returns the args required for an ApiUrls GET request.
func (o *ApiUrls) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = false
	return o.Args
}
