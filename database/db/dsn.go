package db

import (
	"fmt"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/contracts/database"
)

type Dsn interface {
	Mysql(config database.Config) string
	Postgresql(config database.Config) string
	Sqlite(config database.Config) string
	Sqlserver(config database.Config) string
}

type DsnImpl struct {
	config     config.Config
	connection string
}

func NewDsnImpl(config config.Config, connection string) *DsnImpl {
	return &DsnImpl{
		config:     config,
		connection: connection,
	}
}

func (d *DsnImpl) Mysql(config database.Config) string {
	host := config.Host
	if host == "" {
		return ""
	}

	charset := d.config.GetString("database.connections." + d.connection + ".charset")
	loc := d.config.GetString("database.connections." + d.connection + ".loc")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s&multiStatements=true",
		config.Username, config.Password, host, config.Port, config.Database, charset, true, loc)
}

func (d *DsnImpl) Postgresql(config database.Config) string {
	host := config.Host
	if host == "" {
		return ""
	}

	sslmode := d.config.GetString("database.connections." + d.connection + ".sslmode")
	timezone := d.config.GetString("database.connections." + d.connection + ".timezone")

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&timezone=%s",
		config.Username, config.Password, host, config.Port, config.Database, sslmode, timezone)
}

func (d *DsnImpl) Sqlite(config database.Config) string {
	return fmt.Sprintf("%s?multi_stmts=true", config.Database)
}

func (d *DsnImpl) Sqlserver(config database.Config) string {
	host := config.Host
	if host == "" {
		return ""
	}

	charset := d.config.GetString("database.connections." + d.connection + ".charset")

	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&charset=%s&MultipleActiveResultSets=true",
		config.Username, config.Password, host, config.Port, config.Database, charset)
}
