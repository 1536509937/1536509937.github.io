package utils

import (
	"1536509937/ku-bbs/pkg/utils/str"
	"1536509937/ku-bbs/pkg/utils/time"
	"1536509937/ku-bbs/pkg/utils/view"
	"html/template"
	"reflect"
	"strings"
)

func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"DiffForHumans":    time.DiffForHumans,
		"ToDateTimeString": time.ToDateTimeString,
		"Html":             view.Html,
		"RemindName":       view.RemindName,
		"StrLimit":         str.Limit,
		"StrJoin":          strings.Join,
	}
}

func StructToMap(s interface{}) map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
