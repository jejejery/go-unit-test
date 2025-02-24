package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_01(t *testing.T) {
	result := logic1.Logic1_1(10)
	expected := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	not_expected := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 21}
	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_01 done")
}

