// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	core/defines/defines.go
	provides variables defination
*/

package defines

import (
	"encoding/json"
)

var (
	AesKey      string = "1234567812345678"
	Error_Param string
)

//Request Return Info
type RespInfo struct {
	ErrorNo   string
	ErrorInfo string
	Result    string
}

func init() {
	Error_Param = getErrorParam()
}

func getErrorParam() string {
	str, _ := json.Marshal(RespInfo{"1", "\u53c2\u6570\u9519\u8bef", "null"})
	return string(str)
}
