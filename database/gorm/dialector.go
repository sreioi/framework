package gorm

import (
	"fmt"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/database"
	"github.com/sreioi/framework/contracts/database/orm"
	"github.com/sreioi/framework/database/db"

	"github.com/glebarez/sqlite"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DialectorSet = wire.NewSet(NewDialectorImpl, wire.Bind(new(Dialector), new(*DialectorImpl)))
var _ Dialector = &DialectorImpl{}

type Dialector interface {
	Make(configs []database.Config) ([]gorm.Dialector, error)
}

type DialectorImpl struct {
	config     config.Config
	connection string
	dsn        db.Dsn
}

func NewDialectorImpl(config config.Config, connection string) *DialectorImpl {
	return &DialectorImpl{
		config:     config,
		connection: connection,
		dsn:        db.NewDsnImpl(config, connection),
	}
}

func (d *DialectorImpl) Make(configs []database.Config) ([]gorm.Dialector, error) {
	driver := d.config.GetString(fmt.Sprintf("database.connections.%s.driver", d.connection))

	var dialectors []gorm.Dialector
	for _, item := range configs {
		var dialector gorm.Dialector
		var err error
		switch orm.Driver(driver) {
		case orm.DriverMysql:
			dialector = d.mysql(item)
		case orm.DriverPostgresql:
			dialector = d.postgresql(item)
		case orm.DriverSqlite:
			dialector = d.sqlite(item)
		case orm.DriverSqlserver:
			dialector = d.sqlserver(item)
		default:
			err = fmt.Errorf("err database driver: %s, only support mysql, postgresql, sqlite and sqlserver", driver)
		}

		if err != nil {
			return nil, err
		}

		dialectors = append(dialectors, dialector)
	}

	return dialectors, nil
}

func (d *DialectorImpl) mysql(config database.Config) gorm.Dialector {
	dsn := d.dsn.Mysql(config)
	if dsn == "" {
		return nil
	}

	return mysql.New(mysql.Config{
		DSN: dsn,
	})
}

func (d *DialectorImpl) postgresql(config database.Config) gorm.Dialector {
	dsn := d.dsn.Postgresql(config)
	if dsn == "" {
		return nil
	}

	return postgres.New(postgres.Config{
		DSN: dsn,
	})
}

func (d *DialectorImpl) sqlite(config database.Config) gorm.Dialector {
	dsn := d.dsn.Sqlite(config)
	if dsn == "" {
		return nil
	}

	return sqlite.Open(dsn)
}

func (d *DialectorImpl) sqlserver(config database.Config) gorm.Dialector {
	dsn := d.dsn.Sqlserver(config)
	if dsn == "" {
		return nil
	}

	return sqlserver.New(sqlserver.Config{
		DSN: dsn,
	})
}
