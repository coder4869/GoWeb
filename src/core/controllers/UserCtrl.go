// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/* core/controllers/UserCtrl.go

 */

package controllers

import (
	"core/mgodao"
	"core/models"
)

type UserCtrl struct {
}

/*
	Example for mongodb
@param
	addParam:	a map that contains register user information,
				keys for this map refer to defination of User
@retain
	string	: 	succeed - Insert User Id; failed - "false";
*/
func (this *UserCtrl) AddUser(addParam map[string]string) string {
	usr, _ := models.MgoUserWithMap(addParam)

	return mgodao.AddUser(usr)
}
