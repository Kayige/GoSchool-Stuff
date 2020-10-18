package postquery

import (
	"fmt"
	"net/http"
	"strconv"
)

func FormValues(r *http.Request, keys []string) (data map[string]string, err error) {
	data = make(map[string]string)
	for _, k := range keys {
		var v string
		v, err = GetFirstValueByKey(r, k)
		if err != nil {
			return
		}
		data[k] = v
	}
	return
}

func RequiredFormParamString(r *http.Request, name string) (string, error) {
	value, err := GetFirstValueByKey(r, name)
	if err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("%v is required", name)
	}
	return value, nil
}

func OptionalFormParamString(r *http.Request, name string) (string, error) {
	value, err := GetFirstValueByKey(r, name)
	if err != nil {
		return "", err
	}
	return value, nil
}

func RequiredFormParamUint(r *http.Request, name string) (uint64, error) {
	value, err := GetFirstValueByKey(r, name)
	if err != nil {
		return 0, err
	}
	if value == "" {
		return 0, fmt.Errorf("%v is required", name)
	}
	valueUint, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%v is invalid", name)
	}
	return valueUint, nil
}

func OptionalFormParamUint(r *http.Request, name string) (uint64, error) {
	value, err := GetFirstValueByKey(r, name)
	if err != nil {
		return 0, err
	}
	if value == "" {
		return 0, nil
	}
	valueUint, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%v is invalid", name)
	}
	return valueUint, nil
}
