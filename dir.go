
package com

import (
	"os"
)

// check is dir
func IsDir(dir string)bool{
	f, e := os.Stat(dir)
	if  e != nil {
		return false
	}
	return f.IsDir()
}
