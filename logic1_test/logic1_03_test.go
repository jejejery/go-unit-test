package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_03(t *testing.T) {
	result := logic1.Logic1_3(10)
	expected := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}
	not_expected := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 31}
	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_03 done")
}