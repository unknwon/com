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
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetGOPATHs returns all paths in GOPATH variable.
func GetGOPATHs() []string {
	gopath := os.Getenv("GOPATH")
	var paths []string
	if runtime.GOOS == "windows" {
		gopath = strings.Replace(gopath, "\\", "/", -1)
		paths = strings.Split(gopath, ";")
	} else {
		paths = strings.Split(gopath, ":")
	}
	return paths
}

// GetSrcPath returns app. source code path.
// It only works when you have src. folder in GOPATH,
// it returns error not able to locate source folder path.
func GetSrcPath(importPath string) (string, error) {
	// Look up executable in PATH variable.
	appPath, err := exec.LookPath(path.Base(os.Args[0]))
	if err != nil {
		return "", err
	}

	paths := GetGOPATHs()
	for _, p := range paths {
		if IsExist(p + "/src/" + importPath + "/") {
			appPath = p + "/src/" + importPath + "/"
			break
		}
	}

	if len(appPath) == 0 {
		return "", errors.New("Unable to locate source folder path")
	}

	appPath = filepath.Dir(appPath) + "/"
	if runtime.GOOS == "windows" {
		// Replace all '\' to '/'.
		appPath = strings.Replace(appPath, "\\", "/", -1)
	}

	return appPath, nil
}
