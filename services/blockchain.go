package services

import (
	"blockchain_votation_system/config"
	"blockchain_votation_system/entities"
	"blockchain_votation_system/utils"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllBlocks(ctx context.Context) (*entities.Blockchain, error) {
	collection := config.DB.Collection("blocks")
	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)

	var blocks []entities.Block

	if err := res.All(ctx, &blocks); err != nil {
		return nil, err
	}

	blockchain := &entities.Blockchain{Blocks: blocks}

	return blockchain, nil
}

func GetLastBlock(ctx context.Context) (*entities.Block, error) {
	collection := config.DB.Collection("blocks")

	opts := options.FindOne().SetSort(bson.D{{Key: "index", Value: -1}})

	var block entities.Block

	err := collection.FindOne(ctx, bson.D{}, opts).Decode(&block)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &block, nil
}

func AddBlock(ctx context.Context, block *entities.Block) (*entities.Block, error) {
	prevBlock, err := GetLastBlock(ctx)

	if err != nil {
		return nil, err
	}

	block.Index = prevBlock.Index + 1
	block.Timestamp = time.Now().Format(time.RFC3339)
	block.PrevHash = prevBlock.Hash
	block.Hash = utils.CalculateHash(block)

	collection := config.DB.Collection("blocks")
	_, err = collection.InsertOne(ctx, block)

	if err != nil {
		return nil, err
	}

	return block, nil
}

func FindBlockByIndex(ctx context.Context, index int) (*entities.Block, error) {
	collection := config.DB.Collection("blocks")

	var block entities.Block

	filter := bson.M{"index": index}
	err := collection.FindOne(ctx, filter).Decode(&block)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &block, nil
}

func GenerateGenesisBlock(ctx context.Context) {
	block, err := FindBlockByIndex(ctx, 0)

	if err != nil {
		log.Println("Error:", err)
	} else if block == nil {
		genesisBlock := &entities.Block{
			Index:     0,
			Timestamp: time.Now().String(),
			Data:      "Genesis Block",
			PrevHash:  "",
		}

		genesisBlock.Hash = utils.CalculateHash(genesisBlock)
		collection := config.DB.Collection("blocks")
		block, err := collection.InsertOne(ctx, genesisBlock)

		if err != nil {
			log.Println("Error:", err)
		}

		log.Println("Genesis Block generated", block)
	}
}

func ValidateBlock(ctx context.Context, index int) (*entities.Validate, error) {
	prevIndex := index - 1
	block, err := FindBlockByIndex(ctx, index)
	prevBlock, prevErr := FindBlockByIndex(ctx, prevIndex)

	if err != nil {
		return &entities.Validate{Status: false}, err
	}

	if prevErr != nil {
		return &entities.Validate{Status: false}, prevErr
	}

	if block == nil || prevBlock == nil {
		return &entities.Validate{Status: false}, nil
	}

	res := utils.IsBlockValid(*block, *prevBlock)

	return &entities.Validate{Status: res}, nil
}
