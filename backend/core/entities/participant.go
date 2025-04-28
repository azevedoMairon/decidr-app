package entities

type Participant struct {
	ID           string `json:"id"            bson:"_id,omitempty"`
	Name         string `json:"name"          bson:"name"`
	IsNominated  bool   `json:"is_nominated"  bson:"isNominated"`
	IsEliminated bool   `json:"is_eliminated" bson:"isEliminated"`
	ImageUrl     string `json:"image_url"     bson:"imageUrl"`
}
