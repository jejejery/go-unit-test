package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_11(t *testing.T) {
	result := logic1.Logic1_11(10)
	expected := []string {"buzz", "1", "buzz", "3", "buzz", "6", "buzz", "12", "buzz", "24"}
	not_expected := []string {"buzz", "1", "buzz", "3", "buzz", "6", "buzz", "12", "buzz", "25"}

	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_11 done")

}