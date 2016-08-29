// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	core/serves/ServeHelper.go
	provides assistant functions for package serves
*/

package serves

import (
	"core/defines"
	"encoding/json"
	"strings"

	"github.com/coder4869/golibs/glcrypto"
	"github.com/coder4869/golibs/glio"
)

func Base64AESParamDecode(param string, paramMap map[string]string) error {

	param = strings.Replace(param, " ", "+", -1) //replace " " with "+"

	//Base64 Decode --> AES Decrypt --> Json
	jsonStr, err := glcrypto.Base64AesDeResult(param, []byte(defines.AesKey))
	if err != nil { //param base64&Aes decode failed
		glio.FLPrintf("json AES Decrypt Error%v\n", err)
		return err
	}
	glio.FLPrintf("json Decrypt=%v\n", jsonStr)

	err = json.Unmarshal([]byte(jsonStr), &paramMap)
	if err != nil {
		glio.FLPrintf("Json Param Error%v\n", err)
		return err
	}
	delete(paramMap, "param")
	glio.FLPrintf("json Unmarshal=%v\n", paramMap)

	return nil
}
