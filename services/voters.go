package services

import (
	"blockchain_votation_system/config"
	"blockchain_votation_system/entities"
	"blockchain_votation_system/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddVoter(ctx context.Context, voter *entities.Voter) (*entities.Voter, error) {
	collection := config.DB.Collection("voters")

	err := collection.FindOne(ctx, bson.M{"document": voter.Document}).Decode(voter)

	if err == nil {
		return nil, &utils.VoterAlreadyExistsError{Document: voter.Document}
	}

	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	voter.Timestamp = time.Now().Format(time.RFC3339)
	voter.AlreadyVoted = false

	_, err = collection.InsertOne(ctx, voter)

	if err != nil {
		return nil, err
	}

	return voter, nil
}
