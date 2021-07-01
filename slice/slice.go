package slice

import (
	"helper/utils"
	"reflect"
)

func IsContained(container interface{}, item interface{}) (res bool, err error) {
	defer func() {
		if pac := recover(); pac != nil {
			err = pac.(error)
			res = false
			return
		}
	}()
	containerType := reflect.TypeOf(container)
	utils.Assert1(containerType.Kind() == reflect.Slice, "container expected a slice")

	containerElemType := containerType.Elem()
	itemType := reflect.TypeOf(item)
	if containerElemType != itemType {
		return false, nil
	}

	containerValue := reflect.ValueOf(container)
	for i := 0; i < containerValue.Len(); i++ {
		elemInterface := containerValue.Index(i).Interface()
		if elemInterface == item {
			return true, nil
		}
	}
	return false, nil
}
