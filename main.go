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
package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))

}
