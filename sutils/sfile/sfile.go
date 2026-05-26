// Scry Info.  All rights reserved.
// license that can be found in the license file.

package sfile

import (
	"os"
)

// judge if the file or folder exsits, true means existence
func ExistFile(name string) bool {
	cur, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	if cur != nil {
		return !cur.IsDir()
	}
	return false
}
func ExistDir(name string) bool {
	cur, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	if cur != nil {
		return cur.IsDir()
	}
	return false
}

func Exist(name string) bool {
	_, err := os.Stat(name)
	return (err == nil || os.IsExist(err))
}
