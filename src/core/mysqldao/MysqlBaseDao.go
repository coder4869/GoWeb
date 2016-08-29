// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	core/mysqldao/MysqlBaseDao.go
	provides Mysql database initialize and related operations
*/

package mysqldao

import (
	"conf"
	_ "libs/mysql"
	"libs/orm"
	"sync"
)

var once sync.Once

type DbOrm struct {
}

func init() {
	once.Do(conf.LoadDBJsonConf) //execute only once

	//register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL) //orm.DRSqlite, orm.DRPostgres

	// set remote database
	orm.RegisterDataBase("default", conf.DB.Type, conf.DBRegisterURL, 30)

	// set local database
	//	orm.RegisterDataBase("default", "mysql", "db_username:db_password@/db_name?charset=utf8")

	//  create table
	//	orm.RunSyncdb("default", force, verbose) //force:force to create db or not(true/false); verbose:print the table creation process or not(true/false)
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 100)

	orm.Debug = true // open ORM debug mode
}

func (this *DbOrm) NewOrm() orm.Ormer {

	o := orm.NewOrm()
	o.Using(conf.DB.Name) //set database

	return o
}
