// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/GoWeb ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	WebMain.go
	entry of GoWeb
*/

package main

import (
	"core/serves"
	"log"
	"net/http"
)

func main() {
	//log setting
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("start WebServe8080 v2.0 beta")

	//router setting
	http.HandleFunc("/", serves.Index)
	http.HandleFunc("/goweb/user/add", serves.UserAddServe)
	http.HandleFunc("/goweb/user/login", serves.UserLoginServe)

	//listen port setting
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
