package logic2_testing

import (
	"fmt"
	"testing"
	"github.com/jejejery/go-logic/logic2"
	"github.com/stretchr/testify/assert"
)

func TestLogic2_11(t *testing.T) {
	result := logic2.Logic2_11(9)
	expected := [][]int{
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{0, 3, 5, 7, 9, 11, 13, 15, 17},
		{0, 0, 5, 7, 9, 11, 13, 15, 17},
		{0, 0, 0, 7, 9, 11, 13, 15, 17},
		{0, 0, 0, 0, 9, 11, 13, 15, 17},
		{0, 0, 0, 0, 0, 11, 13, 15, 17},
		{0, 0, 0, 0, 0, 0, 13, 15, 17},
		{0, 0, 0, 0, 0, 0, 0, 15, 17},
		{0, 0, 0, 0, 0, 0, 0, 0, 17},
	}

	assert.Equal(t, result, expected)
	fmt.Println("TestLogic2_01 done")
}