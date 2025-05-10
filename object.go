package donggo

/*
Title: Okeys
Desc: Javascript (Object.keys) 와 동일
*/
func OKeys[T any](values map[string]T) []string {

	res := []string{}
	for k := range values {
		res = append(res, k)
	}

	return res
}

/*
Title: Okeys
Desc: Javascript (Object.values) 와 동일
*/
func OValues[T any](values map[string]T) []T {

	res := []T{}
	for _, v := range values {
		res = append(res, v)
	}

	return res
}

/*
Title: Okeys
Desc: Javascript (Object.entries) 와 동일
*/
func OEntries[T any](values map[string]T) [][]any {

	entiresArr := [][]any{}

	for k, v := range values {

		arr := []any{k, v}

		entiresArr = append(entiresArr, arr)
	}

	return entiresArr

}
