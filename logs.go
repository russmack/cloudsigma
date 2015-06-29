package cloudsigma

const (
	EndpointLogs = "logs"
)

// Logs
type Logs struct {
	Args *Args
}

// NewLogs returns a Logs object.
func NewLogs() *Logs {
	o := Logs{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointLogs
	return &o
}

// NewList returns the args required for a Logs GET request.
func (o *Logs) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
