package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func concat(a string, b string) string {
	return fmt.Sprint(a, b)
}

func multiply(a int, b int) int {
	return a * b
}

func doStuffWithInts(someFunc func(int, int) int, a int, b int) int {
	return someFunc(a, b)
}

func main() {
	index := 5
	var answer = 42
	var eight int
	fmt.Printf("%d\n", eight)
	eight = 8
	fmt.Printf("%d, %d, %d\n", index, answer, eight)

	const BALL = 3
	fmt.Println(BALL)

	fmt.Println(concat("con", "cat"))

	fmt.Println(doStuffWithInts(multiply, 4, 5))

	var numbers [2]int
	numbers[0] = 1
	numbers[1] = 2
	//numbers[2] = 3
	//numbers = append(numbers, 3)
	fmt.Println(numbers)

	chars := make([]byte, 2)
	chars[0] = 'a'
	chars[1] = 'b'
	//chars[2] = 'c'
	chars = append(chars, 'c')
	chars[2] = 'c'
	fmt.Println(chars)

	hello := []string{"h", "e", "l", "l", "o"}
	fmt.Println(hello[1:len(hello)])
	bello := hello
	bello[0] = "b"
	fmt.Println(bello)
	fmt.Println(hello)
	yello := make([]string, len(hello))
	copy(yello, bello)
	yello[0] = "y"
	fmt.Println(yello)
	fmt.Println(bello)

	numToString := make(map[int]string)
	numToString[0] = "a"
	numToString[1] = "b"
	// or numToString = map[int]string{0: "a", 1: "b"}
	fmt.Println(numToString)
	value := numToString[1]
	fmt.Println(value)

	_, ok := numToString[5]
	fmt.Println(ok)

	question := "How many roads lead to Rome?"
	var ptr *string
	ptr = &question
	fmt.Println(ptr)
	fmt.Println(*ptr)

	dir, _ := os.Getwd()
	filename := fmt.Sprintf("%s/hey.txt", dir)
	if file, err := ioutil.ReadFile(filename); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s doesn't exist, silly!\n", filename)
		} else {
			panic(err)
		}
	} else {
		fmt.Println(string(file))
	}
}
