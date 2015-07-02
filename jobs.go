package cloudsigma

const (
	EndpointJobs = "jobs"
)

// Jobs
type Jobs struct {
	Args *Args
}

// NewJobs returns a Jobs object.
func NewJobs() *Jobs {
	o := Jobs{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointJobs
	return &o
}

// NewList returns the args required for a Jobs GET request.
func (o *Jobs) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
