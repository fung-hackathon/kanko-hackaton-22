package user

type User struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	StatusCommand int    `json:"statusCommand"`
	Progress      []bool `json:"progress"`
}
