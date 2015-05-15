package cloudsigma

const (
	EndpointNotificationPreferences = "notification_preferences"
)

// NotificationPreferences
type NotificationPreferences struct {
	Args *Args
}

// Preference is the object representing a notification preference.
type Preference struct {
	Contact string `json:"contact"`
	Medium  string `json:"medium"`
	Type    string `json:"type"`
	Value   bool   `json:"value"`
}

// NewNotificationPreferences returns a NotificationPreferences object with
// endpoint url set.
func NewNotificationPreferences() *NotificationPreferences {
	o := NotificationPreferences{}
	o.Args = NewArgs()
	o.Args.Resource = EndpointNotificationPreferences
	return &o
}

// List returns the args required to get all notification preferences.
func (o *NotificationPreferences) List() *Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// Edit returns the args required to update a specified notification preference.
func (o *NotificationPreferences) Edit(pref Preference) *Args {
	o.Args.Verb = "PUT"
	o.Args.RequiresAuth = true
	o.Args.Body = pref
	return o.Args
}
