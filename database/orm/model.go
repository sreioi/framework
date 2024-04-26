package orm

import (
	"errors"
	"github.com/sreioi/framework/contracts/database/orm"
	"github.com/sreioi/framework/support/carbon"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const Associations = clause.Associations

var ErrRecordNotFound = errors.New("record not found")

var Observers = make([]Observer, 0)

type Observer struct {
	Model    any
	Observer orm.Observer
}

type Model struct {
	ID uint `gorm:"primaryKey"`
	Timestamps
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

type Timestamps struct {
	CreatedAt carbon.DateTime `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt carbon.DateTime `gorm:"autoUpdateTime;column:updated_at"`
}
