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
func GetCollectionName(obj interface{}) string {
	field, _ := reflect.TypeOf(obj).FieldByName("collectionName")
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
func (m *model) CreateDocument(v interface{}) (primitive.ObjectID, error) {
	coll := m.db.Collection(GetCollectionName(v))
	res, err := coll.InsertOne(m.ctx, v)
	return res.InsertedID.(primitive.ObjectID), err
}

func structToDoc(v interface{}) (bson.D, error) {
	data, err := bson.Marshal(v)
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

// CAUTIOUS: v is a value not a pointer
func (m *model) GetDocument(v interface{}) ([]byte, error) {
	coll := m.db.Collection(GetCollectionName(v))
	filter, err := structToDoc(v)
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

// mongoDeleteDocument 根据field删除document（多个）
func (m *model) DeleteDocument(v interface{}) (int64, error) {
	coll := m.db.Collection(GetCollectionName(v))
	filter, err := structToDoc(v)
	if err != nil {
		return 0, err
	}
	deleteResult, err := coll.DeleteOne(m.ctx, filter)
	if err != nil {
		return 0, err
	}
	return deleteResult.DeletedCount, nil
}
