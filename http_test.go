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
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", nil, err)
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
	rc, err = HttpGet(&http.Client{}, "https://github.com/foooooooooo", nil)
	if err == nil {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", NotFoundError{}, nil)
	} else if _, ok := err.(NotFoundError); !ok {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", NotFoundError{}, err)
	}
}

func TestHttpGetBytes(t *testing.T) {
	p, err := HttpGetBytes(&http.Client{}, "http://example.com", nil)
	if err != nil {
		t.Errorf("HttpGetBytes:\n Expect => %s\n Got => %s\n", nil, err)
	}
	s := string(p)
	if !strings.HasPrefix(s, examplePrefix) {
		t.Errorf("HttpGet:\n Expect => %s\n Got => %s\n", examplePrefix, s)
	}
}

func TestHttpGetJSON(t *testing.T) {

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
