package utils

import "reflect"

func IsUInt(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Uint {
		return true
	}
	return false
}

func IsUIntArray(value interface{}) bool {
	t := reflect.TypeOf(value)
	if !IsArray(value) {
		return false
	}
	if t.Elem().Kind() == reflect.Uint {
		return true
	}
	return false
}

func IsString(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.String {
		return true
	}
	return false
}

func IsStringArray(value interface{}) bool {
	t := reflect.TypeOf(value)
	if !IsArray(value) {
		return false
	}
	if t.Elem().Kind() == reflect.String {
		return true
	}
	return false
}

func IsArray(value interface{}) bool {
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		return false
	}
	return true
}
