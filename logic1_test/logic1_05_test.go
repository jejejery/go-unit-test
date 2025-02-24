package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_05(t *testing.T) {
	result := logic1.Logic1_5(10)
	expected := []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 2}
	not_expected := []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 3}
	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_05 done")
}