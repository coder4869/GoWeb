// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	tests/UserService_test.go
provides testing related functions, now including testing cases
of following functions:
 	UserAddServe()
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
		go glnet.GETReqWithURL(reqURL)
	}
}
