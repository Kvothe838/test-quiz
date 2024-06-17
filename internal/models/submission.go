package models

type Submission struct {
	ChosenAnswers []ChosenAnswer
}

type ChosenAnswer struct {
	QuestionID int
	AnswerID   int
}
