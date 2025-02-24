package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_04(t *testing.T) {
	result := logic1.Logic1_4(10)
	expected := []int{19, 17 ,15, 13, 11, 9, 7, 5, 3, 1}
	not_expected := []int{19, 17 ,15, 13, 11, 9, 7, 5, 3, 2}
	if result == nil {
		return
	}
	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_04 done")
}