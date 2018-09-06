package main

import "fmt"

func main() {
	nx, ny := 200, 100
	fmt.Println("P3")
	fmt.Println(nx, ny)
	fmt.Println("255")
	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {
			r := float64(j) / float64(ny)
			g := float64(i) / float64(nx)
			b := 0.5
			ir := int64(r * 255.9)
			ig := int64(g * 255.9)
			ib := int64(b * 255.9)
			fmt.Println(ir, ig, ib)
		}
	}
}
