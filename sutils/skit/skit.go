// Scry Info.  All rights reserved.
// license that can be found in the license file.

package skit

import (
	"errors"
	"reflect"
	"unsafe"
)

//check if the final object pointed by interface is empty
func IsNil(any interface{}) bool {
	re := true
	if any != nil {
		v := reflect.ValueOf(any)
		switch v.Kind() {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
			re = v.IsNil()
		default:
			re = false
			return re
		}
		if !re {
			for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface { //如果为指针或interface类型，要检查指向的值
				v = v.Elem() //Ptr或Interface时，返回内部的值
				switch v.Kind() {
				case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
					re = v.IsNil()
				default:
					re = false
					return re
				}
				if re {
					break
				}
			}
		}
	}
	return re
}

//return the final type of t (non-pointer, non-interface)
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

// only support the field type of the pointer and interface
// field is the type of the pointer and interface
//newValue is the corresponding type
// the code depends on reflect.Value which contains the field name 'ptr'of pointer
func SetPrivateField(field *reflect.Value, newValue interface{}) error {

	var err error = nil
	for {
		if !field.CanAddr() {
			err = errors.New("can not get addr")
			break
		}

		if field.Kind() == reflect.Ptr { //Pointer
			fpp := ((**uintptr)(unsafe.Pointer(field.Addr().Pointer()))) //get the address of the field and transfer to the Pointer
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

			//{ // method one， require to decide the type of interface at complile time
			//	var t2 sampleInterface = &sampleImp{F: 20}
			//	fp2 := (*interface{})(unsafe.Pointer(&t2))
			//	*fpp = *fp2
			//}
			{ // method two，generally use the reflection implementation

				vf2 := reflect.ValueOf(newValue)
				if vf2.Elem().Type().AssignableTo(field.Type()) {
					err = errors.New("new value is not interface")
					break
				}

				vf2 = vf2.Convert(field.Type())                //must use convert the field type, because the type of t2 is Interface {},rather than Inter2类型
				ptr := reflect.ValueOf(vf2).FieldByName("ptr") //get ptr value by the reflection,pointer cannot be used and it would cause to panic because the t2 is the pointer type
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

//MergeClone marge and clone
func MergeClone(a []byte, b []byte) []byte {
	var clone []byte
	alen, blen := len(a), len(b)
	switch alen + blen {
	case blen:
		clone = append(b[:0:0], b...)
	case alen:
		clone = append(a[:0:0], a...)
	default:
		clone = append(a[:0:alen], a...) //do not clone
		clone = append(clone, b...)      //clone certainly,because alen and blen  are both not zero,   only alloc memory once
	}
	return clone
}
