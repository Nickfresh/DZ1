package main

import "fmt"

//функция "Циклическая ротация"
func Solution(A []int, K int) []int {
	N := len(A)
	if N > 1 {
		pos := N - (K % N)
		if pos != N {
			return append(A[pos:N], A[0:pos]...)
		} else {
			return A
		}
	}
	return A
}

func main() {
	A := []int{3, 8, 9, 7, 6}
	fmt.Println(Solution(A, 3))
}
