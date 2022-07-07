package model

import (
	"context"

	"blog-1.0/config"
	. "blog-1.0/util/log"
	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
)

var (
	pgDB *pg.DB
)

type dbTrait struct {
	tx *pg.Tx
}

func init() {
	options := &pg.Options{
		Addr:     config.C.DB.Addr,
		User:     config.C.DB.Username,
		Password: config.C.DB.Password,
		Database: config.C.DB.DB,
	}

	d := pg.Connect(options)

	err := d.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	if config.C.Debug {
		d.AddQueryHook(pgdebug.DebugHook{
			Verbose: true,
		})
	}

	pgDB = d

	Logger.Printf("PostgreSQL server connected")
}

func getDBTx(ctx context.Context) dbTrait {
	tx, err := pgDB.BeginContext(ctx)
	if err != nil {
		panic(err)
	}

	return dbTrait{
		tx: tx,
	}
}

func (m *model) Close() {
	if !m.abort {
		_ = m.tx.Commit()
	}
}

func (m *model) Abort() {
	_ = m.tx.Rollback()
	m.abort = true
}
