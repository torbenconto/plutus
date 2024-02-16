package tests

import "reflect"

func GetField(s interface{}, field string) interface{} {
	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}
