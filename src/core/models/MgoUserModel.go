/*	core/models/MgoUserModel.go
	provides MgoUser(mongodb user) defination and related basic operations

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

package models

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Database User Collection Models
type MgoUser struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	UserName      string
	Password      string
	RegistTime    string
	UserType      string //Normal/VIP
	UserStatus    string //Online/Offline
	MobilePhone   string
	Email         string
	LastLoginTime string
	LastLoginIP   string
	Remark        string
	DelFlag       string //Yes/No
}

/* create MgoUser instance with usrMap
@param
	usrMap	: 	contains register user information, usrMap != nil,
				key refers to defination of MgoUser
@return
	MgoUser :	succeed - instance from information of usrMap; failed - nil.
	error	:	succeed - nil; failed - error info while creating instance
*/
func MgoUserWithMap(usrMap map[string]string) (*MgoUser, error) {
	usr := new(MgoUser)

	for key, val := range usrMap {
		if strings.EqualFold(key, "UserName") {
			usr.UserName = val
		} else if strings.EqualFold(key, "Password") {
			usr.Password = val
		}
	}

	return usr, nil
}

/*
now is used for AddUser operations
*/
func (this *MgoUser) MakeDefault() {
	this.RegistTime = time.Now().Format("2006-01-02 15:04:05") //fixd
	this.UserType = "Normal"                                   //Normal/VIP
	this.UserStatus = "Offline"                                //Online/Offline
	this.LastLoginTime = this.RegistTime
	this.DelFlag = "No" //Yes/No
}
