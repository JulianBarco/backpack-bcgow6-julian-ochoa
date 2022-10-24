package dividir

import "fmt"

func Dividir(num1, num2 int) (c int, err error) {

	if num2 == 0 {
		err = fmt.Errorf("no se puede dividir por cero")
	}
	c = num1 / num2
	return c,err
}
