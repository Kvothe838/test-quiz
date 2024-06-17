package controller

import (
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/Kvothe838/fast-track-test-quiz/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
)

func (h *handler) getQuiz(ctx *gin.Context) {
	quiz, err := h.interactor.GetQuiz(ctx)
	if err != nil {
		logger.CtxErrorf(ctx, "error getting quiz: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	quizDTO := toGetQuizDTO(quiz)

	ctx.JSON(http.StatusOK, quizDTO)
}

func toGetQuizDTO(quiz models.Quiz) GetQuizDTO {
	return GetQuizDTO{
		Questions: toGetQuizQuestionsDTO(quiz.Questions),
	}
}

func toGetQuizQuestionsDTO(questions []models.Question) []GetQuizQuestionDTO {
	return lo.Map(questions, func(question models.Question, _ int) GetQuizQuestionDTO {
		return toGetQuizQuestionDTO(question)
	})
}

func toGetQuizQuestionDTO(question models.Question) GetQuizQuestionDTO {
	return GetQuizQuestionDTO{
		ID:          question.ID,
		Description: question.Description,
		Answers:     toGetQuizAnswersDTO(question.Answers),
	}
}

func toGetQuizAnswersDTO(answers []models.AnswerOption) []GetQuizAnswerDTO {
	return lo.Map(answers, func(answer models.AnswerOption, _ int) GetQuizAnswerDTO {
		return toGetQuizAnswerDTO(answer)
	})
}

func toGetQuizAnswerDTO(answer models.AnswerOption) GetQuizAnswerDTO {
	return GetQuizAnswerDTO{
		ID:          answer.ID,
		Description: answer.Description,
	}
}
