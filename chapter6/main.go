package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(math.Sin(math.Pi)-0) < math.Pow(10, -15))
}
