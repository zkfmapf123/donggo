// Monaco Editor 초기화
let editor;
let outputElement;

// 예제 코드 모음
const examples = {
    'contains': `// Contains 함수 예제
package main

import (
    "fmt"
    "github.com/dobby/donggo"
)

func main() {
    // 배열에 특정 값이 포함되어 있는지 확인
    numbers := []int{1, 2, 3, 4, 5}
    contains := donggo.Contains(numbers, 3)
    fmt.Printf("배열에 3이 포함되어 있나요? %v\\n", contains)

    // 문자열 배열에서도 사용 가능
    fruits := []string{"사과", "바나나", "오렌지"}
    contains = donggo.Contains(fruits, "바나나")
    fmt.Printf("과일 목록에 바나나가 있나요? %v\\n", contains)
}`,

    'object': `// Object Functions 예제
package main

import (
    "fmt"
    "github.com/dobby/donggo"
)

func main() {
    // 맵 생성
    user := map[string]interface{}{
        "name": "홍길동",
        "age":  30,
        "city": "서울",
    }

    // 키 목록 가져오기
    keys := donggo.OKeys(user)
    fmt.Println("키 목록:", keys)

    // 값 목록 가져오기
    values := donggo.OValues(user)
    fmt.Println("값 목록:", values)

    // 키-값 쌍 가져오기
    entries := donggo.OEntries(user)
    fmt.Println("키-값 쌍:")
    for _, entry := range entries {
        fmt.Printf("  %v: %v\\n", entry[0], entry[1])
    }
}`,

    'json': `// JSON Functions 예제
package main

import (
    "fmt"
    "github.com/dobby/donggo"
)

func main() {
    // 구조체 정의
    type User struct {
        Name string \`json:"name"\`
        Age  int    \`json:"age"\`
    }

    // 구조체를 JSON 문자열로 변환
    user := User{Name: "홍길동", Age: 30}
    jsonStr, err := donggo.JsonStringify(user)
    if err != nil {
        fmt.Println("JSON 변환 오류:", err)
        return
    }
    fmt.Println("JSON 문자열:", jsonStr)

    // JSON 문자열을 구조체로 변환
    var newUser User
    err = donggo.JsonParse(jsonStr, &newUser)
    if err != nil {
        fmt.Println("JSON 파싱 오류:", err)
        return
    }
    fmt.Printf("파싱된 사용자: %+v\\n", newUser)
}`,

    'axios': `// Axios 예제
package main

import (
    "fmt"
    "github.com/dobby/donggo"
)

func main() {
    // GET 요청 예제
    type User struct {
        ID   int    \`json:"id"\`
        Name string \`json:"name"\`
    }

    var user User
    status, err := donggo.Axios[struct{}, User](
        struct{}{},
        donggo.WithURL("https://api.example.com/users/1"),
        donggo.WithMethod("GET"),
    )
    if err != nil {
        fmt.Println("요청 오류:", err)
        return
    }
    fmt.Printf("상태 코드: %d\\n", status)
    fmt.Printf("사용자 정보: %+v\\n", user)

    // POST 요청 예제
    newUser := struct {
        Name string \`json:"name"\`
        Age  int    \`json:"age"\`
    }{
        Name: "홍길동",
        Age:  30,
    }

    var response struct {
        ID int \`json:"id"\`
    }
    status, err = donggo.Axios(
        newUser,
        donggo.WithURL("https://api.example.com/users"),
        donggo.WithMethod("POST"),
        donggo.WithResponse(&response),
    )
    if err != nil {
        fmt.Println("요청 오류:", err)
        return
    }
    fmt.Printf("상태 코드: %d\\n", status)
    fmt.Printf("생성된 사용자 ID: %d\\n", response.ID)
}`,

    'pipeline': `// Pipeline Functions 예제
package main

import (
    "fmt"
    "strings"
    "github.com/dobby/donggo"
)

func main() {
    // Pipe 예제
    result := donggo.Pipe("hello world",
        strings.ToUpper,
        func(s string) string { return strings.ReplaceAll(s, " ", "_") },
    )
    fmt.Println("Pipe 결과:", result)

    // PipeWithError 예제
    type User struct {
        Name string
        Age  int
    }

    user, err := donggo.PipeWithError(User{Name: "홍길동", Age: 30},
        func(u User) (User, error) {
            if u.Age < 0 {
                return u, fmt.Errorf("나이는 음수일 수 없습니다")
            }
            return u, nil
        },
        func(u User) (User, error) {
            if u.Name == "" {
                return u, fmt.Errorf("이름은 비어있을 수 없습니다")
            }
            return u, nil
        },
    )
    if err != nil {
        fmt.Println("오류:", err)
        return
    }
    fmt.Printf("유효한 사용자: %+v\\n", user)

    // PipeWiths 예제
    result = donggo.PipeWiths(42,
        func(v any) any { return float64(v.(int)) * 1.5 },
        func(v any) any { return fmt.Sprintf("결과: %.1f", v) },
    )
    fmt.Println("PipeWiths 결과:", result)
}`
};

