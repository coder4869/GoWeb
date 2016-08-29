// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/* conf/conf.go
setting config information
*/

package conf

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/coder4869/golibs/glfile"
	"github.com/coder4869/golibs/glio"
)

type DBConf struct {
	Type string // db type, now supports "mysql" or "mongodb"
	User string // user name
	Pwd  string // user password,
	Ip   string // db ip address, format "ip:port", like "127.0.0.1:10001"
	Name string // db name
}

var (
	DB            *DBConf
	DBRegisterURL string
)

/*
called in "core/mysqldao/MysqlBaseDao.go" or "core/mgodao/MgoBaseDao.go"
*/
func LoadDBJsonConf() {
	// if file "./conf/DBConf.json" not found while running,
	// please replace relative path as absolute path and try again
	//	err := glfile.FillObjWithJsonFile(&DB, "./conf/DBConf.json")
	_, conf_go, _, _ := runtime.Caller(0) //file info of current file
	dirPath := filepath.Dir(conf_go)      //dir of conf.go
	err := glfile.FillObjWithJsonFile(&DB, dirPath+"/DBConf.json")
	if err != nil {
		glio.FLPrintf("Load DB Json Config Error:%v\n", err)
	} else {
		if strings.EqualFold(DB.Type, "mysql") {
			DBRegisterURL = DB.User + ":" + DB.Pwd + "@tcp(" + DB.Ip + ")/" + DB.Name + "?charset=utf8"
		} else if strings.EqualFold(DB.Type, "mongodb") {
			DBRegisterURL = DB.Type + "://" + DB.User + ":" + DB.Pwd + "@" + DB.Ip + "/" + DB.Name
		}
	}
	glio.FLPrintf("DBRegisterURL=%v\n", DBRegisterURL)
}
