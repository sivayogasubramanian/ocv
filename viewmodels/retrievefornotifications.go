package viewmodels

type RetrieveNotificationsRequest struct {
	Teacher      string `json:"teacher"`
	Notification string `json:"notification"`
}

type RetrieveNotificationsResponse struct {
	Recipients []string `json:"recipients"`
}
