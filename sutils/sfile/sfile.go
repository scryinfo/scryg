// Scry Info.  All rights reserved.
// license that can be found in the license.txt file.

package sfile

import (
	"os"
)

// 判断文件或文件夹是否存在， true为存在
func ExitFile(file string) bool {
	_, err := os.Stat(file)
	return (err == nil || os.IsExist(err))
}
