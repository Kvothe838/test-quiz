package controller

import (
	"fmt"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	internalerrors "github.com/Kvothe838/fast-track-test-quiz/internal/pkg/errors"
	"github.com/Kvothe838/fast-track-test-quiz/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (h *handler) selectChoice(ctx *gin.Context) {
	var data struct {
		QuestionID int `json:"question_id"`
		ChoiceID   int `json:"choice_id"`
	}

	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	selection := models.ChoiceSelection{
		QuestionID: data.QuestionID,
		ChoiceID:   data.ChoiceID,
	}

	err = h.interactor.SelectChoice(ctx, selection)
	if err != nil {
		if errors.Is(err, internalerrors.QuestionNotExistsErr) {
			msg := fmt.Sprintf("question does not exist, ID: %d", selection.QuestionID)
			logger.CtxInfo(ctx, msg)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": msg,
			})
			return
		}

		if errors.Is(err, internalerrors.ChoiceNotExistsErr) {
			msg := fmt.Sprintf("choice does not exist, questionID: %d, choiceID: %d", selection.QuestionID, selection.ChoiceID)
			logger.CtxInfo(ctx, msg)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": msg,
			})
			return
		}

		logger.CtxErrorf(ctx, "error selecting choice: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}
