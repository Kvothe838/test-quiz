package cli

import (
	"encoding/json"
	"fmt"
	"github.com/Kvothe838/fast-track-test-quiz/cmd/cli/backend"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
	"strings"
)

var submitQuizCmd = &cobra.Command{
	Use:   "submit-quiz",
	Short: "Submit quiz with selected choices",
	Long:  `This command saves current choices selections in a new quiz submission.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := postQuizSubmission()
		if err != nil {
			missingQuestionsError, ok := err.(missingQuestionsErr)
			if ok {
				formattedQuestions := strings.Join(lo.Map(missingQuestionsError.questions, func(questionID int, _ int) string {
					return strconv.Itoa(questionID)
				}), ", ")

				fmt.Println("You're missing to select choice for questions ", formattedQuestions)
				return
			}

			if errors.Is(err, quizAlreadySubmittedErr) {
				fmt.Println("Quiz already submitted. You can't do more than one submission, sorry :(")
				return
			}

			fmt.Printf("An error occurred when submitting quiz: %v", err)
			return
		}

		fmt.Println("Choices posted, and your results are...")
		fmt.Println(res.HitsAmount, " correct answers!")

		if res.BetterThanPercentage == 0 {
			fmt.Println("You were not better than any quizzer, but keep on trying!")
			return
		}

		fmt.Println("You were better than ", res.BetterThanPercentage, "% of all quizzers.")
	},
}

func init() {
	rootCmd.AddCommand(submitQuizCmd)
}

func postQuizSubmission() (SubmitQuizRes, error) {
	url := fmt.Sprintf("%s/quiz-submission", baseUrl)
	resData, statusCode, err := backend.PostData(url, nil)
	if err != nil {
		return SubmitQuizRes{}, errors.Wrap(err, "could not post choices confirmation")
	}

	switch statusCode {
	case http.StatusOK:
		return handleSubmitQuizOK(resData)
	case http.StatusConflict:
		return SubmitQuizRes{}, handleSubmitQuizConflict(resData)
	case http.StatusInternalServerError:
		return SubmitQuizRes{}, errors.New("internal server error")
	case http.StatusNotFound:
		return SubmitQuizRes{}, errors.New("page not found")
	}

	var res struct {
		Message string `json:"message"`
	}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return SubmitQuizRes{}, errors.Wrapf(err, "could not unmarshal quiz submission error for str %s", string(resData))
	}

	return SubmitQuizRes{}, errors.New(res.Message)
}

type SubmitQuizRes struct {
	HitsAmount           int    `json:"hits_amount"`
	BetterThanPercentage int    `json:"better_than_percentage"`
	Message              string `json:"message"`
}

func handleSubmitQuizOK(resData []byte) (SubmitQuizRes, error) {
	var res SubmitQuizRes

	err := json.Unmarshal(resData, &res)
	if err != nil {
		return SubmitQuizRes{}, errors.Wrap(err, "could not unmarshal choices confirmation data")
	}

	return res, nil
}

func handleSubmitQuizConflict(resData []byte) error {
	var res struct {
		ErrCode string `json:"err_code,omitempty"`
	}

	err := json.Unmarshal(resData, &res)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal conflict error")
	}

	switch res.ErrCode {
	case "missing_questions":
		var res struct {
			MissingQuestions []int `json:"missing_questions"`
		}

		err := json.Unmarshal(resData, &res)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal quiz submission missing questions error")
		}

		return missingQuestionsErr{questions: res.MissingQuestions}
	case "quiz_already_submitted":
		return quizAlreadySubmittedErr
	}

	return nil
}

type missingQuestionsErr struct {
	error
	questions []int
}

var quizAlreadySubmittedErr = errors.New("quiz already submitted")
