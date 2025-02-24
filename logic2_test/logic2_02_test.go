package logic2_testing

import (
	"fmt"
	"testing"
	"github.com/jejejery/go-logic/logic2"
	"github.com/stretchr/testify/assert"
)

func TestLogic2_02(t *testing.T) {
	result := logic2.Logic2_2(9)
	expected := [][]int{
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
	}

	assert.Equal(t, result, expected)
	fmt.Println("TestLogic2_02 done")
}