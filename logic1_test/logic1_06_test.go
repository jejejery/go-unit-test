package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_06(t *testing.T) {
	result := logic1.Logic1_6(10)
	expected := []int{30, 27, 24, 21, 18, 15, 12, 9, 6, 3}
	not_expected := []int{30, 27, 24, 21, 18, 15, 12, 9, 6, 4}
	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_06 done")
}