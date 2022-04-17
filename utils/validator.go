package utils

import (
	"errors"
	"reflect"
	"regexp"
)

func Verify(value interface{}) error {
	tp := reflect.TypeOf(value)   // 获取类型
	val := reflect.ValueOf(value) // 获取值

	k := val.Kind()
	if k != reflect.Struct {
		return errors.New("Not a Struct")
	}

	// 遍历结构体中所有字段
	for i := 0; i < val.NumField(); i++ {
		tagVal := tp.Field(i)
		v := val.Field(i)

		if isEmpty(v) {
			return errors.New(tagVal.Name + "不能为空")
		}
		if tagVal.Name == "Email" {
			if !regexpMatch(v.String()) {
				return errors.New(tagVal.Name + "格式不匹配")
			}
		}
	}
	return nil
}

// 非空校验
func isEmpty(value reflect.Value) bool {
	// Kind()获取到具体类型
	switch value.Kind() {
	case reflect.String:
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

// 正则校验
func regexpMatch(matchStr string) bool {
	return regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-z0-9]+`).MatchString(matchStr)
}
