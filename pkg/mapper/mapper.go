package mapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func MapStruct(from, to interface{}) {
	fromValue := reflect.ValueOf(from)
	if fromValue.Kind() != reflect.Struct &&
		(fromValue.Kind() != reflect.Ptr && fromValue.Elem().Kind() != reflect.Struct) {
		panic(errors.New("from is not struct or pointer of struct"))
	}
	fromType := fromValue.Type()
	if fromType.Kind() == reflect.Ptr {
		fromType = fromType.Elem()
		fromValue = fromValue.Elem()
	}

	toValue := reflect.ValueOf(to)
	if toValue.Kind() != reflect.Ptr || toValue.Elem().Kind() != reflect.Struct {
		panic(errors.New("to is not pointer of struct"))
	}
	toType := toValue.Elem().Type()
	toValue = toValue.Elem()

	for i := 0; i < fromType.NumField(); i++ {
		fromTypeField := fromType.Field(i)
		fromFieldTags := getTagNames(fromTypeField)
		fromField := fromValue.FieldByName(fromTypeField.Name)
		fromFieldType := fromTypeField.Type
		if fromFieldType.Kind() == reflect.Ptr {
			fromFieldType = fromFieldType.Elem()
		}
		if fromTypeField.Anonymous {
			MapStruct(fromField.Interface(), to)
		}

		for i := 0; i < toType.NumField(); i++ {
			toTypeField := toType.Field(i)
			toFieldTags := getTagNames(toTypeField)
			toField := toValue.FieldByName(toTypeField.Name)
			toFieldType := toTypeField.Type
			if toFieldType.Kind() == reflect.Ptr {
				toFieldType = toFieldType.Elem()
			}
			if !toField.CanSet() {
				continue
			}

			if compareTags(fromFieldTags, toFieldTags) {
				if fromField.IsZero() {
					toField.Set(reflect.Zero(toField.Type()))
					continue
				}

				if fromFieldType.Kind() == toFieldType.Kind() && toFieldType.Kind() == reflect.Struct {
					val := reflect.New(toFieldType)
					if val.Kind() != reflect.Ptr {
						val = val.Addr()
					}
					MapStruct(fromField.Interface(), val.Interface())
					if toField.Kind() == reflect.Ptr {
						toField.Set(val)
					} else {
						toField.Set(val.Elem())
					}
					continue
				}

				if fromFieldType.Kind() == toFieldType.Kind() && toFieldType.Kind() == reflect.Slice {
					MapSlice(fromField.Interface(), toField.Addr().Interface())
					continue
					//fromFieldTypeElem := fromFieldType.Elem()
					//if fromFieldTypeElem.Kind() == reflect.Ptr {
					//	fromFieldTypeElem = fromFieldTypeElem.Elem()
					//}
					//toFieldTypeElem := toFieldType.Elem()
					//if toFieldTypeElem.Kind() == reflect.Ptr {
					//	toFieldTypeElem = toFieldTypeElem.Elem()
					//}
					//if fromFieldTypeElem.Kind() == toFieldTypeElem.Kind() && toFieldTypeElem.Kind() == reflect.Struct {
					//	slice := reflect.MakeSlice(toFieldType, 0, fromField.Len())
					//	for i := 0; i < fromField.Len(); i++ {
					//		item := reflect.New(toFieldTypeElem)
					//		MapStruct(fromField.Index(i).Interface(), item.Interface())
					//		slice = reflect.Append(slice, item)
					//	}
					//	toField.Set(slice)
					//	continue
					//}
				}

				if fromFieldType.Kind() == toFieldType.Kind() && toFieldType.Kind() == reflect.Map {
					fromFieldTypeElem := fromFieldType.Elem()
					if fromFieldTypeElem.Kind() == reflect.Ptr {
						fromFieldTypeElem = fromFieldTypeElem.Elem()
					}
					toFieldTypeElem := toFieldType.Elem()
					if toFieldTypeElem.Kind() == reflect.Ptr {
						toFieldTypeElem = toFieldTypeElem.Elem()
					}
					if fromFieldTypeElem.Kind() == toFieldTypeElem.Kind() && toFieldTypeElem.Kind() == reflect.Struct {
						dic := reflect.MakeMap(toFieldType)
						for _, key := range fromField.MapKeys() {
							item := reflect.New(toFieldTypeElem)
							MapStruct(fromField.MapIndex(key).Interface(), item.Interface())
							dic.SetMapIndex(key, item)
						}
						toField.Set(dic)
						continue
					}
				}

				if fromField.Type() == toField.Type() {
					toField.Set(fromField)
				} else {
					switch val := fromField.Interface().(type) {
					case time.Time:
						if toField.Kind() == reflect.Int64 {
							toField.Set(reflect.ValueOf(val.Unix()))
							continue
						}
					}
					if fromField.Kind() == reflect.Int64 {
						switch toField.Interface().(type) {
						case time.Time:
							toField.Set(reflect.ValueOf(time.Unix(fromField.Int(), 0)))
							continue
						}
					}
					bin, _ := json.Marshal(fromField.Interface())
					_ = json.Unmarshal(bin, toField.Addr().Interface())
				}
			}
		}
	}
}

