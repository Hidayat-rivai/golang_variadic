package main

import "fmt"

func main(){
	var rata_rata = hello(10,5,4,6,7,8,3,5,4,3,2)
	fmt.Printf("Rata Rata %.2f",rata_rata)
}

func hello(angka ...int)float64{
	var total int = 0
	for _, number := range angka {
		total+= number
	}
	var avg = float64(total) / float64(len(angka))
	return avg
}