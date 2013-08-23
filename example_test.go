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
	// Output:
	// <div id=\"button\" class=\"btn\">Click me</div>\n
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
	// Output:
	// /Users/Joe/Appliations/Go/src/github.com/Unknwon/com
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
	fmt.Println("file.go does not exists")
}

// ------------- END ------------
