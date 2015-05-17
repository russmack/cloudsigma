package cloudsigma

const (
	EndpointServers = "servers"
)

// Servers
type Servers struct {
	Args *Args
}
type ServerRequest struct {
	Name        string `json:"name"`
	Cpu         int    `json:"cpu"`
	Memory      int    `json:"mem"`
	VncPassword string `json:"vnc_password"`
}

// NewServers returns a Servers object.
func NewServers() *Servers {
	o := Servers{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointServers
	return &o
}

// List returns the args required for a Servers GET request.
func (o *Servers) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// Create returns the args required to create a server.
func (o *Servers) Create(servers []ServerRequest) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.AddHeaderPair("Content-Type", "application/json")
	body := struct {
		Objects []ServerRequest `json:"objects"`
	}{servers}
	o.Args.Body = body
	return o.Args
}

// Edit returns the args required to update a server.
func (o *Servers) Edit(objectId string, server ServerRequest) *Args {
	o.Args.Verb = "PUT"
	o.Args.RequiresAuth = true
	o.Args.AddHeaderPair("Content-Type", "application/json")
	o.Args.ObjectId = objectId
	o.Args.Body = server
	return o.Args
}

// Delete returns the args required to delete a server.
func (o *Servers) Delete(objectId string) *Args {
	o.Args.Verb = "DELETE"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	return o.Args
}

// Start returns the args required to start a server.
func (o *Servers) Start(objectId string) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	o.Args.ActionName = "start"
	return o.Args
}

// Stop returns the args required to stop a server.
func (o *Servers) Stop(objectId string) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	o.Args.ActionName = "stop"
	return o.Args
}

// Shutdown returns the args required to shutdown a server.
func (o *Servers) Shutdown(objectId string) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	o.Args.ActionName = "shutdown"
	return o.Args
}
