package logic2_testing

import (
	"fmt"
	"testing"
	"github.com/jejejery/go-logic/logic2"
	"github.com/stretchr/testify/assert"
)

func TestLogic2_03(t *testing.T) {
	result := logic2.Logic2_3(9)
	expected := [][]int{
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{19, 21, 23, 25, 27, 29, 31, 33, 35},
		{37, 39, 41, 43, 45, 47, 49, 51, 53},
		{55, 57, 59, 61, 63, 65, 67, 69, 71},
		{73, 75, 77, 79, 81, 83, 85, 87, 89},
		{91, 93, 95, 97, 99, 101, 103, 105, 107},
		{109, 111, 113, 115, 117, 119, 121, 123, 125},
		{127, 129, 131, 133, 135, 137, 139, 141, 143},
		{145, 147, 149, 151, 153, 155, 157, 159, 161},
	}

	assert.Equal(t, result, expected)
	fmt.Println("TestLogic2_02 done")
}