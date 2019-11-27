// Scry Info.  All rights reserved.
// license that can be found in the license file.

package sfile

import (
	"os"
)

// judge if the file or folder exsits, true means existence
func ExistFile(file string) bool {
	_, err := os.Stat(file)
	return (err == nil || os.IsExist(err))
}

//如果存在且是文件，返回true
//注： !IsDir  !=  IsFile 因为还有出错的情况， 所以写了两个函数
func IsFile(str string) bool {
	fileInfo, err := os.Stat(str)
	return err == nil && !fileInfo.IsDir()
}

//如果存在且是目录，返回true
//注： !IsDir  !=  IsFile 因为还有出错的情况， 所以写了两个函数
func IsDir(str string) bool {
	fileInfo, err := os.Stat(str)
	return err == nil && fileInfo.IsDir()
}
