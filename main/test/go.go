package main

import "fmt"

func main() {
	const CON string = "the Json Map"
	var s string = "在下版本有何贵干"
	fmt.Println(s)
	var a uint8 = 1
	var b uint8 = 2
	fmt.Println(float32(a) / float32(b))
	fmt.Println("Hello, World!")
	fmt.Println(CON)
	fmt.Println("--------------99multiple----------------------")
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d = %d ", x, y, x*y)
		}
		fmt.Println(" ")
	}
	fmt.Println("----------------------------------------------")

}
