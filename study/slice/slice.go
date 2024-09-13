package slice

import (
	"fmt"
)

func SliceInit1() {
	s := make([]int, 3, 5) // int size is 3
	fmt.Println(s)
}

func SliceInit2() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
}

func SliceInit3() {
	s := []int{}
	fmt.Println(s)
}

func SliceInit4() {
	s := []int{}
	s = append(s, 1)
	fmt.Println(s)
}

func SliceInit5() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append(s1, 6)
	fmt.Println(s2)
}
