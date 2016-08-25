/*	core/mgodao/MgoBaseDao.go
	provides mongodb initialize and related operations

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
