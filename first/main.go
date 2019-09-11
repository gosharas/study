package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Sdfg", math.Pi)
	fmt.Println(math.Sqrt(64))
	fmt.Println(add(3, 4))
	fmt.Println(add1(1, 1))
	a, b := swap("One", "Two")
	fmt.Println(a, b)
	fmt.Println(split(5))
	for i := 0; i < 100; i++ {
		if i > 98 {
			fmt.Println(i)
		}
	}
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}

func add(a int, b int) int {
	return a + b
}

func add1(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return a, b
}

func split(sum int) (x, y int) {
	defer fmt.Println("LOLOLOL")
	fmt.Println("loldo")
	x = sum + 2
	y = sum + 4
	return

}
