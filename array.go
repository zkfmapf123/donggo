package donggo

import "reflect"

/*
Title: Containts
Desc: array 에 포함되어있는지 검사하는 함수
*/
func Contains[T any](arr []T, target T) bool {

	for _, v := range arr {

		if reflect.DeepEqual(v, target) {
			return true
		}
	}

	return false
}
