package hello

import "fmt"

func Hi() []string {
	var booking = []string{"name", "hello", "hi", "come", "phew"}

	return booking[:]

}

func Hello(){
	fmt.Print("welcome")
}
