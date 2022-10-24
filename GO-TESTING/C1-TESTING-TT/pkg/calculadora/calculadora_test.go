package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGood(t *testing.T) {
	// //arrange
	// num1 := 1
	// num2 := 2
	// //act
	// result := Add(num1, num2)
	// expected := 22
	// //assert
	// assert.Equal(t, expected, result)
	// // if result != expected {
	// // 	t.Errorf("Expected %d, got %d", expected, result)
	// // }

}

func TestAddBad(t *testing.T) {
	//arrange
	num1 := 0
	num2 := 5
	errorEsperado := fmt.Sprintf("a no puede ser:%d", num1)
	//act
	_, err := Add(num1, num2)

	//assert
	assert.Error(t, err)
	assert.ErrorContains(t, err,errorEsperado)

}
