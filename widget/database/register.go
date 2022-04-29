package database

import (
	"github.com/whf-sky/efficient/widget/database/mysql"
)

//----------------------RegisterModel--------------------------------

type modelDirvers map[string]modelDbnames

type modelDbnames map[string]modelKeys

type modelKeys map[string]ModelInterface

var models = modelDirvers{}

func RegisterModel(model ModelInterface) {
	parseModel(model)
	if _,ok := models[model.DriverName()]; !ok{
		models[model.DriverName()] = modelDbnames{}
	}
	if _,ok := models[model.DriverName()][model.Dbname()]; !ok{
		models[model.DriverName()][model.Dbname()] = modelKeys{}
	}
	models[model.DriverName()][model.Dbname()][model.Key()] = model
}

//----------------------RegisterSQLS--------------------------------

type SQLInterface interface {
	InsertQuery(table string, columns []string, valsLen int) (sql string, err error)
	UpdateQuery(table string, sets []string, where string) (sql string, err error)
	DeleteQuery(table string, where string) (sql string, err error)
}

var sqls = map[string]SQLInterface{}

func RegisterSQLS(name string, sql SQLInterface) {
	sqls[name] = sql
}

func SQL(name string) SQLInterface {
	return sqls[name]
}

//-----------------------init register-------------------------------------

func init() {
	RegisterSQLS("mysql", mysql.NewMysql())
}