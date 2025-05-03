package entities

type VoteResult struct {
	ParticipantId string `json:"participant_id" bson:"_id,omitempty"`
	Count         int    `json:"count" bson:"count"`
}
