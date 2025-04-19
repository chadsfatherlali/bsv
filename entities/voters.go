package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Voter struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName    string             `json:"firstName"`
	LastName     string             `json:"lastName"`
	Document     int                `json:"document"`
	Picture      string             `json:"picture"`
	AlreadyVoted bool               `json:"alreadyVoted"`
	Timestamp    string             `json:"timestamp"`
}
