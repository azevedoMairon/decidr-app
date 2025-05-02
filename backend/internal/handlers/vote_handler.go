package handlers

import (
	"net/http"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/azevedoMairon/decidr-app/internal/services"
	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	service services.VoteService
}

func NewVoteHandler(s services.VoteService) *VoteHandler {
	return &VoteHandler{service: s}
}

func (h *VoteHandler) PostVote(c *gin.Context) {
	ctx := c.Request.Context()

	var req entities.VoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	participants, err := h.service.PostVote(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, participants)
}
