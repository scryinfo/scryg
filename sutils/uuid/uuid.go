// Scry Info.  All rights reserved.
// license that can be found in the license file.

package uuid

import (
	"github.com/google/uuid"
)

// GetUuid return uuid
func GetUuid() string {
	u1 := uuid.New()
	return u1.String()
}
