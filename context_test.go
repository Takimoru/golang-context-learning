package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)
	// Context background merupakan context yang isinya kosong

	todo := context.TODO()
	fmt.Println(todo)
	// sama dengan background todo merupakan context yang isinya kosong

}
