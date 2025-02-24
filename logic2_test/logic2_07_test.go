package logic2_testing

import (
	"fmt"
	"testing"
	"github.com/jejejery/go-logic/logic2"
	"github.com/stretchr/testify/assert"
)

func TestLogic2_07(t *testing.T) {
	result := logic2.Logic2_7(9)
	expected := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 5, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 7, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 9, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 11, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 13, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 15, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 17},
	}

	assert.Equal(t, result, expected)
	fmt.Println("TestLogic2_02 done")
}