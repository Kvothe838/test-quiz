package errors

import "github.com/pkg/errors"

var (
	ChoiceNotExistsErr         = errors.New("choice ID does not exist")
	QuestionNotExistsErr       = errors.New("question ID does not exist")
	SelectedChoiceNotExistsErr = errors.New("selected choice does not exist")
	QuizAlreadySubmittedErr    = errors.New("quiz already submitted")
)

type MissingChoicesSelectionErr struct {
	QuestionIDs []int
}

func (MissingChoicesSelectionErr) Error() string {
	return "some questions have missing choice selection"
}

func (m MissingChoicesSelectionErr) GetQuestionIDs() []int {
	return m.QuestionIDs
}

type MissingChoicesSelectionError interface {
	error
	GetQuestionIDs() []int
}
