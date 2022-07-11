package model

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"

	"blog/config"
	"blog/util/log"
)

var (
	mongoClient *mongo.Client
)

const (
	collectionInfoName    = "info"
	collectionPostName    = "post"
	collectionCommentName = "comment"
)

type dbTrait struct {
	db *mongo.Database
}

func init() {
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

	// Ping test
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Logger.Error(err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Logger.Error(err)
	}
	if err == nil {
		log.Logger.Println("MongoDB connected successfulliy!")
	}
}

func GetMongoGlobalClient() *mongo.Client {
	return mongoClient
}

// session 事务，但是需要mongo在cluster模式，慎用
func session(ctx context.Context, f func(ctx mongo.SessionContext) error) error {
	session, err := mongoClient.StartSession()
	if err != nil {
		return err
	}

	err = session.StartTransaction()
	if err != nil {
		return err
	}

	err = mongo.WithSession(ctx, session, f)

	if err != nil {
		err = session.AbortTransaction(ctx)
	} else {
		err = session.CommitTransaction(ctx)
	}

	session.EndSession(ctx)
	return err
}

func getDBTx(ctx context.Context) dbTrait {
	err := mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return dbTrait{
		db: mongoClient.Database(config.C.DB.DB),
	}
}

func (m *model) Close() {
	// DO NOTHING
}

func (m *model) Abort() {
	m.abort = true
}
