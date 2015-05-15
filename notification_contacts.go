package cloudsigma

const (
	EndpointNotificationContacts = "notification_contacts"
)

// NotificationContacts
type NotificationContacts struct {
	Args *Args
}

// Contact object.
type Contact struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// NewNotificationContacts returns a NotificationContacts object.
func NewNotificationContacts() *NotificationContacts {
	o := NotificationContacts{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointNotificationContacts
	return &o
}

// List returns the args required for a NotificationContacts GET request.
func (o *NotificationContacts) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// Create returns the args required to create a specified notification contact.
func (o *NotificationContacts) Create(contact Contact) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.Body = contact
	return o.Args
}

// Edit returns the args required to update a specified notification contact.
func (o *NotificationContacts) Edit(objectId string, contact Contact) *Args {
	o.Args.Verb = "PUT"
	o.Args.RequiresAuth = true
	o.Args.Body = contact
	o.Args.ObjectId = objectId
	return o.Args
}

// Delete returns the args required to delete a specified notification contact.
func (o *NotificationContacts) Delete(objectId string) *Args {
	o.Args.Verb = "DELETE"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	return o.Args
}
