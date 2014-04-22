package com

import (
	"encoding/base64"
	"net/url"
)

// url encode string, is + not %20
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// url decode string
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// base64 encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64 decode
func Base64Decode(str string) (string, error) {
	s, e := base64.StdEncoding.DecodeString(str)
	return string(s), e
}
