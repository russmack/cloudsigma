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

// NewSet returns the args required to update a specified notification contact.
func (o *NotificationContacts) NewSet(contact Contact) *Args {
	o.Args.Verb = "PUT"
	o.Args.RequiresAuth = true
	o.Args.Body = contact
	return o.Args
}
