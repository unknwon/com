// Copyright 2013 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package com_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Unknwon/com"
)

// ------------------------------
// cmd.go
// ------------------------------

func ExampleColorLogS() {
	coloredLog := com.ColorLogS(fmt.Sprintf(
		"[WARN] This is a tesing log that should be colored, path( %s ),"+
			" highlight # %s #, error [ %s ].",
		"path to somewhere", "highlighted content", "tesing error"))
	fmt.Println(coloredLog)
}

func ExampleColorLog() {
	com.ColorLog(fmt.Sprintf(
		"[WARN] This is a tesing log that should be colored, path( %s ),"+
			" highlight # %s #, error [ %s ].",
		"path to somewhere", "highlighted content", "tesing error"))
}

// ------------- END ------------

// ------------------------------
// html.go
// ------------------------------

func ExampleHtml2JS() {
	htm := "<div id=\"button\" class=\"btn\">Click me</div>\n\r"
	js := string(com.Html2JS([]byte(htm)))
	fmt.Println(js)
	// Output: <div id=\"button\" class=\"btn\">Click me</div>\n
}

// ------------- END ------------

// ------------------------------
// path.go
// ------------------------------

func ExampleGetGOPATHs() {
	gps := com.GetGOPATHs()
	fmt.Println(gps)
}

func ExampleGetSrcPath() {
	srcPath, err := com.GetSrcPath("github.com/Unknwon/com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(srcPath)
}

// ------------- END ------------

// ------------------------------
// file.go
// ------------------------------

func ExampleIsExist() {
	if com.IsExist("file.go") {
		fmt.Println("file.go exists")
		return
	}
	fmt.Println("file.go is not a file or does not exist")
}

// ------------- END ------------

// ------------------------------
// dir.go
// ------------------------------

func ExampleIsDir() {
	if com.IsDir("files") {
		fmt.Println("directory 'files' exists")
		return
	}
	fmt.Println("'files' is not a directory or does not exist")
}

// ------------- END ------------

// ------------------------------
// string.go
// ------------------------------

func ExampleIsLetter() {
	fmt.Println(com.IsLetter('1'))
	fmt.Println(com.IsLetter('['))
	fmt.Println(com.IsLetter('a'))
	fmt.Println(com.IsLetter('Z'))
	// Output:
	// false
	// false
	// true
	// true
}

func ExampleExpand() {
	match := map[string]string{
		"domain":    "gowalker.org",
		"subdomain": "github.com",
	}
	s := "http://{domain}/{subdomain}/{0}/{1}"
	fmt.Println(com.Expand(s, match, "Unknwon", "gowalker"))
	// Output: http://gowalker.org/github.com/Unknwon/gowalker
}

// ------------- END ------------

// ------------------------------
// http.go
// ------------------------------

func ExampleHttpGet() ([]byte, error) {
	rc, err := com.HttpGet(&http.Client{}, "http://gowalker.org", nil)
	if err != nil {
		return nil, err
	}
	p, err := ioutil.ReadAll(rc)
	rc.Close()
	return p, err
}

// ------------- END ------------
