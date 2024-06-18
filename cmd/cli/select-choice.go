package cli

import (
	"fmt"
	"github.com/Kvothe838/fast-track-test-quiz/cmd/cli/backend"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
)

var selectChoiceCmd = &cobra.Command{
	Use:   "select-choice",
	Short: "Select a choice for a question",
	Long:  `This command selects a choice to a question of the quiz, overwriting last selected option.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		questionID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("First arg must be a number representing question ID")
			return
		}

		choiceID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Second arg must be a number representing selected choice ID")
			return
		}

		err = postChoiceSelection(questionID, choiceID)
		if err != nil {
			fmt.Println("An error occurred when selecting choice.")
			return
		}

		fmt.Printf("Successfully saved choice %d for question %d\n", choiceID, questionID)
	},
}

func init() {
	rootCmd.AddCommand(selectChoiceCmd)
}

func postChoiceSelection(questionID, choiceID int) error {
	url := fmt.Sprintf("%s/choice-selection", baseUrl)
	var data struct {
		QuestionID int `json:"question_id"`
		ChoiceID   int `json:"choice_id"`
	}

	data.QuestionID = questionID
	data.ChoiceID = choiceID
	resData, statusCode, err := backend.PostData(url, data)
	if err != nil {
		return errors.Wrap(err, "could not post data")
	}

	if statusCode != http.StatusCreated {
		return backend.BuildResponseError(resData)
	}

	return nil
}
