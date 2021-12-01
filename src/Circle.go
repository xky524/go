package main

import "fmt"
import "math"

func main()  {
	var radius1 float64 = 10
	var radius2 float64 = 20
	var area1 = radius1*radius1*math.Pi
	var area2 = radius2*radius2*math.Pi
	fmt.Printf("%.2f\n", area2+area1)
}
