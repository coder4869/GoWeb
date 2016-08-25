/*	tests/UserService_test.go
provides testing related functions, now including testing cases
of following functions:
 	UserAddServe()

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

package test

import (
	"core/defines"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/coder4869/golibs/glcrypto"
	"github.com/coder4869/golibs/glio"
	"github.com/coder4869/golibs/glnet"
)

var reqIP string = "http://127.0.0.1:8080"

/*************************************************************************/
/********************* Test for UserAddServe WEB API *********************/
/*************************************************************************/

type NewUserParam struct {
	UserName string
	Password string
}

func newUserParam() string {
	paramObj := &NewUserParam{"zhangsan", "e10adc3949ba59abbe56e057f20f883e"}
	jsonParam, jsonErr := json.Marshal(paramObj)
	if jsonErr != nil {
		return "null"
	}

	//Base64 Decode --> AES Decrypt --> Json
	jsonStr, err := glcrypto.Base64AesEnResult(jsonParam, []byte(defines.AesKey))
	if err != nil { //param base64&Aes encode failed
		glio.FLPrintf("json AES Encrypt Error%v\n", err)
		return "null"
	}
	glio.FLPrintf("json Decrypt=%v\n", jsonStr)
	return jsonStr
}

func TestUserAddServe(t *testing.T) {
	reqURL := reqIP + "/goweb/user/add?param="
	reqURL += newUserParam()
	fmt.Println(reqURL)

	str, err := glnet.GETReqWithURL(reqURL)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(str)
	}
}

//Parallel Test
func Benchmark_UserAddServeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			reqURL := reqIP + "/goweb/user/add?param="
			reqURL += newUserParam()
			fmt.Println(reqURL)

			go glnet.GETReqWithURL(reqURL)
		}
	})
}

func Benchmark_UserAddServe(b *testing.B) {
	// must run b.N times. b.N will auto-adjust in running,
	// this ensure both time cost and caculated test data is reasonable.
	for i := 0; i < b.N; i++ {
		reqURL := reqIP + "/goweb/user/add?param="
		reqURL += newUserParam()

		fmt.Println(reqURL)
		go glnet.GETReqWithURL(reqURL)
	}
}
