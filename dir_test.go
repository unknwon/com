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
)

func TestIsDir(t *testing.T) {
	if IsDir("file.go") {
		t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", false, true)
	}

	if !IsDir("testdata") {
		t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", true, false)
	}

	if IsDir("foo") {
		t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", false, true)
	}
}

func BenchmarkIsDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsDir("file.go")
	}
}
