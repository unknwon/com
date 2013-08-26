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
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var examplePrefix = `<!doctype html>
<html>
<head>
    <title>Example Domain</title>
`

func TestHttpGet(t *testing.T) {
	// 200.
	rc, err := HttpGet(&http.Client{}, "http://example.com", nil)
	if err != nil {
		t.Errorf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
	p, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", nil, err)
	}
	s := string(p)
	if !strings.HasPrefix(s, examplePrefix) {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", examplePrefix, s)
	}

	// 404.
	rc, err = HttpGet(&http.Client{}, "http://example.com/foo", nil)
	if err == nil {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", NotFoundError{}, nil)
	} else if _, ok := err.(NotFoundError); !ok {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", NotFoundError{}, err)
	}
}

func TestHttpGetBytes(t *testing.T) {
	p, err := HttpGetBytes(&http.Client{}, "http://example.com", nil)
	if err != nil {
		t.Errorf("HttpGetBytes:\n Expect => %v\n Got => %s\n", nil, err)
	}
	s := string(p)
	if !strings.HasPrefix(s, examplePrefix) {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", examplePrefix, s)
	}
}

func TestHttpGetJSON(t *testing.T) {

}

type rawFile struct {
	rawURL string
	data   []byte
}

func (rf *rawFile) RawURL() string {
	return rf.rawURL
}

func (rf *rawFile) Data() []byte {
	return rf.data
}

func (rf *rawFile) SetData(p []byte) {
	rf.data = p
}

func TestFetchFiles(t *testing.T) {
	files := []RawFile{
		&rawFile{rawURL: "http://example.com"},
		&rawFile{rawURL: "http://example.com"},
	}
	err := FetchFiles(&http.Client{}, files, nil)
	if err != nil {
		t.Errorf("FetchFiles:\n Expect => %v\n Got => %s\n", nil, err)
	} else if len(files[0].Data()) != 1270 {
		t.Errorf("FetchFiles:\n Expect => %d\n Got => %d\n", 1270, len(files[0].Data()))
	} else if len(files[1].Data()) != 1270 {
		t.Errorf("FetchFiles:\n Expect => %d\n Got => %d\n", 1270, len(files[1].Data()))
	}
}

func BenchmarkHttpGet(b *testing.B) {
	client := &http.Client{}
	for i := 0; i < b.N; i++ {
		HttpGet(client, "http://example.com", nil)
	}
}

func BenchmarkHttpGetBytes(b *testing.B) {
	client := &http.Client{}
	for i := 0; i < b.N; i++ {
		HttpGetBytes(client, "http://example.com", nil)
	}
}

func BenchmarkHttpGetJSON(b *testing.B) {
	client := &http.Client{}
	for i := 0; i < b.N; i++ {
		HttpGetJSON(client, "http://example.com", nil)
	}

}

func BenchmarkFetchFiles(b *testing.B) {
	files := []RawFile{
		&rawFile{rawURL: "http://example.com"},
		&rawFile{rawURL: "http://www.baidu.com"},
	}
	for i := 0; i < b.N; i++ {
		FetchFiles(&http.Client{}, files, nil)
	}
}
