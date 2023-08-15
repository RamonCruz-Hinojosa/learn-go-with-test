package main

// practice exercises here
func main() {

	// //strings
	// // var nameOne string = "mario"
	// // var nameTwo = "luigi"
	// // var nameThree string
	// // fmt.Println(nameOne, nameTwo, nameThree)

	// // nameOne = "peach"
	// // nameThree = "bowser"
	// // fmt.Println(nameOne, nameTwo, nameThree)

	// // nameFour := "yoshi"
	// // fmt.Println(nameFour)

	// //ints
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// // fmt.Println(ageOne, ageTwo, ageThree)

	// // // //bits and memory
	// // // var numOne int8 = 25
	// // // var numTwo int8 = -128
	// // // var numThree uint = 24

	// // var scoreOne float32 = -1.5
	// // var scroeTwo float64 = 1234123412341234.7
	// // scoreThree := 1.5

	// //Print
	// fmt.Print("hello, ")
	// fmt.Print("world! \n")
	// fmt.Print("world! \n")

	// fmt.Println("hello ninjas!")
	// fmt.Println("goodbye ninjas!")

	// age := 36
	// name := "shaun"

	// // fmt.Println("my age is", age, "and my name is ", name)

	// // //Printf (formatted strings) %_ = format specifier
	// // fmt.Printf("may age is %q and my name is %q \n", age, name)
	// // fmt.Printf("age is of type %T\n", age)
	// // fmt.Printf("you scored %f points! \n", 225.55)
	// // fmt.Printf("you scored %0.1f points! \n", 225.55) //rounds up

	// // //Sprintf (save formatted strings)
	// // var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	// // fmt.Println("the saved string is:", str)

	// //arrays and slices
	// // var ages [3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}

	// var names = []string{}
	// names[1] = "luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// //slices (use arrays under the hood)
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 865)

	// fmt.Println(scores, len(scores))

	// //slice ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]
	// rangeFour := names[1:]

	// fmt.Println(rangeFour)

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "koopa")
	// fmt.Println(rangeOne)

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, " "))
	//the original value is unchanged
	// fmt.Println("original string value =", greeting)

	// ages := []int{45, 20, 35, 30, 75, 50, 60, 25}

	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "bowser"))

	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	// for _, value := range names {
	// 	fmt.Printf("the value is %v \n", value)
	// 	value = "new string"
	// }

	// fmt.Println(names)
}

func Hello() {}
