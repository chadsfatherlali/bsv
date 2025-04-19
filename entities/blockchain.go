package entities

type Blockchain struct {
	Blocks []Block `bson:"blocks" json:"blocks"`
}
