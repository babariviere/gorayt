package main

import (
	"fmt"
	p "rayt/primitive"
)

const (
	tmin = 0.1
	tmax = 100
)

func color(r p.Ray, w p.World) p.Color {
	hited, rec := w.Hit(r, tmin, tmax)
	if hited {
		return p.NewVector(rec.Normal.Direction.X+1, rec.Normal.Direction.Y+1, rec.Normal.Direction.Z+1).Mul(0.5).Color()
	}
	dir := r.Direction.Normalize()
	t := 0.5 * (dir.Y + 1.0)
	invt := 1.0 - t
	return p.NewColor(invt+t*0.5, invt+t*0.7, invt+t)
}

func main() {
	nx, ny := 200, 100
	fmt.Println("P3")
	fmt.Println(nx, ny)
	fmt.Println("255")
	lowerLeft := p.NewVector(-2, -1, -1)
	horizontal := p.NewVector(4, 0, 0)
	vertical := p.NewVector(0, 2, 0)
	origin := p.NewPoint(0, 0, 0)
	world := p.World{}
	world.Add(p.NewSphere(p.NewPoint(0, 0, -1), 0.5))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := p.NewRay(origin, lowerLeft.Add(horizontal.Mul(u).Add(vertical.Mul(v))))
			col := color(r, world)
			ir := uint8(255.99 * col.R)
			ig := uint8(255.99 * col.G)
			ib := uint8(255.99 * col.B)
			fmt.Println(ir, ig, ib)
		}
	}
}
