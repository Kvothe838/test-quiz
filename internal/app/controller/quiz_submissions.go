package controller

import (
	internalerrors "github.com/Kvothe838/fast-track-test-quiz/internal/pkg/errors"
	"github.com/Kvothe838/fast-track-test-quiz/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (h *handler) submitQuiz(ctx *gin.Context) {
	quiz, err := h.interactor.SubmitQuiz(ctx)
	if err != nil {
		missingChoicesSelectionErr, isMissingChoicesSelectionErr := errors.Cause(err).(internalerrors.MissingChoicesSelectionErr)
		if isMissingChoicesSelectionErr {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"err_code":          "missing_questions",
				"missing_questions": missingChoicesSelectionErr.GetQuestionIDs(),
			})
			return
		}

		if errors.Is(err, internalerrors.QuizAlreadySubmittedErr) {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"err_code": "quiz_already_submitted",
			})
			return
		}

		logger.CtxErrorf(ctx, "error submitting quiz: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	betterThanPercentage, err := h.interactor.CalcBetterThanPercentage(ctx, quiz.ID)
	if err != nil {
		logger.CtxErrorf(ctx, "error calculating better than percentage: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resultDTO := SubmitQuizResultDTO{
		HitsAmount:           quiz.HitsAmount,
		BetterThanPercentage: betterThanPercentage,
	}

	ctx.JSON(http.StatusOK, resultDTO)
}
