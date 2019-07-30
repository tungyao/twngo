package tnjson

import (
	"reflect"
	"strconv"
)

type JSON struct {
	json string
}

func (j *JSON) _format(obj interface{}) *JSON {
	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		f := reflect.ValueOf(obj).Elem().Interface().(map[string]interface{})
		for k, v := range f {
			if reflect.TypeOf(v).Kind() == reflect.Map {
				j.json += "\"" + k + "\":{"
				j._format(v)
			} else {
				switch reflect.TypeOf(v).Kind() {
				case reflect.String:
					j.json += "\"" + k + "\":" + "\"" + reflect.ValueOf(v).String() + "\"" + ","
				case reflect.Int:
					j.json += "\"" + k + "\":" + "\"" + strconv.Itoa(reflect.ValueOf(v).Interface().(int)) + "\"" + ","
				}
			}
		}
	} else {
		f := reflect.ValueOf(obj).Interface().(map[string]interface{})
		for k, v := range f {
			if reflect.TypeOf(v).Kind() == reflect.Map {
				j._format(v)
			} else {
				switch reflect.TypeOf(v).Kind() {
				case reflect.String:
					j.json += "\"" + k + "\":" + "\"" + reflect.ValueOf(v).String() + "\""
				case reflect.Int:
					j.json += "\"" + k + "\":" + "\"" + strconv.Itoa(reflect.ValueOf(v).Interface().(int)) + "\""
				}
				j.json += "},"

			}
		}
	}
	return j
}
func (j *JSON) Encode(obj interface{}) string {
	js := j._format(obj).json
	js = js[:len(js)-1]
	return "{" + js + "}"
}
func (j *JSON) Decode() interface{} {

	return 0
}
