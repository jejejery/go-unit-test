package logic1_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jejejery/go-logic/logic1"
	"fmt"
)

func TestLogic1_10(t *testing.T) {
	result := logic1.Logic1_10(10)
	expected := []string {"2", "fizz", "4", "fizz", "8", "fizz", "16", "fizz", "32", "fizz"}
	not_expected := []string {"2", "fizz", "4", "fizz", "8", "fizz", "16", "fizz", "64", "fizz"}

	assert.Equal(t, result, expected)
	assert.NotEqual(t, result, not_expected)
	fmt.Println("TestLogic1_10 done")

}