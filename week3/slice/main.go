package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	s1 := arr[1:4]
	s2 := s1[2:6]

	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	fmt.Printf("arr: %T, s1: %T\n", arr, s1)

	s1[2] = "z"

	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := []int{1, 2, 3, 4, 5} // Slice literal, no value inside the square brackets

	fmt.Printf("%T, slice: %v, len: %d, cap: %d\n", s3, s3, len(s3), cap(s3))

	s4 := make([]int, 10) // make([]T, len)
	s5 := make([]int, 10, 20) // make([]T, len, cap)

	fmt.Printf("s4 => %T, slice: %v, len: %d, cap: %d\n", s4, s4, len(s4), cap(s4))
	fmt.Printf("S4 pointer: %p\n", s4)
	s4 = append(s4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("S4 pointer: %p\n", s4)
	fmt.Printf("s4 => %T, slice: %v, len: %d, cap: %d\n", s4, s4, len(s4), cap(s4))

	fmt.Printf("s5 => %T, slice: %v, len: %d, cap: %d\n", s5, s5, len(s5), cap(s5))
	fmt.Printf("S5 pointer: %p\n", s5)
	s5 = append(s5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("S5 pointer: %p\n", s5)
	fmt.Printf("s5 => %T, slice: %v, len: %d, cap: %d\n", s5, s5, len(s5), cap(s5))
}
