package main

import (
	"fmt"
)

func Curry(f func(string) string) func(string) string {
	return func(s string) string {
		return f(s)
	}
}
func CurryAddXYZ(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x + y + z
		}
	}
}

func main() {
	fmt.Println(CurryAddXYZ(1)(2)(3)) // 6

}
