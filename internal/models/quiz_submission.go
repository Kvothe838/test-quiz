package models

type QuizSubmission struct {
	ID         int
	Selections []ChoiceSelection
	HitsAmount int
}
