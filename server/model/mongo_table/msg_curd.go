package unrelation

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MsgDocModelInter interface {
	PushMsgsToDoc(docID string, msgsToMongo []MsgInfoModel) error
	Create(model *MsgDocModel) error
	UpdateMsg(docID string, index int64, key string, value any) (*mongo.UpdateResult, error)
}

type MsgMongoDriver struct {
	MsgCollection *mongo.Collection
}

func NewMsgMongoDriver(database *mongo.Database) MsgDocModelInter {
	collection := database.Collection(MsgDocModel{}.TableName())
	return &MsgMongoDriver{MsgCollection: collection}
}

func (m MsgMongoDriver) PushMsgsToDoc(docID string, msgsToMongo []MsgInfoModel) error {
	return m.MsgCollection.FindOneAndUpdate(
		context.Background(),
		bson.M{"doc_id": docID}, bson.M{"$push": bson.M{"msgs": bson.M{"$each": msgsToMongo}}},
	).Err()
}

func (m MsgMongoDriver) Create(model *MsgDocModel) error {
	_, err := m.MsgCollection.InsertOne(context.Background(), model)
	return err
}

func (m MsgMongoDriver) UpdateMsg(docID string, index int64, key string, value any) (*mongo.UpdateResult, error) {
	var field string
	if key == "" {
		field = fmt.Sprintf("msgs.%d", index)
	} else {
		field = fmt.Sprintf("msgs.%d.%s", index, key)
	}
	filter := bson.M{"doc_id": docID}
	update := bson.M{"$set": bson.M{field: value}}
	res, err := m.MsgCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return res, nil
}
