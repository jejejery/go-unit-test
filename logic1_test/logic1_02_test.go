package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_02(t *testing.T) {
	result := logic1.Logic1_2(10)
	expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	not_expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 21}
	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_02 done")
}