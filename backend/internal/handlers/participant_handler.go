package handlers

import (
	"net/http"

	"github.com/azevedoMairon/decidr-app/internal/services"
	"github.com/gin-gonic/gin"
)

type ParticipantHandler struct {
	service services.ParticipantService
}

func NewHandler(s services.ParticipantService) *ParticipantHandler {
	return &ParticipantHandler{service: s}
}

func (h *ParticipantHandler) GetParticipants(c *gin.Context) {
	ctx := c.Request.Context()

	var isNominated *bool = nil
	if q := c.Query("isNominated"); q != "" {
		val := q == "true"
		isNominated = &val
	}

	participants, err := h.service.GetParticipants(ctx, isNominated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, participants)
}
