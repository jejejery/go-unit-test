package logic1_testing

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_08(t *testing.T) {
	result1 := logic1.Logic1_8(10)
	expected1 := []int{2, 4, 6, 8, 10, 10, 8, 6, 4, 2}
	result2 := logic1.Logic1_8(11)
	expected2 := []int{2, 4, 6, 8, 10, 12, 10, 8, 6, 4, 2}

	assert.Equal(t, result1, expected1)
	assert.Equal(t, result2, expected2)
	fmt.Println("TestLogic1_08 done")

}
