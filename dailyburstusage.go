package cloudsigma

const (
	EndpointDailyBurstUsage = "dailyburstusage"
)

// DailyBurstUsage
type DailyBurstUsage struct {
	Args *Args
}

// NewDailyBurstUsage returns a DailyBurstUsage object.
func NewDailyBurstUsage() *DailyBurstUsage {
	o := DailyBurstUsage{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointDailyBurstUsage
	return &o
}

// NewList returns the args required for a DailyBurstUsage GET request.
func (o *DailyBurstUsage) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
