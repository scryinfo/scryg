// Scry Info.  All rights reserved.
// license that can be found in the license file.

package skit

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type sampleT2 struct {
	f2 int
}

type sampleInterface interface {
	Get() int
}

type sampleT struct {
	fPtr       *sampleT2
	fInterface sampleInterface
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

func TestIsNil(t *testing.T) {
	var d *int = nil
	var d2 interface{} = d
	var d3 interface{} = &d2
	var d4 interface{} = &d3

	r := IsNil(1)
	assert.Equal(t, false, r)

	r = IsNil(nil)
	assert.Equal(t, true, r)
	r = IsNil(d)
	assert.Equal(t, true, r)
	r = IsNil(d2)
	assert.Equal(t, true, r)
	r = IsNil(d3)
	assert.Equal(t, true, r)
	r = IsNil(d4)
	assert.Equal(t, true, r)

	var dd = 1
	d = &dd
	r = IsNil(d)
	assert.Equal(t, false, r)
	r = IsNil(d2)
	assert.Equal(t, true, r)
	r = IsNil(d3)
	assert.Equal(t, true, r)
	r = IsNil(d4)
	assert.Equal(t, true, r)
}

func TestMergeClone(t *testing.T) {
	data1 := []byte{1}
	data2 := []byte{4}
	clone := MergeClone(data1, data2)
	assert.Equal(t, data1, clone[0:len(data1)])
	assert.Equal(t, data2, clone[len(data1):])
	assert.NotSame(t, &data1[0], &clone[0]) //cloned

	data2 = nil
	clone = MergeClone(data1, data2)
	assert.Equal(t, data1, clone[0:len(data1)])
	assert.Equal(t, data1, clone)
	assert.NotSame(t, &data1[0], &clone[0]) //cloned

	data2 = nil
	data1 = nil
	clone = MergeClone(data1, data2)
	assert.Equal(t, data1, clone)

	data1 = nil
	data2 = []byte{4}
	clone = MergeClone(data1, data2)
	assert.Equal(t, data2, clone[0:len(data2)])
	assert.NotSame(t, &data2[0], &clone[0]) //cloned
}
