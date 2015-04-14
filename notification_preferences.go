package cloudsigma

const (
	EndpointNotificationPreferences = "notification_preferences"
)

// NotificationPreferences
type NotificationPreferences struct {
	Args Args
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
	o.Args = Args{Resource: EndpointNotificationPreferences}
	return &o
}

// NewGet returns the args required to get all notification preferences.
func (o *NotificationPreferences) NewGet() Args {
	o.Args.Verb = "GET"
	o.Args.RequiresAuth = true
	return o.Args
}

// NewSet returns the args required to update a specified notification preference.
func (o *NotificationPreferences) NewSet(pref Preference) Args {
	o.Args.Verb = "PUT"
	o.Args.RequiresAuth = true
	o.Args.Body = pref
	return o.Args
}
