package injector

import (
	"fmt"
	"reflect"
	"strings"
)

// Inject value inside a struct based on struct tags
func Inject(target interface{}, tagName string, data map[string]interface{}) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() || targetValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	targetValue = targetValue.Elem()
	targetType := targetValue.Type()

	for i := 0; i < targetValue.NumField(); i++ {
		field := targetType.Field(i)
		fieldValue := targetValue.Field(i)
		fieldType := fieldValue.Type()

		tagValue, tagExist := field.Tag.Lookup(tagName)
		if !tagExist {
			continue
		}

		// Handle json tag format (e.g. json:"name,omitempty")
		tagValue = strings.SplitN(tagValue, ",", 2)[0]

		dataInterface, dataExist := data[tagValue]
		if !dataExist {
			continue
		}

		dataValue := reflect.ValueOf(dataInterface)
		dataType := dataValue.Type()
		if !dataType.AssignableTo(fieldType) {
			return fmt.Errorf("trying to assign incompatible type %s to %s for tag %s", dataValue.Type().Name(), fieldType.Name(), tagValue)
		}

		fieldValue.Set(dataValue)
	}

	return nil
}
