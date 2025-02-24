package logic2_testing

import (
	"fmt"
	"testing"
	"github.com/jejejery/go-logic/logic2"
	"github.com/stretchr/testify/assert"
)

func TestLogic2_04(t *testing.T) {
	result := logic2.Logic2_4(9)
	expected := [][]int{
		{1, 4, 7, 10, 13, 16, 19, 22, 25},
		{28, 31, 34, 37, 40, 43, 46, 49, 52},
		{55, 58, 61, 64, 67, 70, 73, 76, 79},
		{82, 85, 88, 91, 94, 97, 100, 103, 106},
		{109, 112, 115, 118, 121, 124, 127, 130, 133},
		{136, 139, 142, 145, 148, 151, 154, 157, 160},
		{163, 166, 169, 172, 175, 178, 181, 184, 187},
		{190, 193, 196, 199, 202, 205, 208, 211, 214},
		{217, 220, 223, 226, 229, 232, 235, 238, 241},
	}

	assert.Equal(t, result, expected)
	fmt.Println("TestLogic2_02 done")
}