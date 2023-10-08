package protobufparser

import (
	"errors"
	"reflect"
	"unsafe"
)

type ValueType byte

const (
	VARINT ValueType = 0 //varint := int32 | int64 | uint32 | uint64 | bool | enum | sint32 | sint64
	String ValueType = 1 //string := valid UTF-8 string (e.g. ASCII)
	I32    ValueType = 2 //i32 := sfixed32 | fixed32 | float
	I64    ValueType = 3 //i64 := sfixed64 | fixed64 | double
	//len-prefix := size (message | string | bytes | packed)
	//bytes := any sequence of 8-bit bytes
	//packed := varint* | i32* | i64*
)

type BaseValue struct {
	ThisObject interface{}
}

func NewBaseValue() *BaseValue {
	return &BaseValue{}
}
func ToBaseValue(object interface{}) (dst *BaseValue, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	var isStruct = false
	var v reflect.Value
	switch reflect.TypeOf(object).Kind() {
	case reflect.Ptr, reflect.Interface:
		v = reflect.ValueOf(object).Elem().FieldByName("BaseValue")
	case reflect.Struct:
		isStruct = true
		v = reflect.ValueOf(object).FieldByName("BaseValue")
	default:
		return nil, errors.New("ToBaseValue: object type error: " + reflect.TypeOf(object).Kind().String())
	}
	switch v.Kind() {
	case reflect.Ptr:
		return (*BaseValue)(unsafe.Pointer(v.Pointer())), err
	case reflect.Struct:
		if isStruct {
			return nil, errors.New("ToBaseValue: At least one of the parent and the child is a pointer")
		}
		return (*BaseValue)(unsafe.Pointer(v.Addr().Pointer())), err
	default:
		return nil, errors.New("ToBaseValue: BaseValue type error: " + v.Kind().String())
	}

}
