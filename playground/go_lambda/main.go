package main

import (
	"fmt"
)

func Curry(f func(string) string) func(string) string {
	return func(s string) string {
		return f(s)
	}
}

func main() {
	fmt.Println("Hello, " + Curry(func(s string) string {
		return s + "!"
	})("World"))

}
