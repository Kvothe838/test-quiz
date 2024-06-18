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
		hitsAmount, err := postQuizSubmission()
		if err != nil {
			missingQuestionsError, ok := err.(missingQuestionsErr)
			if ok {
				formattedQuestions := strings.Join(lo.Map(missingQuestionsError.questions, func(questionID int, _ int) string {
					return strconv.Itoa(questionID)
				}), ", ")

				fmt.Println("You're missing to select choice for questions ", formattedQuestions)
				return
			}

			fmt.Printf("An error occurred when submitting quiz: %v", err)
			return
		}

		fmt.Println("Choices posted, and your results are...")
		fmt.Println(hitsAmount, " correct answers!")
	},
}

func init() {
	rootCmd.AddCommand(submitQuizCmd)
}

func postQuizSubmission() (int, error) {
	url := fmt.Sprintf("%s/quiz-submission", baseUrl)
	resData, statusCode, err := backend.PostData(url, nil)
	if err != nil {
		return 0, errors.Wrap(err, "could not post choices confirmation")
	}

	if statusCode != http.StatusOK {
		if statusCode == http.StatusConflict {
			var res struct {
				MissingQuestions []int `json:"missing_questions"`
			}

			err = json.Unmarshal(resData, &res)
			if err != nil {
				return 0, errors.Wrap(err, "could not unmarshal quiz submission missing questions error")
			}

			return 0, missingQuestionsErr{questions: res.MissingQuestions}
		}

		if statusCode == http.StatusInternalServerError {
			return 0, errors.New("internal server error")
		}

		if statusCode == http.StatusNotFound {
			return 0, errors.New("page not found")
		}

		var res struct {
			Message string `json:"message"`
		}

		err = json.Unmarshal(resData, &res)
		if err != nil {
			return 0, errors.Wrapf(err, "could not unmarshal quiz submission error for str %s", string(resData))
		}

		return 0, errors.New(res.Message)

	}

	var res struct {
		HitsAmount int    `json:"hits_amount"`
		Message    string `json:"message"`
	}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return 0, errors.Wrap(err, "could not unmarshal choices confirmation data")
	}

	return res.HitsAmount, nil
}

type missingQuestionsErr struct {
	error
	questions []int
}
