/*	core/mysqldao/MysqlBaseDao.go
	provides Mysql database initialize and related operations

MIT License
Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb )

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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
