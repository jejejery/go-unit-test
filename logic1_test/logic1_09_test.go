package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_09(t *testing.T) {
	result1 := logic1.Logic1_9(10)
	expected1 := []int{3, 6, 9, 12, 15, 15, 12, 9, 6, 3}
	result2 := logic1.Logic1_9(11)
	expected2 := []int{3, 6, 9, 12, 15, 18, 15, 12, 9, 6, 3}

	assert.Equal(t, result1, expected1)
	assert.Equal(t, result2, expected2)
	fmt.Println("TestLogic1_09 done")
}