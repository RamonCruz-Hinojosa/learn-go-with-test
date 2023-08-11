package integers_test

import (
	"fmt"

	"github.com/RamonCruz-Hinojosa/learn-go-with-test/integers"
)

func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
