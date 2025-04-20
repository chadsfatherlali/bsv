package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Block struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Index     int                `json:"index"`
	Timestamp string             `json:"timestamp"`
	Data      string             `json:"data" validate:"required,uuid4"`
	PrevHash  string             `json:"prev_hash"`
	Hash      string             `json:"hash"`
}
