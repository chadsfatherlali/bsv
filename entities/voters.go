package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Voter struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName    string             `json:"firstName" validate:"required,lettersonly"`
	LastName     string             `json:"lastName" validate:"required,lettersonly"`
	Document     int                `json:"document" validate:"required,numeric,max=9999999999"`
	Picture      string             `json:"picture" validate:"required,url"`
	AlreadyVoted bool               `json:"alreadyVoted"`
	Timestamp    string             `json:"timestamp"`
}
