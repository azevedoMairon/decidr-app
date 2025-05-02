package entities

type VoteRequest struct {
	ParticipantId string `json:"participant_id" binding:"required"`
}
