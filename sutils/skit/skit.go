package skit

import (
	"fmt"
	"reflect"
)
//检查interface最终指向的对象是否为空
func IsNil(any interface{}) bool {
	fmt.Println()
	re := false
	if any != nil {
		v := reflect.ValueOf(any)

		if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			re = v.IsNil()
			if !re {
				for {
					//fmt.Println(v.Type())
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
	}
	return re
}


//返回t最终的类型（非指针，非interface）
//todo test it
func RealType(t interface{}) (ty reflect.Type) {
	ty = nil
	v := reflect.ValueOf(t)
	if !v.IsNil() {
		if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			for {
				v = v.Elem()
				if (v.Kind() != reflect.Ptr && v.Kind() != reflect.Interface) || v.IsNil() {
					break
				}
			}
		}
	}
	ty = v.Type()
	return
}
