package surferua

import (
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {
	fmt.Println(New().Android().Chrome().String())
	fmt.Println(New().Firefox().String())
	fmt.Println(New().String())
	fmt.Println(New().Phone().String())
	fmt.Println(New().String())
}
