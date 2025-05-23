package handlers

import (
	"log/slog"
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
		slog.Warn("[VoteHandler.PostVote] Failed to bind vote request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	result, err := h.service.PostVote(ctx, req)
	if err != nil {
		slog.Error("[VoteHandler.PostVote] Error processing vote", "error", err.Error(), "participant_id", req.ParticipantId)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *VoteHandler) GetResults(c *gin.Context) {
	ctx := c.Request.Context()

	var byHour *bool = nil
	if q := c.Query("byHour"); q != "" {
		val := q == "true"
		byHour = &val
		slog.Info("[VoteHandler.GetResults] Parsed byHour query parameter", "byHour", *byHour)
	}

	result, err := h.service.GetResults(ctx, byHour)
	if err != nil {
		slog.Error("[VoteHandler.GetResults] Failed to retrieve results", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
