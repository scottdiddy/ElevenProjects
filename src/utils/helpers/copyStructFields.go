package helpers

import "reflect"

func CopyStructFields(original interface{}, incoming interface{}) {
	originalValue := reflect.ValueOf(original).Elem()
	originalType := originalValue.Type()
	incomingValue := reflect.ValueOf(incoming).Elem()

	for i := 0; i < originalValue.NumField(); i++ {
		originalField := originalValue.Field(i)
		fieldName := originalType.Field(i).Name
		incomingField := incomingValue.FieldByName(fieldName)

		if originalField.Kind() == reflect.Struct && incomingField.Kind() == reflect.Struct {
			CopyStructFields(originalField.Addr().Interface(), incomingField.Addr().Interface())
		} else if !isEmpty(incomingField) && incomingField.IsValid() {
			originalField.Set(incomingField)
		}

	}
}
func isEmpty(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.String:
		return field.String() == ""
	case reflect.Ptr, reflect.Interface:
		return field.IsNil()
	default:
		return false
	}
}
