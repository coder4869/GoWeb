// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	core/serves/UserServe.go
	provides user serve related functions
*/

package serves

import (
	"core/controllers"
	"core/defines"
	"net/http"

	"github.com/coder4869/golibs/glio"
	"github.com/coder4869/golibs/glnet"
)

var (
	usrCtrl *controllers.UserCtrl
)

func UserAddServe(w http.ResponseWriter, r *http.Request) {
	addUsrParam := glnet.FormParse(w, r)

	err := Base64AESParamDecode(addUsrParam["param"], addUsrParam) //ServeHelper.go
	if err != nil {
		glio.FFLPrintf(w, defines.Error_Param) //output to client with w
		return
	}

	respJson := usrCtrl.AddUser(addUsrParam)
	glio.FFLPrintf(w, respJson) //output to client with w
}

func UserLoginServe(w http.ResponseWriter, r *http.Request) {
	glio.FFLPrintf(w, "UserLoginServe function is in developing\n") //output to client with w
}
