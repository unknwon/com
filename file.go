package com

import (
	"path"
	"os"
	"io"
	"io/ioutil"
)

// get filepath base name
func Basename(file string) string {
	return path.Base(file)
}

// get filepath dir name
func Dirname(file string) string {
	return path.Dir(file)
}

// get absolute filepath, based on built excutable file
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

// check is file
func IsFile(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}

// copy file to new path
func Copy(file string, to string) (bool, error) {
	sf, e := os.Open(file)
	if e != nil {
		return false, e
	}
	defer sf.Close()
	df, e2 := os.Create(to)
	if e2 != nil {
		return false, e2
	}
	defer df.Close()
	// buffer reader, do chunk transfer
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := sf.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := df.Write(buf[:n]); err != nil {
			return false, err
		}
	}
	return true, nil
}

// move file to new path
func Move(file string, to string) (bool, error) {
	// copy
	r, e := Copy(file, to)
	if e != nil {
		return r, e
	}
	// then remove
	e = os.Remove(file)
	if e != nil {
		r = false
	}
	return r, e
}

// delete file
func Unlink(file string) error {
	return os.Remove(file)
}

// rename file name
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// put string to file
func FilePutContent(file string, content string) (int, error) {
	fs, e := os.Create(file)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}

// get string from text file
func FileGetContent(file string) (string, error) {
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}
