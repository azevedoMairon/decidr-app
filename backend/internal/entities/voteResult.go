package entities

import "time"

type VoteResult struct {
	ParticipantId string     `json:"participant_id" bson:"participantId"`
	Count         int        `json:"count"          bson:"count"`
	Hour          *time.Time `json:"hour,omitempty" bson:"hour,omitempty"`
}
