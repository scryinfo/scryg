package skit

import (
	"reflect"
)

func IsNil(any interface{}) bool {
	re := false
	if any != nil {
		v := reflect.ValueOf(any)

		if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			re = v.IsNil()
			if !re {
				for {
					v2 := v.Elem()
					if v2.Kind() != reflect.Ptr && v2.Kind() != reflect.Interface {
						break
					}
					re = v2.IsNil()
					if re {
						break
					}
					v = v2
				}
			}

		}
	} else {
		re = true
	}
	return re
}

func RealType(t interface{}) (ty reflect.Type) {
	ty = nil
	//tt := reflect.TypeOf(t)
	return
}
