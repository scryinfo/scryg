package sfile

import (
	"os"
)

//
func ExitFile(file string) bool {
	_, err := os.Stat(file)
	return (err == nil || os.IsExist(err))
}
