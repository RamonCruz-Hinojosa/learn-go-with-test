package iteration

import "fmt"

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	//Output: "aaaaa"
}
