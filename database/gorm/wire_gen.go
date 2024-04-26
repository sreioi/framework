package gorm

import (
	"context"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/database/db"
)

func InitializeGorm(config2 config.Config, connection string) *GormImpl {
	configImpl := db.NewConfigImpl(config2, connection)
	dialectorImpl := NewDialectorImpl(config2, connection)
	gormImpl := NewGormImpl(config2, connection, configImpl, dialectorImpl)
	return gormImpl
}

func InitializeQuery(ctx context.Context, config2 config.Config, connection string) (*QueryImpl, error) {
	configImpl := db.NewConfigImpl(config2, connection)
	dialectorImpl := NewDialectorImpl(config2, connection)
	gormImpl := NewGormImpl(config2, connection, configImpl, dialectorImpl)
	queryImpl, err := BuildQueryImpl(ctx, config2, connection, gormImpl)
	if err != nil {
		return nil, err
	}
	return queryImpl, nil
}
