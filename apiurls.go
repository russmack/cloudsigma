package cloudsigma

const (
	EndpointApiUrls = ""
)

// ApiUrls
type ApiUrls struct {
	Args *Args
}

// NewApiUrls returns a ApiUrls object.
func NewApiUrls() *ApiUrls {
	o := ApiUrls{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointApiUrls
	return &o
}

// NewGet returns the args required for a ApiUrls GET request.
func (o *ApiUrls) NewGet() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = false
	return o.Args
}
