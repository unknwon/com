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

package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsFile(t *testing.T) {
	if !IsFile("file.go") {
		t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", true, false)
	}

	if IsFile("testdata") {
		t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", false, true)
	}

	if IsFile("files.go") {
		t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", false, true)
	}
}

func TestIsExist(t *testing.T) {
	Convey("Check if file or directory exists", t, func() {
		Convey("Pass a file name that exists", func() {
			So(IsExist("file.go"), ShouldEqual, true)
		})
		Convey("Pass a directory name that exists", func() {
			So(IsExist("testdata"), ShouldEqual, true)
		})
		Convey("Pass a directory name that does not exist", func() {
			So(IsExist(".hg"), ShouldEqual, false)
		})
	})
}

func TestSaveFile(t *testing.T) {
	s := "TestSaveFile"
	n, err := SaveFile("testdata/SaveFile.txt", []byte(s))
	if err != nil {
		t.Errorf("SaveFile:\n Expect => %v\n Got => %v\n", nil, err)
	} else if n == 0 {
		t.Errorf("SaveFile:\n Expect => %d\n Got => %d\n", len(s), n)
	}
}

func TestSaveFileS(t *testing.T) {
	s := "TestSaveFileS"
	n, err := SaveFileS("testdata/SaveFileS.txt", s)
	if err != nil {
		t.Errorf("SaveFileS:\n Expect => %v\n Got => %v\n", nil, err)
	} else if n == 0 {
		t.Errorf("SaveFileS:\n Expect => %d\n Got => %d\n", len(s), n)
	}
}

func TestReadFile(t *testing.T) {
	b, err := ReadFile("testdata/SaveFile.txt")
	if err != nil {
		t.Errorf("ReadFile:\n Expect => %v\n Got => %v\n", nil, err)
	} else if string(b) != "TestSaveFile" {
		t.Errorf("ReadFile:\n Expect => %s\n Got => %s\n", "TestSaveFile", string(b))
	}
}

func TestReadFileS(t *testing.T) {
	s, err := ReadFileS("testdata/SaveFileS.txt")
	if err != nil {
		t.Errorf("ReadFileS:\n Expect => %v\n Got => %v\n", nil, err)
	} else if s != "TestSaveFileS" {
		t.Errorf("ReadFileS:\n Expect => %s\n Got => %s\n", "TestSaveFileS", s)
	}
}

func TestUnzip(t *testing.T) {

}

func TestUnTarGz(t *testing.T) {

}

func BenchmarkIsFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsFile("file.go")
	}
}

func BenchmarkIsExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExist("file.go")
	}
}

func BenchmarkSaveFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SaveFile("testdata/SaveFile.txt", []byte("TestSaveFile"))
	}
}

func BenchmarkSaveFileS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SaveFileS("testdata/SaveFileS.txt", "TestSaveFileS")
	}
}

func BenchmarkUnzip(b *testing.B) {

}

func BenchmarkUnTarGz(b *testing.B) {

}
