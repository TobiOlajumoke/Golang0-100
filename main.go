// 	age := 22
// 	name := "Tobi"
// 	// // //String
// 	// // var nameOne string = "Bankai"
// 	// // var nameTwo = "katoom"
// 	// // var nameThree string

// 	// // fmt.Println(nameOne, nameTwo, nameThree)

// 	// // var nameOnes = "Bank"
// 	// // var nameThrees = "boy"
// 	// // fmt.Println(nameOnes, nameTwo, nameThrees)

// 	// // nameFour := 30

// 	// // fmt.Println(nameFour)

// 	// // int
// 	// var ageOne int = 10
// 	// var ageTwo = 30
// 	// agethree := 15

// 	// fmt.Println(ageOne, ageTwo, agethree)

// 	// // // bits & memory

// 	// // var numOne int8 = 25
// 	// // var numTwo int8 = -128
// 	// // var numThree uint16 = 256

// 	// var scoreOne float32 = 25.98
// 	// var scoreTwo float64 = 122345667
// 	// scoreThree := 3.4444
// 	// fmt.Println(scoreOne, scoreTwo, scoreThree)

// 	//Print
// 	fmt.Print("hello, ")
// 	fmt.Print("World! \n")
// 	fmt.Print("new line \n")
// 	//Println
// 	fmt.Println("hello ninjas")
// 	fmt.Println("goodbye ninjas")
// 	fmt.Println("my age is", age, "and my name is", name)

// 	// Printf (formatted strings) %_ = format specifier

// 	fmt.Printf("my age is %v and my name is %v\n", age, name)
// 	fmt.Printf("my age is %q and my name is %q\n", age, name)
// 	fmt.Printf("age is of type %T\n", age)
// 	fmt.Printf("you scored %f points! \n", 225.55)
// 	fmt.Printf("you scored %0.1f points! \n", 225.55)

// 	// Sprintf (save formatted strings)
// 	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
// 	fmt.Println("the saved string is :", str)
// }

// package main

// import "fmt"

// func main() {
// 	//  var ages [3]int = [3]int{20, 25, 30}
// 	var ages = [3]int{20, 25, 30}
// 	names := [4]string{"tobi", "ayo", "segun", "victor"}
// 	names[1] = "bito"
// 	fmt.Println(ages, len(ages))
// 	fmt.Println(names, len(names))

// 	// slices (use arrays under the hood)
// 	var scores = []int{100, 50, 60}
// 	scores[2] = 25
// 	scores = append(scores, 85)

// 	fmt.Println(scores, len(scores))

// 	//slice ranges
// 	rangeOne := names[1:3]
// 	rangeTwo := names[2:]
// 	rangeThree := names[:3]

// 	fmt.Println(rangeOne, rangeTwo, rangeThree)

// 	rangeOne = append(rangeOne, "koopa")
// 	fmt.Println(rangeOne)

// }
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	// greeting := "hello there friends!"
// 	// // fmt.Println(strings.Contains(greeting, "hello!"))
// 	// // fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
// 	// // fmt.Println(strings.ToUpper(greeting))
// 	// // fmt.Println(strings.Index(greeting, "ll"))
// 	// // fmt.Println(strings.Index(greeting, "th"))
// 	// fmt.Println(strings.Split(greeting, " "))

// 	// // the original value is unchanged
// 	// fmt.Println("orignal string value =", greeting)

// 	// ages := []int{45, 20, 35, 75, 60, 50, 25}
// 	// sort.Ints(ages)
// 	// fmt.Println(ages)
// 	// index := sort.SearchInts(ages, 30)
// 	// fmt.Println(index)

// 	// names := []string{"tobi", "victor", "ayo", "segun", "mayowa"}
// 	// sort.Strings(names)
// 	// fmt.Println(names)
// 	// fmt.Println(sort.SearchStrings(names, "tobi"))
// }

// package main

// import "fmt"

// func main() {
// 	// 	x := 0
// 	// 	for x < 5 {
// 	// 		fmt.Println("value of x is :", x)
// 	// 		x++
// 	// 	}

// 	// }

// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Println("value of i is:", i)
// 	// }

// 	// running through slice of string
// 	names := []string{"victor", "tobi", "ayo", "mayowa"}
// 	// 	for i := 0; i < len(names); i++ {
// 	// 		fmt.Println(names[i])
// 	// 	}

// 	// for index, value := range names {
// 	// 	fmt.Printf("the position at index %v is %v \n", index, value)
// 	// }

// 	// to get the value alone
// 	// for _, value := range names {
// 	// 	fmt.Printf("the value is %v \n", value)
// 	// 	value = "new string"
// 	// }

// // fmt.Println(names)
package main

import "crypto/rand"

func main() {
	// age := 45
	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	// names := []string{"tobi", "victor", "ayo", "segun", "bito"}
	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break
	// 	}
	// 	fmt.Printf("the value at pos %v is %v \n ", index, value)
	// }

	// side quest rasing 2 to the power of 11
	// var x float64
	// x = math.Exp2(11)
	// fmt.Println(x)

	//sum of numers 1-10
	// sum:= 0
	// for i:= 0; i < 10; i++ {
	// 	sum += i

	// }

	// fmt.Println(sum)
	// fmt.Println("The list of number from 1 - 1000")
	// for i := 1; i <= 1000; i++ {
	// 	fmt.Print(i, " ")
	// }

	min := 1
	max := 10
	v := rand.Intn(max-min) + min
	println(v)

}
