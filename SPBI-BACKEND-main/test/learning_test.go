package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	//go print(5, "Rifqi")
	//print(5, "Dimas")

	var message = make(chan string)

	var print = func(text string) {
		data := fmt.Sprintf("Closure Print: %s", text)
		message <- data
	}

	go print("Test 123")
	go print("Test 456")

	var message1 = <-message
	fmt.Printf("%s\n", message1)

	var message2 = <-message
	fmt.Printf("%s\n", message2)

	//var input string
	//fmt.Scanln(&input)
}

//func print(qty int, name string) {
//	for i := 1; i <= qty; i++ {
//		fmt.Println(strconv.Itoa(i) + " " + name)
//	}
//}

func helloWorld() {
	fmt.Println("Hello World")
}

func TestGoRoutine(t *testing.T) {
	go helloWorld()
	fmt.Println("Finish")

	time.Sleep(10000)
}
