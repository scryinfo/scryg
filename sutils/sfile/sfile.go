// Scry Info.  All rights reserved.
// license that can be found in the license file.

package sfile

import (
	"os"
)

// judge if the file or folder exsits, true means existence
func ExitFile(file string) bool {
	_, err := os.Stat(file)
	return (err == nil || os.IsExist(err))
}
