package cloudsigma

const (
	EndpointCloudStatus = "cloud_status"
)

// CloudStatus
// {"free_tier": {"dssd": 53687091200, "mem": 1073741824}, "free_tier_monthly": {"tx": 5497558138880}, "guest": true, "signup": true, "sso": ["github", "twitter", "google", "facebook", "linkedin"], "trial": true}
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
