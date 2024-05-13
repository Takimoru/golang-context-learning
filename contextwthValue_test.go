package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWthValue(t *testing.T) {

	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextA, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("a"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("d"))
	fmt.Println(contextF.Value("e"))
	// context hanya bisa menanyakan value yang ada di parent tidak bisa mendapatkan value dari child
	// alurnya child bertanya pada parent
}
