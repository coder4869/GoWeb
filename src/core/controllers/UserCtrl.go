/*	core/controllers/UserCtrl.go

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
