package main

import "fmt"

type Team int

const (
	SERVICES Team = iota
	INTERFACES
	CORE
	OPIT
)

func deferMe() {
	defer func() {
		fmt.Print(" World")
	}()
	fmt.Print("Hello")
}

func checkAnswer(answer int) {
	if answer != 42 {
		panic("NOOO!!")
	}
	fmt.Println(answer)
}

func goPanic() {
	checkAnswer(42)
	checkAnswer(41)
}

func goPanicAndRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Keep calm and eat your sushi")
		}
	}()
	checkAnswer(42)
	checkAnswer(41)
	checkAnswer(43)
}

func stringOrInt(i int) interface{} {
	if i < 0 {
		return "negative"
	}
	return i
}

func main() {
	fmt.Printf("%d, %d\n", SERVICES, CORE)

	deferMe()
	fmt.Println()

	goPanicAndRecover()
	fmt.Println("Phew, that was close.")

	if val, ok := stringOrInt(1).(int); ok {
		fmt.Printf("Int: %d\n", val)
	}
	if val, ok := stringOrInt(-1).(string); ok {
		fmt.Printf("String: %s\n", val)
	}

	bla := 2

	checkAnswer(41)
	fmt.Println("That was the right answer!") // This never happens
}
