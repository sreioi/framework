package facades

import "github.com/sreioi/framework/contracts/database/orm"

func Orm() orm.Orm {
	return App().MakeOrm()
}
