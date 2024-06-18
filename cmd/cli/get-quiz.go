package cli

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

var quizCmd = &cobra.Command{
	Use:   "get-quiz",
	Short: "Get a quiz",
	Long:  `This command fetches a quiz from quiz api`,
	Run: func(cmd *cobra.Command, args []string) {
		getQuiz()
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)
}

var baseUrl = "http://localhost:8080"

func getQuiz() {
	url := fmt.Sprintf("%s/quiz", baseUrl)
	responseBytes := getQuizData(url)
	quiz := quiz{}

	if err := json.Unmarshal(responseBytes, &quiz); err != nil {
		fmt.Printf("Could not unmarshal responseBytes. %v", err)
	}

	printQuiz(quiz)
}

func printQuiz(quiz quiz) {
	fmt.Println(quiz.Title)
	fmt.Println(quiz.Description)
	fmt.Println()
	for _, question := range quiz.Questions {
		fmt.Printf("Question %d\n", question.ID)
		fmt.Println(question.Description)
		for _, answerOption := range question.AnswerOptions {
			fmt.Printf("\t%d. %s\n", answerOption.ID, answerOption.Description)
		}

		fmt.Println()
	}
}

func getQuizData(url string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		log.Printf("Error on building quiz request: %v", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a quiz request: %v", err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body: %v", err)
	}

	return responseBytes
}

type quiz struct {
	Title       string
	Description string
	Questions   []question `json:"questions"`
}

type question struct {
	ID            int            `json:"id"`
	Description   string         `json:"description"`
	AnswerOptions []answerOption `json:"answers"`
}

type answerOption struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
