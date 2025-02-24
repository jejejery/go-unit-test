package logic2_testing

import (
	"fmt"
	"testing"
	"github.com/jejejery/go-logic/logic2"
	"github.com/stretchr/testify/assert"
)

func TestLogic2_05(t *testing.T) {
	result := logic2.Logic2_5(9)
	expected := [][]int{
		{1,3,5,7,9,11,13,15,17},
		{35,33,31,29,27,25,23,21,19},
		{37,39,41,43,45,47,49,51,53},
		{71,69,67,65,63,61,59,57,55},
		{73,75,77,79,81,83,85,87,89},
		{107,105,103,101,99,97,95,93,91},
		{109,111,113,115,117,119,121,123,125},
		{143,141,139,137,135,133,131,129,127},
		{145,147,149,151,153,155,157,159,161},

	}

	assert.Equal(t, result, expected)
	fmt.Println("TestLogic2_02 done")
}