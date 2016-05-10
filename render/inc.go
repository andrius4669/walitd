package render

import (
	"errors"
	"reflect"
)

var (
	errBadType = errors.New("bad type")
)

func intCast(v reflect.Value) (int64, error) {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(v.Uint()), nil
	}
	return 0, errBadType
}

func inc(args ...interface{}) (int64, error) {
	v := reflect.ValueOf(args[0])
	rv, err := intCast(v)
	if err != nil {
		return 0, err
	}
	return rv + 1, nil
}
