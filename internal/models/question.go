package models

type Quiz struct {
	Title       string
	Description string
	Questions   []Question
}
type Question struct {
	ID          int
	Description string
	Choices     []Choice
}

type Choice struct {
	ID          int
	Description string
	IsCorrect   bool
}
