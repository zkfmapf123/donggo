package donggo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type PostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func TestAxios_GET(t *testing.T) {
	response, statusCode, err := Axios[struct{}, PostResponse](
		struct{}{},
		WithUrl("https://jsonplaceholder.typicode.com/posts/1"),
		WithMethod("GET"),
	)

	if err != nil {
		t.Errorf("GET 요청 실패: %v", err)
	}

	if statusCode != 200 {
		t.Errorf("예상 상태 코드 200, 실제: %d", statusCode)
	}

	if response.ID != 1 {
		t.Errorf("예상 ID 1, 실제: %d", response.ID)
	}
}

func TestAxios_POST(t *testing.T) {
	// POST 요청 테스트
	post := Post{
		Title:  "테스트 제목",
		Body:   "테스트 내용",
		UserID: 1,
	}

	response, statusCode, err := Axios[Post, PostResponse](
		post,
		WithUrl("https://jsonplaceholder.typicode.com/posts"),
		WithMethod("POST"),
		WithHeaders(map[string]string{
			"Content-Type": "application/json",
		}),
	)

	if err != nil {
		t.Errorf("POST 요청 실패: %v", err)
	}

	if statusCode != 201 {
		t.Errorf("예상 상태 코드 201, 실제: %d", statusCode)
	}

	if response.Title != post.Title {
		t.Errorf("예상 제목 %s, 실제: %s", post.Title, response.Title)
	}

	if response.Body != post.Body {
		t.Errorf("예상 내용 %s, 실제: %s", post.Body, response.Body)
	}

	if response.UserID != post.UserID {
		t.Errorf("예상 UserID %d, 실제: %d", post.UserID, response.UserID)
	}
}

func TestAxios_Error(t *testing.T) {
	// 잘못된 URL로 요청하여 에러 테스트
	_, statusCode, _ := Axios[struct{}, PostResponse](
		struct{}{},
		WithUrl("https://jsonplaceholder.typicode.com/invalid"),
		WithMethod("GET"),
	)

	assert.Equal(t, statusCode, 404)

}

func TestAxios_InvalidMethod(t *testing.T) {
	// 잘못된 HTTP 메소드 테스트
	_, statusCode, err := Axios[struct{}, PostResponse](
		struct{}{},
		WithUrl("https://jsonplaceholder.typicode.com/posts/1"),
		WithMethod("INVALID"),
	)

	if err == nil {
		t.Error("잘못된 메소드에 대한 에러가 발생해야 하는데 발생하지 않았습니다")
	}

	if statusCode != 0 {
		t.Errorf("잘못된 메소드의 경우 상태 코드는 0이어야 하는데, 실제: %d", statusCode)
	}
}
