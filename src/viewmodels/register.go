package viewmodels

type RegisterRequest struct {
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}
