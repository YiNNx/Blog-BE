package model

import (
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"

	"blog/config"
	"blog/util/log"
)

var (
	mongoClient *mongo.Client
)

// GetCollectionName get the collectionName tag by reflect
func GetCollectionName(p interface{}) string {
	field := reflect.ValueOf(p).Elem().Type().Field(0)
	labelValue := field.Tag.Get("collection")
	return labelValue
}

type dbTrait struct {
	db *mongo.Database
}

func ConnectMongo() {
	var err error
	uri := fmt.Sprintf(
		"mongodb://%s:%s@mongo:%s",
		config.C.DB.Username,
		config.C.DB.Password,
		config.C.DB.Addr,
	)
	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Logger.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.C.DB.Timeout))
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Logger.Error(err)
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Logger.Error(err)
	}

	if err == nil {
		log.Logger.Info("MongoDB connected successfulliy!")
	}
}

func Disconnect() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		log.Logger.Error(err)
	}
}

func getDBTx(ctx context.Context) dbTrait {
	err := mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Logger.Error(err)
	}

	return dbTrait{
		db: mongoClient.Database(config.C.DB.DB),
	}
}

// mongoCreateDocument 添加document
func (m *Model) CreateDocument(p interface{}) (primitive.ObjectID, error) {
	coll := m.db.Collection(GetCollectionName(p))
	res, err := coll.InsertOne(m.ctx, p)
	return res.InsertedID.(primitive.ObjectID), err
}

func (m *Model) UpdateDocument(_id primitive.ObjectID, p interface{}) (int64, error) {
	coll := m.db.Collection(GetCollectionName(p))
	update, err := structToDoc(p)
	if err != nil {
		return 0, err
	}
	res, err := coll.UpdateOne(
		m.ctx,
		bson.D{{Key: "_id", Value: _id}},
		bson.D{{Key: "$set", Value: update}},
	)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func structToDoc(p interface{}) (bson.D, error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return nil, err
	}
	var doc bson.D
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// CAUTIOUS: p is a pointer
func (m *Model) GetDocument(p interface{}) ([]byte, error) {
	coll := m.db.Collection(GetCollectionName(p))
	filter, err := structToDoc(p)
	if err != nil {
		return nil, err
	}
	var res bson.D
	err = coll.FindOne(m.ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	doc, err := bson.Marshal(res)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (m *Model) GetAllDocuments(p interface{}) ([]bson.D, error) {
	coll := m.db.Collection(GetCollectionName(p))
	cursor, err := coll.Find(m.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var res []bson.D
	err = cursor.All(m.ctx, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// mongoDeleteDocument 根据field删除document
func (m *Model) DeleteDocument(p interface{}) (int64, error) {
	coll := m.db.Collection(GetCollectionName(p))
	filter, err := structToDoc(p)
	if err != nil {
		return 0, err
	}
	deleteResult, err := coll.DeleteOne(m.ctx, filter)
	if err != nil {
		return 0, err
	}
	return deleteResult.DeletedCount, nil
}
