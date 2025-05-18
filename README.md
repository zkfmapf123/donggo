# Donggo Golang Packages

<h1> <a href="https://zkfmapf123.github.io/donggo/"> donggo homepage </h1>

Go 언어로 구현된 유틸리티 패키지입니다. JavaScript 스타일의 함수들을 Go 언어에서 사용할 수 있도록 구현했습니다.

## Install

```sh
go get github.com/zkfmapf123/donggo
```

## Features

### Array Functions

#### Contains[T any](arr []T, target T) bool
배열에 특정 값이 포함되어 있는지 검사하는 함수입니다.
```go
arr := []int{1, 2, 3, 4, 5}
result := Contains(arr, 3) // true
```

### Object Functions

#### OKeys[T any](values map[string]T) []string
JavaScript의 `Object.keys`와 동일한 기능을 수행합니다. 맵의 모든 키를 배열로 반환합니다.
```go
m := map[string]int{"a": 1, "b": 2}
keys := OKeys(m) // []string{"a", "b"}
```

#### OValues[T any](values map[string]T) []T
JavaScript의 `Object.values`와 동일한 기능을 수행합니다. 맵의 모든 값을 배열로 반환합니다.
```go
m := map[string]int{"a": 1, "b": 2}
values := OValues(m) // []int{1, 2}
```

#### OEntries[T any](values map[string]T) [][]any
JavaScript의 `Object.entries`와 동일한 기능을 수행합니다. 맵의 키-값 쌍을 2차원 배열로 반환합니다.
```go
m := map[string]int{"a": 1, "b": 2}
entries := OEntries(m) // [][]any{{"a", 1}, {"b", 2}}
```

### JSON Functions

#### JsonStringify[T any](v T) ([]byte, string, error)
JavaScript의 `JSON.stringify`와 동일한 기능을 수행합니다. 객체를 JSON 문자열로 변환합니다.
```go
data := map[string]interface{}{"name": "John", "age": 30}
bytes, str, err := JsonStringify(data)
// str: `{"name":"John","age":30}`
```

#### JsonParse[T any](v []byte) T
JavaScript의 `JSON.parse`와 동일한 기능을 수행합니다. JSON 문자열을 객체로 변환합니다.
```go
jsonStr := []byte(`{"name":"John","age":30}`)
var result map[string]interface{}
result = JsonParse[map[string]interface{}](jsonStr)
```

### HTTP Client (Axios)

JavaScript의 Axios와 유사한 HTTP 클라이언트를 제공합니다.

#### Axios[T, Res any](body T, options ...axiosOptions) (Res, int, error)
HTTP 요청을 보내고 응답을 받는 함수입니다. 제네릭을 사용하여 요청과 응답 타입을 지정할 수 있습니다.

옵션 설정 함수들:
- `WithUrl(url string)`: 요청 URL 설정
- `WithMethod(method string)`: HTTP 메소드 설정 (GET, POST, PUT, PATCH)
- `WithContentType(contentType string)`: Content-Type 헤더 설정
- `WithHeaders(headers map[string]string)`: 커스텀 헤더 설정

```go
// GET 요청 예제
response, statusCode, err := Axios[struct{}, PostResponse](
    struct{}{},
    WithUrl("https://api.example.com/posts/1"),
    WithMethod("GET"),
)

// POST 요청 예제
post := Post{Title: "제목", Body: "내용", UserID: 1}
response, statusCode, err := Axios[Post, PostResponse](
    post,
    WithUrl("https://api.example.com/posts"),
    WithMethod("POST"),
    WithHeaders(map[string]string{
        "Content-Type": "application/json",
    }),
)
```

### Pipeline Functions

함수형 프로그래밍을 위한 파이프라인 패턴을 제공합니다.

#### Pipe[T any](v T, fns ...pipeFunc[T]) T
일반적인 파이프라인 함수입니다. 동일한 타입의 함수들을 연결하여 실행합니다.
```go
result := Pipe(3,
    func(v int) int { return v * 2 },
    func(v int) int { return v + 1 },
) // 7
```

#### PipeWithError[T any](v T, fns ...PipeWithErrorFunc[T]) (T, error)
에러 처리가 가능한 파이프라인 함수입니다. 각 단계에서 에러가 발생하면 파이프라인이 중단됩니다.
```go
result, err := PipeWithError(3,
    func(v int) (int, error) { return v * 2, nil },
    func(v int) (int, error) { 
        if v > 5 {
            return v, errors.New("value is over 5")
        }
        return v + 1, nil 
    },
)
```

#### PipeWiths(v any, fns ...func(any) any) any
타입이 다른 함수들을 연결할 수 있는 파이프라인 함수입니다.
```go
result := PipeWiths(3,
    func(v any) any { return v.(int) * 10 },
    func(v any) any { return strconv.Itoa(v.(int)) },
) // "30"
```

## License

MIT License
