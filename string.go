package com

import (
	"strings"
	"crypto/md5"
	"crypto/sha1"
	"io"
	"fmt"
	"unicode"
)

// explode string with proper chars
func Explode(str string, sep string) []string {
	return strings.Split(str, sep)
}

// join string array to string with connection chars
func Join(str []string, sep string) string {
	return strings.Join(str, sep)
}

// substring from start position and belong length
func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// md5 hash string
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

// sha1 hash string
func Sha1(str string) string {
	m := sha1.New()
	io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

// trim space on left
func Ltrim(str string) string {
	return strings.TrimLeftFunc(str, unicode.IsSpace)
}

// trim space on right
func Rtrim(str string) string {
	return strings.TrimRightFunc(str, unicode.IsSpace)
}

// trim space in all string lenth
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// repeat string times
func StrRepeat(str string, times int) string {
	return strings.Repeat(str, times)
}

// replace find all occurs to string
func StrReplace(str string, find string, to string) string {
	return strings.Replace(str, find, to, -1)
}

// find the first occur position, if not found return -1
func StrPos(str string, find string) int {
	return strings.Index(str, find)
}

// convert to upper
func StrToUpper(str string) string {
	return strings.ToUpper(str)
}

// convert to lower
func StrToLower(str string) string {
	return strings.ToLower(str)
}

// convert first letter to upper
func UcFirst(str string) string {
	return strings.Title(str)
}
