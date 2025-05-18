package donggo

import "fmt"

type pipeFunc[T any] func(T) T
type PipeWithErrorFunc[T any] func(T) (T, error)

/*
Title: Pipe
Desc: 함수를 연결하여 실행 (실행결과 return)
*/
func Pipe[T any](v T, fns ...pipeFunc[T]) T {
	result := v
	for _, fn := range fns {
		result = fn(result)
	}
	return result
}

/*
Title: PipeWithError
Desc: 함수를 연결하여 실행 (실행결과 return) + Error 리턴
*/
func PipeWithError[T any](v T, fns ...PipeWithErrorFunc[T]) (T, error) {
	result := v
	var err error

	for _, fn := range fns {
		result, err = fn(result)

		fmt.Println(">>>", result, err)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

/*
Title: PipeWiths
Desc: 함수를 연결하여 실행 (실행결과 return) + any 타입 리턴
*/
func PipeWiths(v any, fns ...func(any) any) any {
	result := v
	for _, fn := range fns {
		result = fn(result)
	}
	return result
}
