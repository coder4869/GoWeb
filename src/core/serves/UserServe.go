/*	core/serves/UserServe.go
	provides user serve related functions

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
