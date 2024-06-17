package models

type Quiz struct {
	Title       string
	Description string
	Questions   []Question
}
type Question struct {
	ID          int
	Description string
	Answers     []AnswerOption
}

type AnswerOption struct {
	ID          int
	Description string
	IsCorrect   bool
}
