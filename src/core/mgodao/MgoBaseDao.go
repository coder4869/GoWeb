// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	core/mgodao/MgoBaseDao.go
	provides mongodb initialize and related operations
*/

package mgodao

import (
	"conf"
	"sync"

	"github.com/coder4869/golibs/glio"
	"gopkg.in/mgo.v2"
)

var (
	once       sync.Once
	mgoSession *mgo.Session
)

func init() {
	once.Do(conf.LoadDBJsonConf) //execute only once
}

//public: copy db session if exist
func GetDBSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(conf.DBRegisterURL)
		if err != nil {
			glio.FLPrintf("mongodb %v connect Error: %v\n", conf.DB.Name, err)
			//			panic(err)
			defer mgoSession.Close()
			return nil
		}
	}
	//	sessionClone := mgoSession.Clone() //default max connection pool: 4096
	sessionCopy := mgoSession.Copy() //default max connection pool: 4096
	glio.FLPrintf("mongodb %v connect succeed!\n", conf.DB.Name)
	sessionCopy.SetMode(mgo.Monotonic, true)
	return sessionCopy
}

//public
func GetDBCollection(collection string, s func(*mgo.Collection) error) error {
	session := GetDBSession()
	defer session.Close()
	c := session.DB(conf.DB.Name).C(collection)
	return s(c)
}
