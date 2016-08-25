/*	core/mgodao/MgoUserDao.go
	provides mongodb user related operations

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
	"core/models"

	"github.com/coder4869/golibs/glio"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* 	Insert
@param
	usr		: struct instance of MgoUser, used for insert.
@return
	string	: succeed - insert id; failed - "false".
*/
func AddUser(usr *models.MgoUser) string {

	usr.Id = bson.NewObjectId()
	usr.MakeDefault()
	insert := func(c *mgo.Collection) error {
		return c.Insert(usr)
	}

	err := GetDBCollection("MgoUser", insert)
	if err != nil {
		glio.FLPrintf("AddUser Error: %v!\n", err)
		return "false"
	}

	return string(usr.Id.Hex())
}

/*	Query
@param
	idValue	:	string type, key index for user query.
@return
	*models.MgoUser	: succeed - MgoUser object; failed - nil.
*/
func GetUserById(idValue string) *models.MgoUser {

	usr := new(models.MgoUser)
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"id": bson.ObjectIdHex(idValue)}).All(usr)
	}

	err := GetDBCollection("MgoUser", query)
	if err != nil {
		glio.FLPrintf("GetUserById Error: %v! \n", err)
		return nil
	}

	return usr
}

/*	Query
@return
	[]models.MgoUser	: succeed - MgoUser list; failed - nil.
*/
func GetAllUsersInfo() []models.MgoUser {

	var usrs []models.MgoUser
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&usrs)
	}

	err := GetDBCollection("MgoUser", query)
	if err != nil {
		glio.FLPrintf("GetAllUsersInfo Error: %v! \n", err)
		return nil
	}

	return usrs
}

/*	Update
@return
	bool	: succeed - true; failed - false.
*/
func UpdateUserInfo(query bson.M, change bson.M) bool {

	updat := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}

	err := GetDBCollection("MgoUser", updat)
	if err != nil {
		glio.FLPrintf("UpdateUserInfo Error: %v! \n", err)
		return false
	}

	return true
}

///**
// * 执行查询，此方法可拆分做为公共方法
// * [SearchPerson description]
// * @param {[type]} collectionName string [description]
// * @param {[type]} query          bson.M [description]
// * @param {[type]} sort           bson.M [description]
// * @param {[type]} fields         bson.M [description]
// * @param {[type]} skip           int    [description]
// * @param {[type]} limit          int)   (results      []interface{}, err error [description]
// */
//func SearchPerson(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
//	exop := func(c *mgo.Collection) error {
//		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
//	}
//	err = witchCollection(collectionName, exop)
//	return
//}
