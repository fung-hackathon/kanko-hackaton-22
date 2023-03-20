package user

type User struct {
	ID         string `json:"userid"`
	QuizStatus int    `json:"quizStatus"`
	Quizid     int    `json:"quizid"`
	Progress   []bool `json:"progress"`
}
