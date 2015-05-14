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

// NewGet returns the args required for a NotificationContacts GET request.
func (o *NotificationContacts) NewGet() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// NewAdd returns the args required to create a specified notification contact.
func (o *NotificationContacts) NewAdd(contact Contact) *Args {
	o.Args.Verb = "POST"
	o.Args.RequiresAuth = true
	o.Args.Body = contact
	return o.Args
}

// NewEdit returns the args required to update a specified notification contact.
func (o *NotificationContacts) NewEdit(objectId string, contact Contact) *Args {
	o.Args.Verb = "PUT"
	o.Args.RequiresAuth = true
	o.Args.Body = contact
	o.Args.ObjectId = objectId
	return o.Args
}

// NewDelete returns the args required to delete a specified notification contact.
func (o *NotificationContacts) NewDelete(objectId string) *Args {
	o.Args.Verb = "DELETE"
	o.Args.RequiresAuth = true
	o.Args.ObjectId = objectId
	return o.Args
}
