package uuid

import "github.com/satori/go.uuid"

func GetUuid() string {
	// 创建
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}
