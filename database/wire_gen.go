package database

import (
	"context"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/database/db"
	"github.com/sreioi/framework/database/gorm"
)

func InitOrm(ctx context.Context, config config.Config, connection string) (*OrmImpl, error) {
	// 获取配置文件
	configImpl := db.NewConfigImpl(config, connection)
	// 初始化数据库
	dialectorImpl := gorm.NewDialectorImpl(config, connection)
	// 初始化gorm
	gormImpl := gorm.NewGormImpl(config, connection, configImpl, dialectorImpl)
	// 初始化query
	queryImpl, err := gorm.BuildQueryImpl(ctx, config, connection, gormImpl)
	if err != nil {
		return nil, err
	}
	// 初始化orm
	ormImpl, err := NewOrmImpl(ctx, config, connection, queryImpl)
	if err != nil {
		return nil, err
	}
	return ormImpl, nil
}
