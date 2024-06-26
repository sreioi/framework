package gorm

import (
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/database/orm"
	"gorm.io/gorm"
)

type Transaction struct {
	orm.Query
	instance *gorm.DB
}

func NewTransaction(tx *gorm.DB, config config.Config, connection string) *Transaction {
	return &Transaction{Query: NewQueryImpl(tx.Statement.Context, config, connection, tx, nil), instance: tx}
}

func (r *Transaction) Commit() error {
	return r.instance.Commit().Error
}

func (r *Transaction) Rollback() error {
	return r.instance.Rollback().Error
}
