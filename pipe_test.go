package donggo

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pipe(t *testing.T) {

	v := 3

	res := Pipe(v, func(v int) int {
		return v * 10
	}, func(v int) int {
		return v + 10
	})

	assert.Equal(t, res, 40)
}

func Test_PipeWiths(t *testing.T) {

	v := 3

	res := PipeWiths(v,
		func(v any) any { return v.(int) * 10 },
		func(v any) any { return v.(int) * 10 },
		func(v any) any { return strconv.Itoa(v.(int)) },
	)

	assert.Equal(t, res, "300")
}

func Test_PipeWithError(t *testing.T) {
	v := 3

	res, err := PipeWithError(v,
		func(v int) (int, error) {
			return v * 10, nil
		}, func(v int) (int, error) {
			if v > 5 {
				return v, errors.New("value is over 5")
			}

			return v + 10, nil
		})

	fmt.Println(res, err)

	assert.Error(t, err, "value is over 5")

}
