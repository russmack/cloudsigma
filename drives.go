package cloudsigma

const (
	EndpointDrives = "drives"
)

// Drives
type Drives struct {
	Args *Args
}
type DriveRequest struct {
	Media string `json:"media"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
}
type Drive struct {
	BootOrder  int    `json:"boot_order"`
	DevChannel string `json:"dev_channel"`
	Device     string `json:"device"`
	Uuid       string `json:"drive"`
}

// NewDrives returns a Drives object.
func NewDrives() *Drives {
	o := Drives{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointDrives
	return &o
}

// NewList returns the args required for a Drives GET request.
func (o *Drives) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// Create returns the args required to create a drive.
func (o *Drives) NewCreate(drives []DriveRequest) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.AddHeaderPair("Content-Type", "application/json")
	body := struct {
		Objects []DriveRequest `json:"objects"`
	}{drives}
	o.Args.Body = body
	return o.Args
}

// Delete returns the args required to delete a drive.
func (o *Drives) NewDelete(objectId string) *Args {
	o.Args.Verb = "DELETE"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	return o.Args
}

// NewListDetailed returns the args required for a Drives GET request.
func (o *Drives) NewListDetailed() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/detail"
	return o.Args
}

// NewGet returns the args required for a Drives GET request.
func (o *Drives) NewGet(uuid string) *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	o.Args.Resource += "/" + uuid + "/"
	return o.Args
}
