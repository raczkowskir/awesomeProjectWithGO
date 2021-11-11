package main

import (
	"fmt"
	_ "strings"
)

func functionWhichReturnNothing() {
	fmt.Println("Lets see what happen.")
}

func functionWhichReturnMultipleValues() (int, int) {
	var myFirstInt int = 0
	mySecondInt := 10
	//mySecondInt int = 10

	return myFirstInt, mySecondInt
}

func main() {
	fmt.Println("//////////// The basel line")

	functionWhichReturnNothing()
	print(functionWhichReturnMultipleValues())
	//fmt.Println("///////////// The last line")
}

//func someShit() {
//fmt.Println(runAnotherFunction(hello))
//fmt.Println(runAnotherFunction(world))
//
//fmt.Println(mySecondFunction(functionWithoutArguments))
//
//var myFirstVar, mySecondVar, myThirdVar interface{}
//
//mySecondVar = true
//mySecondVar = 0
//
//fmt.Println(myFirstVar, mySecondVar, myThirdVar)
//
//fmt.Println(myFirstFunction(true))
//
//var myFunnyVar = 11
//var myFunnyVar2 = 0
//
//if myFunnyVar == 10 {
//	fmt.Println("dziala jedwabiscie")
//} else if myFunnyVar != 11 || myFunnyVar2 == 0 {
//	fmt.Println("dziala mniej jedwabiscie")
//}
//
//var movieTitle interface{}
//movieTitle = "Raki"
//
//switch a := movieTitle.(type) {
//case string:
//	fmt.Println("That was a string type", a)
//case int:
//	fmt.Println("That was an int")
//case bool:
//	fmt.Println("That was a bool")
//default:
//	fmt.Println("We do not know what was that")
//
//
//
//}

//switch movieTitle {
//case "Braveheart":
//	fmt.Println("Cast: Mell Gibson")
//case "Me Irene and Me":
//	fmt.Println("Cast: Jim Carrey")
//case "Touchless":
//	fmt.Println("Cast: Shon Connery")
//case "Just Friends":
//	fmt.Println("Cast: Emy Smart")
//default:
//	fmt.Println("Some random actors")
//
//}

//myInt := 99
//
//var myUniversalVar interface{}
//
//
//fmt.Println("Napisz jakies slowo:")
////fmt.Scanln(myUniversalVar)
//fmt.Println("Jestes typem: ", myUniversalVar)
//myUniversalVar = 10
//
//fmt.Println("Jestes typem: ", myUniversalVar)
//myUniversalVar = "dupa"
//
//fmt.Println("Jestes typem: ", myUniversalVar)
//myUniversalVar = true
//
//fmt.Println("Jestes typem: ", myUniversalVar)
//myUniversalVar = nil
//fmt.Println("Jestes typem: ", myUniversalVar)
//
//var myBool = true
//var myString string = "a to ciekawe"
//var dupa = 0.89
//var ears float64 = 0.005
//var weirdVariable = true
//
//myInt2 := 0.1
//
//myInt2 = 1
//
//myInt2 = 0.8
////myInt2 = int(0.8)
//
//
//fmt.Println(myInt2)
//fmt.Println(weirdVariable)
//fmt.Println(dupa, ears)
//fmt.Println(myString, myString)
//
//fmt.Println(
//	myBool,
//	"Dzialaj dziadu!",
//)
//fmt.Printf("Moja pierwsza zmienna %v\n", myInt )
//}

//func myFirstFunction(myFunnyBool bool) int {
//	if bool(true) {
//		return 0
//	}
//	return -1
//}
//
//func functionWithoutArguments() int {
//	return 999
//}
//
//func mySecondFunction(f func() int) int {
//	var dupa1 string = "dupa123"
//	print(dupa1[2:])
//	//var dupa2 string = strings.Trim
//	fmt.Println("Spacja")
//	print(dupa1)
//	//print(dupa2)
//
//	return f()
//}
//
//func runAnotherFunction(f func()string) string {
//	//return fmt.Sprintf("the function returned: %s", f())
//	return f()
//}
//
//func hello()string {
//	return "Hello"
//}
//
//func world()string {
//	return "World!"
//}

//func anotherFunction(int i)string {
//	return "I do not fit:( "
//}
