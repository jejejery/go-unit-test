package logic2_testing

import (
	"fmt"
	"testing"
	"github.com/jejejery/go-logic/logic2"
	"github.com/stretchr/testify/assert"
)

func TestLogic2_12(t *testing.T) {
	result := logic2.Logic2_12(9)
	expected := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 17},
		{1, 3, 0, 0, 0, 0, 0, 15, 17},
		{1, 3, 5, 0, 0, 0, 13, 15, 17},
		{1, 3, 5, 7, 0, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 0, 11, 13, 15, 17},
		{1, 3, 5, 0, 0, 0, 13, 15, 17},
		{1, 3, 0, 0, 0, 0, 0, 15, 17},
		{1, 0, 0, 0, 0, 0, 0, 0, 17},
	}

	assert.Equal(t, result, expected)
	fmt.Println("TestLogic2_01 done")
}