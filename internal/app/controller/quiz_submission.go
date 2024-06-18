package controller

import (
	internalerrors "github.com/Kvothe838/fast-track-test-quiz/internal/pkg/errors"
	"github.com/Kvothe838/fast-track-test-quiz/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (h *handler) submitQuiz(ctx *gin.Context) {
	hitsAmount, err := h.interactor.SubmitQuiz(ctx)
	if err != nil {
		missingChoicesSelectionErr, isMissingChoicesSelectionErr := errors.Cause(err).(internalerrors.MissingChoicesSelectionErr)
		if isMissingChoicesSelectionErr {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"missing_questions": missingChoicesSelectionErr.GetQuestionIDs(),
			})
			return
		}

		logger.CtxErrorf(ctx, "error submitting quiz: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resultDTO := SubmitQuizResultDTO{
		HitsAmount: hitsAmount,
	}

	ctx.JSON(http.StatusOK, resultDTO)
}
