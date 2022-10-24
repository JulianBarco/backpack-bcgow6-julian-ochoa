package dividir

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividrGood(t *testing.T) {
	//arrange
	num1 := 20
	num2 := 2
	//act
	c, err := Dividir(num1, num2)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, c, 10)
}

func TestDividirBad(t *testing.T) {
	//arrange
	num1 := 20
	num2 := 0
	errorEsperado := fmt.Sprintf("num2 no puede ser %d", num2)
	//act
	_, err := Dividir(num1, num2)
	//assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado)
}
