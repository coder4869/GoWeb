// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	core/models/MgoUserModel.go
	provides MgoUser(mongodb user) defination and related basic operations
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
