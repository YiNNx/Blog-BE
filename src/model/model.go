package model

import (
	"context"
	"time"

	"blog/config"
)

type Model struct {
	dbTrait
	ctx    context.Context
	cancel context.CancelFunc
	abort  bool
}

func GetModel() *Model {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.C.DB.Timeout))

	m := &Model{
		dbTrait: getDBTx(ctx),
		ctx:     ctx,
		cancel:  cancel,
		abort:   false,
	}

	return m
}

func (m *Model) Close() {
	m.cancel()
}

func (m *Model) Abort() {
	m.abort = true
	m.cancel()
}
