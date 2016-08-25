/*	core/serves/ServeHelper.go
	provides assistant functions for package serves

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
