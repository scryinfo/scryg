package skit

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
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

//只支持指针与interface类型的字段
//field 是指针或interface类型
//newValue是对应的类型
// 代码依赖于 reflect.Value中存放指针的字段名为 “ptr”
func SetPrivateField(field *reflect.Value, newValue interface{}) error {

	var err error = nil
	for {
		if !field.CanAddr() {
			err = errors.New("can not get addr")
			break
		}

		if field.Kind() == reflect.Ptr { //指针
			fpp := ((**uintptr)(unsafe.Pointer(field.Addr().Pointer()))) //得到字段的地址, 转换为指指针
			vf2 := reflect.ValueOf(newValue)
			if vf2.Kind() != reflect.Ptr || (!vf2.Type().AssignableTo(field.Type())) {
				err = errors.New("new value is not pointer")
				break
			}
			fp2 := (*uintptr)(unsafe.Pointer(vf2.Pointer()))
			*fpp = fp2
			break
		}

		if field.Kind() == reflect.Interface { //interface
			fpp := ((*interface{})(unsafe.Pointer(field.Addr().Pointer())))

			//{ // 方式一，需要在编译时确定interface的类型
			//	var t2 sampleInterface = &sampleImp{F: 20}
			//	fp2 := (*interface{})(unsafe.Pointer(&t2))
			//	*fpp = *fp2
			//}
			{ // 方式二，通用使用反射实现

				vf2 := reflect.ValueOf(newValue)
				if vf2.Elem().Type().AssignableTo(field.Type()) {
					err = errors.New("new value is not interface")
					break
				}

				vf2 = vf2.Convert(field.Type()) //一定使用Convert 转换为字段的类型， 因为t2的中的类型 为interface{},  不是 Inter2类型
				ptr := reflect.ValueOf(vf2).FieldByName("ptr") //通过反射取到 ptr的值， 这里不能使用Pointer，会panic，因为 t2不是指针类型
				fp2 := (*interface{})(unsafe.Pointer(ptr.Pointer()))
				*fpp = *fp2
			}

			break
		}

		err = errors.New("not support type")
		break
	}
	return err
}
