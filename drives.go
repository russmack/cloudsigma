package cloudsigma

const (
	EndpointDrives = "drives"
)

// Drives
type Drives struct {
	Args *Args
}

// NewDrives returns a Drives object.
func NewDrives() *Drives {
	o := Drives{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointDrives
	return &o
}

// NewGet returns the args required for a Drives GET request.
func (o *Drives) NewGet() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
