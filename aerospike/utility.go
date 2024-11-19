package aerospike_db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/aerospike/aerospike-client-go/v7"
)

func BinsToStruct(record *aerospike.Record, result interface{}) error {


	// fmt.Print("before record %v", record)

	if record == nil {
		return errors.New("nil record")
	}

	

	v := reflect.ValueOf(result)

	

	if v.Kind() != reflect.Ptr || v.IsNil() {

		return errors.New("result must be a non-nil pointer")
	}

	

	v = v.Elem()
	t := v.Type()

	// fmt.Print("after elem & type of  %v", record)
	

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("as")
		if tag == "" {
			continue
		}

		binName := strings.Split(tag, ",")[0]
		fmt.Printf(" binName %v", binName)
		binValue, exists := record.Bins[binName]

		fmt.Printf(" binValue %v", binValue)

		if !exists {
			continue
		}

		fieldValue := v.Field(i)
		if !fieldValue.CanSet() {
			continue
		}

		// Handle different types with better type checking
		switch fieldValue.Kind() {
		case reflect.String:
			if str, ok := binValue.(string); ok {
				fieldValue.SetString(str)
			}
		case reflect.Int, reflect.Int64:
			switch num := binValue.(type) {
			case int:
				fieldValue.SetInt(int64(num))
			case int64:
				fieldValue.SetInt(num)
			default:
				continue // Type mismatch, do nothing
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
		default:
			continue // Unsupported type, do nothing
		}
	}

	return nil
}

func StructToBins(data interface{}) (aerospike.BinMap, error) {
	bins := aerospike.BinMap{}
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("input must be a struct or pointer to struct")
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("as")
		if tag == "" {
			continue
		}

		tagParts := strings.Split(tag, ",")
		binName := tagParts[0]
		omitempty := len(tagParts) > 1 && tagParts[1] == "omitempty"

		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				if !omitempty {
					return nil, errors.New("nil value for required field: " + field.Name)
				}
				continue
			}
			value = value.Elem()
		}

		if omitempty && value.IsZero() {
			continue
		}

		switch value.Kind() {
		case reflect.String, reflect.Int, reflect.Int64, reflect.Float64, reflect.Bool:
			bins[binName] = value.Interface()
		case reflect.Struct:
			if value.Type() == reflect.TypeOf(time.Time{}) {
				bins[binName] = value.Interface()
			} else {
				nestedBins, err := StructToBins(value.Interface())
				if err != nil {
					return nil, err
				}
				for k, v := range nestedBins {
					bins[binName+"_"+k] = v
				}
			}
		default:
			fmt.Printf("value type : ", value.Type())
			return nil, errors.New("unsupported field type: " + field.Name)
		}
	}

	return bins, nil
}