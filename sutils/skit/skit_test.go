// Scry Info.  All rights reserved.
// license that can be found in the license file.
package skit

import (
	"reflect"
	"testing"
)

type sampleT struct {
	fPtr       *sampleT2
	fInterface sampleInterface
}

type sampleT2 struct {
	f2 int
}

type sampleInterface interface {
	Get() int
}

type sampleImp struct {
	F int
}

func (c *sampleImp) Get() int {
	return c.F
}

func TestSetPrivateField(t *testing.T) {

	{ //pointer
		st := &sampleT{fPtr: nil, fInterface: nil}
		field := reflect.ValueOf(st).Elem().FieldByName("fPtr")

		err := SetPrivateField(&field, &sampleT2{f2: 10})
		if err != nil || st.fPtr == nil || st.fPtr.f2 != 10 {
			t.Error("pointer fail")
		}
	}
	{ //interface
		st := &sampleT{fPtr: nil, fInterface: nil}
		field := reflect.ValueOf(st).Elem().FieldByName("fInterface")
		err := SetPrivateField(&field, &sampleImp{F: 15})
		if err != nil || st.fInterface == nil || st.fInterface.Get() != 15 {
			t.Error("interface fail")
		}

		var inter2 sampleInterface = &sampleImp{F: 12}
		err = SetPrivateField(&field, inter2)
		if err != nil || st.fInterface == nil || st.fInterface.Get() != 12 {
			t.Error("interface fail")
		}

		var inter3 interface{} = &sampleImp{F: 11}
		err = SetPrivateField(&field, inter3)
		if err != nil || st.fInterface == nil || st.fInterface.Get() != 11 {
			t.Error("interface fail")
		}
	}

}
