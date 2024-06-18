package errors

import "github.com/pkg/errors"

var (
	ChoiceNotExistsErr         = errors.New("choice ID does not exist")
	QuestionNotExistsErr       = errors.New("question ID does not exist")
	SelectedOptionNotExistsErr = errors.New("selected option does not exist")
)
