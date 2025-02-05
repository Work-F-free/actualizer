package scheduler

import (
	"actualizer/internal/common/errors"
	"actualizer/internal/common/responses"
	"actualizer/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service.Scheduler
}

func NewHandler(srvs service.Scheduler) *Handler {
	return &Handler{
		Scheduler: srvs,
	}
}

func (h *Handler) Start(c *gin.Context) {
	err := h.Scheduler.Start()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Error:   "can't start scheduler",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (h *Handler) Stop(c *gin.Context) {
	err := h.Scheduler.Stop()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Error:   "can't stop scheduler",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}
