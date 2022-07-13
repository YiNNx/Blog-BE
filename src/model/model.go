package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog/config"
)

type Model interface {
	// 关闭数据库连接
	Close()
	// 终止操作，用于如事务的取消
	Abort()
	// TODO: 将Model层的实现列在这里，然后再去实现model结构体中的对应实现
	GetDocument(interface{}) ([]byte, error)
	CreateDocument(v interface{}) (primitive.ObjectID, error)
	DeleteDocument(v interface{}) (int64, error) 
}

type model struct {
	dbTrait
	ctx    context.Context
	cancel context.CancelFunc
	abort  bool
}

func GetModel() Model {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.C.DB.Timeout))

	m := &model{
		dbTrait: getDBTx(ctx),
		ctx:     ctx,
		cancel:  cancel,
		abort:   false,
	}

	return m
}

func (m *model) Close() {
	m.cancel()
}

func (m *model) Abort() {
	m.abort = true
	m.cancel()
}
