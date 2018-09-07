package main

import (
	"image"
	c "image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	p "rayt/primitive"
)

const (
	tmin = 0.1
	tmax = 100
)

var (
	while = p.NewColor(1, 1, 1)
	blue  = p.NewColor(0.5, 0.7, 1)
)

func color(r p.Ray, w p.World) p.Color {
	hited, rec := w.Hit(r, tmin, tmax)
	if hited {
		return rec.Normal.Direction.Add(1.0).Mul(0.5).Color()
	}
	dir := r.Direction.Normalize()
	t := 0.5 * (dir.Y + 1.0)
	return while.Vec().Mul(1 - t).Add(blue.Vec().Mul(t)).Color()
}

func main() {
	ns := 100
	camera := p.DefaultCamera()
	world := p.World{}
	world.Add(p.NewSphere(p.NewPoint(0, 0, -1), 0.5))
	world.Add(p.NewSphere(p.NewPoint(0, -100.5, -1), 100))
	img := image.NewNRGBA(image.Rect(0, 0, camera.Width, camera.Height))
	for y := camera.Height - 1; y >= 0; y-- {
		for x := 0; x < camera.Width; x++ {
			var col p.Color = p.NewColor(0, 0, 0)
			for s := 0; s < ns; s++ {
				r := camera.GetRay(float64(x)+rand.Float64(), float64(y)+rand.Float64())
				col = col.Add(color(r, world))
			}
			col = col.Div(float64(ns))
			ir := uint8(255.99 * col.R)
			ig := uint8(255.99 * col.G)
			ib := uint8(255.99 * col.B)
			img.Set(x, camera.Height-y-1, c.NRGBA{
				R: ir,
				G: ig,
				B: ib,
				A: 255,
			})
		}
	}
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
