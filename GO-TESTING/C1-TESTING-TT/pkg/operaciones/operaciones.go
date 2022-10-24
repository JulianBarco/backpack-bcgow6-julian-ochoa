package operaciones

import "fmt"

func Add(a int, b int) (int, error) {
	if a == 0 {
		return 0,fmt.Errorf("a no puede ser:%d", a)
	}

	return a + b, nil
}
