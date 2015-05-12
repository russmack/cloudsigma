package cloudsigma

const (
	EndpointNotificationContacts = "notification_contacts"
)

// NotificationContacts
type NotificationContacts struct {
	Args Args
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
	o.Args = Args{Resource: EndpointNotificationContacts}
	return &o
}

// NewGet returns the args required for a NotificationContacts GET request.
func (o *NotificationContacts) NewGet() Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}
