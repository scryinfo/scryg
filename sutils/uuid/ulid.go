package uuid

import "github.com/oklog/ulid/v2"

// GetUuidForDB return uid for DB primary key
func GetUuidForDB() string {
	return ulid.Make().String()
}
