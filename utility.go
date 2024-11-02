package main

import (
	"errors"
	"reflect"
	"strings"
	"time"
	as "github.com/aerospike/aerospike-client-go/v7"
)



// binsToStruct converts Aerospike record bins to a struct
func binsToStruct(record *as.Record, result interface{}) error {
	if record == nil {
		return errors.New("nil record")
	}

	v := reflect.ValueOf(result)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("result must be a non-nil pointer")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("as")
		if tag == "" {
			continue
		}

		binName := strings.Split(tag, ",")[0]
		binValue, exists := record.Bins[binName]
		if !exists {
			continue
		}

		fieldValue := v.Field(i)
		if !fieldValue.CanSet() {
			continue
		}

		// Handle different types
		switch fieldValue.Kind() {
		case reflect.String:
			if str, ok := binValue.(string); ok {
				fieldValue.SetString(str)
			}
		case reflect.Int, reflect.Int64:
			if num, ok := binValue.(int); ok {
				fieldValue.SetInt(int64(num))
			}
		case reflect.Float64:
			if num, ok := binValue.(float64); ok {
				fieldValue.SetFloat(num)
			}
		case reflect.Struct:
			if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
				if timeVal, ok := binValue.(time.Time); ok {
					fieldValue.Set(reflect.ValueOf(timeVal))
				}
			}
		}
	}

	return nil
}
