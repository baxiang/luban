package container

import "reflect"

func IsArray(value interface{}) bool {
	rv := reflect.ValueOf(value)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

// RemoveStringSliceItemByIndex 根据下标删除 string slice 中的元素
func RemoveStringSliceItemByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// SliceCopy does a shallow copy of slice <data> for most commonly used slice type
// []interface{}.
func SliceCopy(data []interface{}) []interface{} {
	newData := make([]interface{}, len(data))
	copy(newData, data)
	return newData
}

// SliceDelete deletes an element at <index> and returns the new slice.
// It does nothing if the given <index> is invalid.
func SliceDelete(data []interface{}, index int) (newSlice []interface{}) {
	if index < 0 || index >= len(data) {
		return data
	}
	// Determine array boundaries when deleting to improve deletion efficiency.
	if index == 0 {
		return data[1:]
	} else if index == len(data)-1 {
		return data[:index]
	}
	// If it is a non-boundary delete,
	// it will involve the creation of an array,
	// then the deletion is less efficient.
	return append(data[:index], data[index+1:]...)
}
