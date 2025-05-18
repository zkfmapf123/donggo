package donggo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type axiosOptionParams struct {
	Url         string            `json:"url"`
	Method      string            `json:"method"`
	ContentType string            `json:"contentType"`
	Headers     map[string]string `json:"headers"`
}

type axiosOptions func(option *axiosOptionParams) error

func defultAxiosOptions() *axiosOptionParams {
	return &axiosOptionParams{
		Url:         "",
		Method:      "GET",
		ContentType: "application/json",
		Headers:     map[string]string{},
	}
}

func WithUrl(url string) axiosOptions {
	return func(option *axiosOptionParams) error {
		option.Url = url
		return nil
	}
}

func WithMethod(method string) axiosOptions {
	return func(option *axiosOptionParams) error {
		option.Method = method
		return nil
	}
}

func WithContentType(contentType string) axiosOptions {
	return func(option *axiosOptionParams) error {
		option.ContentType = contentType
		return nil
	}
}

func WithHeaders(headers map[string]string) axiosOptions {
	return func(option *axiosOptionParams) error {
		option.Headers = headers
		return nil
	}
}

func Axios[T, Res any](body T, options ...axiosOptions) (Res, int, error) {
	var emptyRes Res
	params := defultAxiosOptions()

	for _, option := range options {
		if err := option(params); err != nil {
			return emptyRes, 0, fmt.Errorf("옵션 설정 실패: %w", err)
		}
	}

	var req *http.Request
	var resp *http.Response
	var err error

	switch params.Method {
	case "GET":
		req, err = http.NewRequest("GET", params.Url, nil)
		if err != nil {
			return emptyRes, 0, fmt.Errorf("GET 요청 생성 실패: %w", err)
		}

	case "POST", "PUT", "PATCH":
		data, err := json.Marshal(body)
		if err != nil {
			var zero Res
			return zero, 0, fmt.Errorf("JSON 변환 실패: %w", err)
		}

		req, err = http.NewRequest(params.Method, params.Url, bytes.NewBuffer(data))
		if err != nil {
			return emptyRes, 0, fmt.Errorf("%s 요청 생성 실패: %w", params.Method, err)
		}

	default:
		return emptyRes, 0, fmt.Errorf("지원하지 않는 HTTP 메소드: %s", params.Method)
	}

	// 헤더 설정
	for k, v := range params.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		return emptyRes, resp.StatusCode, fmt.Errorf("HTTP 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyRes, resp.StatusCode, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	jsonData := JsonParse[Res](b)
	return jsonData, resp.StatusCode, nil
}
