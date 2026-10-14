package highway

import "reflect"

func GetReflectType(obj any) string{
	return reflect.TypeOf(obj).String()
}