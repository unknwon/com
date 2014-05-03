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
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
)

const (
	Byte  = 1
	KByte = Byte * 1024
	MByte = KByte * 1024
	GByte = MByte * 1024
	TByte = GByte * 1024
	PByte = TByte * 1024
	EByte = PByte * 1024
)

var bytesSizeTable = map[string]uint64{
	"b":  Byte,
	"kb": KByte,
	"mb": MByte,
	"gb": GByte,
	"tb": TByte,
	"pb": PByte,
	"eb": EByte,
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// HumaneFileSize calculates the file size and generate user-friendly string.
func HumaneFileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}

// get absolute filepath, based on built executable file
func RealPath(file string) (string, error) {
	if path.IsAbs(file) {
		return file, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, file), err
}

// get file modified time
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// Copy copies file from source to target path.
// It returns false and error when error occurs in underlying functions.
func Copy(srcPath, destPath string) error {
	sf, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sf.Close()

	si, err := os.Lstat(srcPath)
	if err != nil {
		return err
	}

	// Symbolic link.
	if si.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(srcPath)
		if err != nil {
			return err
		}
		return os.Symlink(target, destPath)
	}

	df, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer df.Close()

	// buffer reader, do chunk transfer
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := sf.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := df.Write(buf[:n]); err != nil {
			return err
		}
	}

	return os.Chmod(destPath, si.Mode())
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// SaveFile saves content type '[]byte' to file by given path.
// It returns error when fail to finish operation.
func SaveFile(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

// SaveFileS saves content type 'string' to file by given path.
// It returns error when fail to finish operation.
func SaveFileS(filePath string, s string) (int, error) {
	return SaveFile(filePath, []byte(s))
}

// ReadFile reads data type '[]byte' from file by given path.
// It returns error when fail to finish operation.
func ReadFile(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

// ReadFileS reads data type 'string' from file by given path.
// It returns error when fail to finish operation.
func ReadFileS(filePath string) (string, error) {
	b, err := ReadFile(filePath)
	return string(b), err
}

// Unzip unzips .zip file to 'destPath' and returns sub-directories.
// It returns error when fail to finish operation.
func Unzip(srcPath, destPath string) ([]string, error) {
	// Open a zip archive for reading
	r, err := zip.OpenReader(srcPath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	dirs := make([]string, 0, 5)
	// Iterate through the files in the archive
	for _, f := range r.File {
		// Get files from archive
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}

		dir := path.Dir(f.Name)
		// Create directory before create file
		os.MkdirAll(destPath+"/"+dir, os.ModePerm)
		dirs = AppendStr(dirs, dir)

		if f.FileInfo().IsDir() {
			continue
		}

		// Write data to file
		fw, _ := os.Create(path.Join(destPath, f.Name))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(fw, rc)

		if fw != nil {
			fw.Close()
		}
		if err != nil {
			return nil, err
		}
	}
	return dirs, nil
}

// UnTarGz ungzips and untars .tar.gz file to 'destPath' and returns sub-directories.
// It returns error when fail to finish operation.
func UnTarGz(srcFilePath string, destDirPath string) ([]string, error) {
	// Create destination directory
	os.Mkdir(destDirPath, os.ModePerm)

	fr, err := os.Open(srcFilePath)
	if err != nil {
		return nil, err
	}
	defer fr.Close()

	// Gzip reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	// Tar reader
	tr := tar.NewReader(gr)

	dirs := make([]string, 0, 5)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}

		// Check if it is directory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create directory before create file
			dir := path.Dir(hdr.Name)
			os.MkdirAll(destDirPath+"/"+dir, os.ModePerm)
			dirs = AppendStr(dirs, dir)

			// Write data to file
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				return nil, err
			}
		}
	}
	return dirs, nil
}