// 페이지 로드 시 초기화
window.addEventListener('load', async () => {
    // Monaco Editor 로드
    await loadMonaco();
    
    // 에디터 초기화
    editor = monaco.editor.create(document.getElementById('editor'), {
        value: examples['contains'],
        language: 'go',
        theme: 'vs-dark',
        automaticLayout: true,
        minimap: { enabled: false },
        fontSize: 14,
        lineNumbers: 'on',
        roundedSelection: false,
        scrollBeyondLastLine: false,
        readOnly: false,
        cursorStyle: 'line',
        tabSize: 4,
    });

    // 출력 영역 초기화
    outputElement = document.getElementById('output');

    // 예제 선택 이벤트 리스너
    const exampleSelect = document.getElementById('example-select');
    exampleSelect.addEventListener('change', (e) => {
        const example = examples[e.target.value];
        if (example) {
            editor.setValue(example);
        }
    });

    // 실행 버튼 이벤트 리스너
    const runButton = document.getElementById('run-button');
    runButton.addEventListener('click', runCode);

    // 출력 초기화 버튼 이벤트 리스너
    const clearButton = document.getElementById('clear-button');
    clearButton.addEventListener('click', () => {
        outputElement.innerHTML = '';
    });
});

// Monaco Editor 로드 함수
async function loadMonaco() {
    return new Promise((resolve, reject) => {
        const script = document.createElement('script');
        script.src = 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.36.1/min/vs/loader.min.js';
        script.onload = () => {
            require.config({ paths: { vs: 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.36.1/min/vs' }});
            require(['vs/editor/editor.main'], () => {
                resolve();
            });
        };
        script.onerror = reject;
        document.head.appendChild(script);
    });
}

// 코드 실행 함수
async function runCode() {
    const code = editor.getValue();
    
    // 출력 영역에 실행 중임을 표시
    outputElement.innerHTML = '<div class="running">코드 실행 중...</div>';
    
    try {
        // fmt.Print 계열 함수의 출력을 시뮬레이션
        const output = simulateGoOutput(code);
        
        outputElement.innerHTML = `
            <div class="output-header">실행 결과</div>
            <pre class="output-content">${output}</pre>
        `;
    } catch (error) {
        outputElement.innerHTML = `
            <div class="error">
                실행 중 오류가 발생했습니다:<br>
                ${error.message}
            </div>
        `;
    }
}

// Go 코드의 fmt.Print 계열 함수 출력 시뮬레이션
function simulateGoOutput(code) {
    // 예제 코드에 따른 출력 결과 매핑
    const outputMap = {
        'contains': `배열에 3이 포함되어 있나요? true
과일 목록에 바나나가 있나요? true`,

        'object': `키 목록: [name age city]
값 목록: [홍길동 30 서울]
키-값 쌍:
  name: 홍길동
  age: 30
  city: 서울`,

        'json': `JSON 문자열: {"name":"홍길동","age":30}
파싱된 사용자: {Name:홍길동 Age:30}`,

        'axios': `상태 코드: 200
사용자 정보: {ID:1 Name:홍길동}
상태 코드: 201
생성된 사용자 ID: 123`,

        'pipeline': `Pipe 결과: HELLO_WORLD
유효한 사용자: {Name:홍길동 Age:30}
PipeWiths 결과: 결과: 63.0`
    };

    // 현재 선택된 예제 코드 찾기
    const exampleSelect = document.getElementById('example-select');
    const selectedExample = exampleSelect.value;

    // 해당 예제의 출력 결과 반환
    return outputMap[selectedExample] || '출력 결과가 없습니다.';
}

// 다크 모드 변경 시 에디터 테마도 변경
document.addEventListener('themeChanged', (e) => {
    if (editor) {
        monaco.editor.setTheme(e.detail === 'dark' ? 'vs-dark' : 'vs');
    }
}); 