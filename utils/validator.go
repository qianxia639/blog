package utils

import (
	"errors"
	"reflect"
)

func Verify(value interface{}) error {
	tp := reflect.TypeOf(value)   // 获取类型
	val := reflect.ValueOf(value) // 获取值

	switch val.Kind() {
	case reflect.Struct:
		return validatorStruct(tp, val)
	case reflect.Map:
		return validatorMap(tp, val)
	}

	return nil
}

func validatorStruct(t reflect.Type, value reflect.Value) error {
	for i := 0; i < value.NumField(); i++ {
		tagVal := t.Field(i)
		v := value.Field(i)
		if isEmpty(v) {
			return errors.New(tagVal.Name + " cannot be empty")
		}
	}
	return nil
}

func validatorMap(t reflect.Type, value reflect.Value) error {
	return nil
}

// 非空校验
func isEmpty(value reflect.Value) bool {
	// Kind()获取到具体类型
	switch value.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
