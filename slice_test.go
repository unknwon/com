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
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAppendStr(t *testing.T) {
	s := []string{"a"}
	s = AppendStr(s, "a")
	s = AppendStr(s, "b")
	sR := []string{"a", "b"}
	if !CompareSliceStr(s, sR) {
		t.Errorf("AppendStr:\n Expect => %s\n Got => %s\n", sR, s)
	}
}

func TestCompareSliceStr(t *testing.T) {
	Convey("Compares two 'string' type slices with elements and indexes", t, func() {
		So(CompareSliceStr(
			[]string{"1", "2", "3"}, []string{"1", "2", "3"}), ShouldEqual, true)
	})
}

func BenchmarkAppendStr(b *testing.B) {
	s := []string{"a"}
	for i := 0; i < b.N; i++ {
		s = AppendStr(s, fmt.Sprint(b.N%3))
	}
}

func BenchmarkCompareSliceStr(b *testing.B) {
	s1 := []string{"1", "2", "3"}
	s2 := []string{"1", "2", "3"}
	for i := 0; i < b.N; i++ {
		CompareSliceStr(s1, s2)
	}
}
