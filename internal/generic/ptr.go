package generic

import "reflect"

func indirectType(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Pointer {
		return typ.Elem()
	}

	return typ
}
