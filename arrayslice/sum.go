package arrayslice

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// func NewFunc(numbers []int, numbers2 []int) []int {
// 	var numbers3 []int
// 	for i, value := range numbers {
// 		numbers3 = append(numbers3, numbers2[i]+value)
// 	}
// 	return numbers3
// }

// func calsfunc() {
// 	slice1 := []int{4, 5, 6, 7, 8}
// 	slice2 := []int{1, 2, 3, 4, 5}
// 	slice3 := append(slice2, slice1...)
// 	numbers3 := NewFunc(slice1, slice2)
// 	fmt.Println(numbers3, slice3)
// }
