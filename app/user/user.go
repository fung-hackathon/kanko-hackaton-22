package user

type User struct {
	ID              string `json:"userid"`
	QuizStatus      int    `json:"quizStatus"`
	Quizid          int    `json:"quizid"`
	RecommendStatus int    `json:"recommendStatus"`
	Progress        []bool `json:"progress"`
}
