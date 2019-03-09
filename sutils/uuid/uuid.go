package uuid

import (
	"github.com/google/uuid"
)

func GetUuid() string {
	// 创建
	u1 := uuid.New()
	return u1.String()
}
