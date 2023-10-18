package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(mul(2, 3))
	fmt.Println(sub(2, 3))
	fmt.Println(div(4, 2))
}

func sum(i ...float64) float64 {
	sum := 0.0
	for _, n := range i {
		sum += n
	}
	return sum
}

func sub(i ...float64) float64 {
	result := i[0]
	for _, n := range i[1:] {
		result -= n
	}
	return result
}

func mul(i ...float64) float64 {
	mul := 1.0
	for _, n := range i {
		mul *= n
	}
	return mul
}

func div(i float64, n float64) float64 {
	return i / n
}
