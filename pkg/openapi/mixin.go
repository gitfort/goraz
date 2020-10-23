package openapi

import (
	"reflect"
)

func Mixin(prim *Swagger, mix *Swagger) *Swagger {
	primRef := reflect.ValueOf(prim)
	mixRef := reflect.ValueOf(mix)
	recursive(primRef, mixRef)
	return primRef.Interface().(*Swagger)
}

func recursive(primRef reflect.Value, mixRef reflect.Value) {

	if !mixRef.IsValid() || mixRef.IsZero() {
		return
	} else if primRef.IsZero() {
		primRef.Set(mixRef)
		return
	}

	if primRef.Kind() == reflect.Ptr {
		primRef = primRef.Elem()
	}
	if mixRef.Kind() == reflect.Ptr {
		mixRef = mixRef.Elem()
	}

	switch mixRef.Kind() {
	case reflect.Struct:
		for i := 0; i < mixRef.NumField(); i++ {
			mixField := mixRef.Field(i)
			primField := primRef.Field(i)
			//if primField.IsZero() {
			//	primField.Set()
			//}
			//log.Print(mixField, primField)
			recursive(primField, mixField)
		}
	case reflect.Map:
		for _, key := range mixRef.MapKeys() {
			mixValue := mixRef.MapIndex(key)
			primValue := primRef.MapIndex(key)
			if !primValue.IsValid() {
				primRef.SetMapIndex(key, mixValue)
			} else {
				recursive(primValue, mixValue)
			}
		}
	}
}
