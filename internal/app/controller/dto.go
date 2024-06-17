package controller

type GetQuizDTO struct {
	Title       string
	Description string
	Questions   []GetQuizQuestionDTO `json:"questions"`
}

type GetQuizQuestionDTO struct {
	ID          int                `json:"id"`
	Description string             `json:"description"`
	Answers     []GetQuizAnswerDTO `json:"answers"`
}

type GetQuizAnswerDTO struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
