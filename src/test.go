package main

import "fmt"

func main()  {
	var a uint = 60
	var b uint = 13
	var c uint = 0
	c=a & b
	fmt.Println(c)
	c= a | b
	fmt.Println(c)
	c =a ^ b
	fmt.Println(c)
	c=a<<2
	fmt.Println(c)
	c=a>>2
	fmt.Println(c)
	fmt.Println(" .\n" +
					"...\n" +
		           " ...")

}