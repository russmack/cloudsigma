package cloudsigma

const (
	EndpointNotificationContacts = "notification_contacts"
)

// NotificationContacts
type NotificationContacts struct {
	Args *Args
}

// Contact Request object, has writable fields.
type ContactRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Contact Response object, includes read-only fields.
type ContactResponse struct {
	Uuid        string `json:"uuid"`
	Main        bool   `json:"main"`
	ResourceUri string `json:"resource_uri"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
}

// NewNotificationContacts returns a NotificationContacts object.
func NewNotificationContacts() *NotificationContacts {
	o := NotificationContacts{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointNotificationContacts
	return &o
}

// NewList returns the args required for a NotificationContacts GET request.
func (o *NotificationContacts) NewList() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// Create returns the args required to create a specified notification contact.
func (o *NotificationContacts) NewCreate(contacts []ContactRequest) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.AddHeaderPair("Content-Type", "application/json")
	body := struct {
		Objects []ContactRequest `json:"objects"`
	}{contacts}
	o.Args.Body = body
	return o.Args
}

// Edit returns the args required to update a specified notification contact.
func (o *NotificationContacts) NewEdit(objectId string, contact ContactRequest) *Args {
	o.Args.Verb = "PUT"
	o.Args.RequiresAuth = true
	o.Args.AddHeaderPair("Content-Type", "application/json")
	o.Args.ObjectId = objectId
	o.Args.Body = contact
	return o.Args
}

// Delete returns the args required to delete a specified notification contact.
func (o *NotificationContacts) NewDelete(objectId string) *Args {
	o.Args.Verb = "DELETE"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	return o.Args
}
