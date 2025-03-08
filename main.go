package main

import (
	"fmt"
	"math"
)

type combo struct {
	x, y, z int
}

func fermat(w, n int) []*combo {
	var ok bool
	var result []*combo

	for x := 1; x <= w; x++ {
		for y := 1; y <= w; y++ {
			for z := 1; float64(z) <= math.Pow(float64(x), float64(n))+math.Pow(float64(y), float64(n)); z++ {
				ok = math.Pow(float64(x), float64(n))+math.Pow(float64(y), float64(n)) == math.Pow(float64(z), float64(n))
				if ok {
					result = append(result, &combo{x, y, z})
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("input: [%d %d]\n", 10, 2)
	for _, v := range fermat(10, 2) {
		fmt.Printf("%d %d %d\n", v.x, v.y, v.z)
	}

	fmt.Println("----")

	fmt.Printf("input: [%d %d]\n", 20, 2)
	for _, v := range fermat(20, 2) {
		fmt.Printf("%d %d %d\n", v.x, v.y, v.z)
	}

	fmt.Println("----")

	fmt.Printf("input: [%d %d]\n", 10, 3)
	for _, v := range fermat(10, 3) {
		fmt.Printf("%d %d %d\n", v.x, v.y, v.z)
	}

	fmt.Println("----")

	fmt.Printf("input: [%d %d]\n", 100, 7)
	for _, v := range fermat(100, 7) {
		fmt.Printf("%d %d %d\n", v.x, v.y, v.z)
	}
}
