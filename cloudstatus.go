package cloudsigma

const (
	EndpointCloudStatus = "cloud_status"
)

// CloudStatus
type CloudStatus struct {
	Args Args
}

// NewCloudStatus returns a CloudStatus object.
func NewCloudStatus() *CloudStatus {
	o := CloudStatus{}
	o.Args = Args{Resource: EndpointCloudStatus}
	return &o
}

// NewGet returns the args required for a CloudStatus GET request.
func (o *CloudStatus) NewGet() Args {
	o.Args.Verb = "GET"
	//o.Args.RequiresAuth = true
	return o.Args
}
