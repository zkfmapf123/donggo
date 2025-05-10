package donggo

import "encoding/json"

/*
Title: JsonStringify
Desc: Javascript (JSON.stringify) 와 동일
*/
func JsonStringify[T any](v T) ([]byte, string, error) {

	b, err := json.Marshal(v)
	if err != nil {
		return nil, "", err
	}

	return b, string(b), nil
}

/*
Title: JsonParse
Desc: Javascript (JSON.parse) 와 동일
*/
func JsonParse[T any](v []byte) T {

	var format T
	json.Unmarshal(v, &format)
	return format
}
