package helper

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func TerraformToStruct(source interface{}, target interface{}) {
	tType := reflect.TypeOf(target)
	if tType.Kind() != reflect.Ptr {
		panic("target is not pointer")
	}
	tType = tType.Elem()
	//log.Println("tType.Kind:", tType.Kind())

	tVal := reflect.ValueOf(target).Elem()
	//log.Println("tVal.Kind:", tVal.Kind())

	switch input := source.(type) {
	case []interface{}:
		//log.Println("list")
		if len(input) == 1 {
			TerraformToStruct(input[0], target)
		}

	case map[string]interface{}:
		//log.Println("map")

		for i := 0; i < tType.NumField(); i++ {
			f := tType.Field(i)
			name := getTerraformFieldName(f)
			if name == "" {
				// skip if dest targetValue have no tag
				continue
			}
			//name = strings.Split(name, ",")[0]
			val, ok := input[name]
			if !ok {
				// skip if no field in data source
				continue
			}

			sourceValue := reflect.ValueOf(val)
			if sourceValue.Type().Kind() == reflect.Slice && sourceValue.Len() == 0 {
				// skip empty []interface{}
				continue
			}


			targetValue := tVal.Field(i)
			//log.Printf("val: %+v, type: %t\n", val, val)
			targetValue.Set(convertVal(sourceValue, targetValue))
		}
	}
}

func getTerraformFieldName(field reflect.StructField) string {
	var name string
	if val, ok := field.Tag.Lookup("terraform"); ok {
		name = val
	} else if val, ok := field.Tag.Lookup("json"); ok {
		name = val
	}
	name = strings.Split(name, ",")[0]
	return toSnakeCase(name)
}

func convertVal(from, to reflect.Value) reflect.Value {
	//log.Println("from >> Kind:", from.Kind(), "Type:", from.Type())
	//log.Println("to >> Kind:", to.Kind(), "Type:", to.Type())

	switch from.Type().Kind() {
	case reflect.Int, reflect.Int64:
		return reflect.ValueOf(from.Int())
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(from.Float())

	case reflect.String, reflect.Bool:
		return from

	case reflect.Slice:
		if to.Type().Kind() == reflect.Slice {
			slice := reflect.MakeSlice(to.Type(), from.Len(), from.Len())
			for i := 0; i < from.Len(); i++ {
				// convert source slice items and set to new slice
				var item reflect.Value

				switch to.Type().Elem().Kind() {
				// case when we have basic types
				case reflect.String, reflect.Bool, reflect.Int, reflect.Int64, reflect.Float64:
					item = from.Index(i).Elem()
				default:
					// case when we have pointer elements
					if to.Type().Elem().Kind() == reflect.Ptr {
						item = reflect.New(to.Type().Elem().Elem())
					} else {
						item = reflect.New(to.Type().Elem())
					}

					item = convertVal(from.Index(i), item)
				}

				// if target slice is not slice of pointers
				if item.Type().Kind() == reflect.Ptr && to.Type().Elem().Kind() != reflect.Ptr {
					item = item.Elem()
				}

				slice.Index(i).Set(item)
			}

			return slice
		}

	case reflect.Map:
		if to.Type().Kind() == reflect.Map {
			m := reflect.MakeMapWithSize(to.Type(), from.Len())
			for _, key := range from.MapKeys() {
				// convert source slice items and set to new slice
				var item reflect.Value

				switch to.Type().Elem().Kind() {
				// case when we have basic types
				case reflect.String, reflect.Bool, reflect.Int, reflect.Int64, reflect.Float64:
					item = from.MapIndex(key).Elem()
				//default:
				//	// case when we have pointer elements
				//	if to.Type().Elem().Kind() == reflect.Ptr {
				//		item = reflect.New(to.Type().Elem().Elem())
				//	} else {
				//		item = reflect.New(to.Type().Elem())
				//	}
				//
				//	item = convertVal(from.Index(i), item)
				}

				// if target slice is not slice of pointers
				//if item.Type().Kind() == reflect.Ptr && to.Type().Elem().Kind() != reflect.Ptr {
				//	item = item.Elem()
				//}

				m.SetMapIndex(key, item)
			}

			return m
		}
	}



	if to.Kind() == reflect.Ptr {
		if to.IsNil() {
			subStruct := reflect.New(to.Type().Elem()).Interface()
			TerraformToStruct(from.Interface(), subStruct)
			return reflect.ValueOf(subStruct)
		}

		TerraformToStruct(from.Interface(), to.Interface())
		return to
	}

	panic(fmt.Sprintf("can't convert value from %s, to %s", from.Type(), to.Type()))
}

func StructToTerraform(source interface{}) []interface{} {
	//log.Println("-------------------------")
	sType := reflect.TypeOf(source)
	if sType.Kind() != reflect.Ptr {
		panic("source is not pointer")
	}
	sType = sType.Elem()
	sVal := reflect.ValueOf(source)
	sVal = sVal.Elem()
	//log.Println("sType.Kind:", sType.Kind())

	result := make([]interface{}, 0)
	if sVal.Kind() == reflect.Slice {

		sType = sType.Elem()
		if sType.Kind() == reflect.Ptr {
			sType = sType.Elem()
		}

		for i := 0; i < sVal.Len(); i++ {
			s := sVal.Index(i)
			if s.Kind() == reflect.Ptr {
				s = s.Elem()
			}

			out := make(map[string]interface{})
			for i := 0; i < sType.NumField(); i++ {
				fType := sType.Field(i)
				fVal := s.Field(i)

				name := getTerraformFieldName(fType)
				if name == "" {
					// skip if dest field have no tag
					continue
				}
				//name = strings.Split(name, ",")[0]
				//name = toSnakeCase(name)

				//log.Println("field:", name)

				out[name] = convertStructVal(fVal)
			}
			result = append(result, out)
		}

	} else {

		out := make(map[string]interface{})
		for i := 0; i < sType.NumField(); i++ {
			fType := sType.Field(i)
			fVal := sVal.Field(i)

			name := getTerraformFieldName(fType)
			if name == "" {
				// skip if dest field have no tag
				continue
			}
			//fieldName = strings.Split(fieldName, ",")[0]
			//fieldName = toSnakeCase(fieldName)

			//log.Println("field:", fieldName)

			out[name] = convertStructVal(fVal)
		}
		result = append(result, out)
	}

	return result
}

func convertStructVal(val reflect.Value) interface{} {
	switch val.Kind() {
	case reflect.String:
		return val.String()
	case reflect.Int, reflect.Int64:
		return val.Int()
	case reflect.Float64:
		return val.Float()
	case reflect.Bool:
		return val.Bool()
	case reflect.Slice:
		//log.Println("val.Type:", val.Type())
		switch val.Type().Elem().Kind() {
		case reflect.String, reflect.Bool, reflect.Int, reflect.Int64, reflect.Float64:
			return val.Interface()
		}

		slice := make([]interface{}, val.Len(), val.Len())
		for i := 0; i < val.Len(); i++ {
			data := StructToTerraform(val.Index(i).Interface())
			slice[i] = data[0]
		}

		return slice

	case reflect.Map:
		switch val.Type().Elem().Kind() {
		case reflect.String, reflect.Bool, reflect.Int, reflect.Int64, reflect.Float64:
			return val.Interface()
		}
		return nil

	case reflect.Ptr:
		if val.IsNil() {
			return nil
		}

		return StructToTerraform(val.Interface())
	}
	return nil
}
