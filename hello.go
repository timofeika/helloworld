package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {

	//	var quantity int
	//	var length, width float64
	//	var customerName string

	//	quantity = 2
	//	length, width = 1.2, 2.5
	//	customerName = "Mc Murphy"

	// if u declare vars with something = not need to put exactly var type
	//	var testint = 4
	//	var testfl = 1.1
	//	var testtr = "just expain variables"

	// if u dont initialize vars it will be contain 0
	//	var Myint int
	//	var Myfloat float64

	// for string it will be '', for bool it will be false
	//	var Mystring string
	//	var Mybool bool

	// also we can do like that
	//	quantit := 0
	floantit := 1.2
	//	stringit := "just one more sample"

	fmt.Println("Just playing with GO, by .devil")

	fmt.Println("Hello, world!")
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(2.34))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hello"))
	fmt.Println(int(floantit))

	// play with time
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)

	//play with strings
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	//play stdin
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if grade <= 100 {
		fmt.Println("Nice grade!")
	} else if grade > 100 {
		fmt.Println("too old")
	} else {
		fmt.Println("Fail!")
		log.Fatal(err)
	}

	fmt.Println(grade)

	fmt.Println("my executable name is:", os.Args[0])
	//playing string to numbers

}
