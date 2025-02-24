package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_07(t *testing.T) {
	result1 := logic1.Logic1_7(10)
	result2 := logic1.Logic1_7(11)
	expected1 := []int{1, 3, 5, 7, 9, 9, 7, 5, 3, 1}
	expected2 := []int{1, 3, 5, 7, 9, 11, 9, 7, 5, 3, 1}

	assert.Equal(t, result1, expected1)
	assert.Equal(t, result2, expected2)
	fmt.Println("TestLogic1_07 done")
}