func MapSlice(from, to interface{}) {
	fromValue := reflect.ValueOf(from)
	if fromValue.Kind() != reflect.Slice &&
		(fromValue.Kind() != reflect.Ptr && fromValue.Elem().Kind() != reflect.Slice) {
		panic(errors.New("from is not slice or pointer of slice"))
	}
	fromType := fromValue.Type()
	if fromType.Kind() == reflect.Ptr {
		fromType = fromType.Elem()
		fromValue = fromValue.Elem()
	}

	toValue := reflect.ValueOf(to)
	if toValue.Kind() != reflect.Ptr || toValue.Elem().Kind() != reflect.Slice {
		panic(errors.New("to is not pointer of slice"))
	}
	toType := toValue.Elem().Type()
	toValue = toValue.Elem()

	fromTypeElem := toType.Elem()
	if fromTypeElem.Kind() == reflect.Ptr {
		fromTypeElem = fromTypeElem.Elem()
	}
	toTypeElem := toType.Elem()
	if toTypeElem.Kind() == reflect.Ptr {
		toTypeElem = toTypeElem.Elem()
	}
	toValue.Set(reflect.MakeSlice(toType, 0, toValue.Cap()))

	if fromTypeElem.Kind() == toTypeElem.Kind() && toTypeElem.Kind() == reflect.Struct {
		for i := 0; i < fromValue.Len(); i++ {
			fromItem := fromValue.Index(i)
			isExists := false
			if fromID, ok := getIDField(fromValue.Index(i)); ok {
				for j := 0; j < toValue.Len(); j++ {
					toItem := toValue.Index(j)
					if toID, ok := getIDField(toItem); ok {
						if fmt.Sprint(fromID.Interface()) == fmt.Sprint(toID.Interface()) {
							MapStruct(fromItem.Interface(), toItem.Interface())
							isExists = true
							break
						}
					}
				}
			}
			if !isExists {
				toItem := reflect.New(toTypeElem)
				MapStruct(fromItem.Interface(), toItem.Interface())
				toValue.Set(reflect.Append(toValue, toItem))
			}
		}
		return
	}

	if fromTypeElem.Kind() == toTypeElem.Kind() && toTypeElem.Kind() == reflect.Slice {
		for i := 0; i < fromValue.Len(); i++ {
			item := reflect.New(toTypeElem)
			MapSlice(fromValue.Index(i).Interface(), item.Interface())
			toValue.Set(reflect.Append(toValue, item.Elem()))
		}
		return
	}

	bin, _ := json.Marshal(fromValue.Interface())
	_ = json.Unmarshal(bin, toValue.Addr().Interface())
}

func getIDField(value reflect.Value) (reflect.Value, bool) {
	if value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}
	t := value.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			if v, ok := getIDField(value.FieldByName(field.Name)); ok {
				return v, true
			}
		}
		if strings.ToLower(field.Name) == "id" {
			return value.FieldByName(field.Name), true
		}
	}
	return reflect.ValueOf(nil), false
}
