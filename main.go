package main

import (
	"fmt"
	"image"
	c "image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	p "rayt/primitive"
	"sync"
	"time"
)

const (
	tmin = 0.001
	tmax = 100
)

const maxDepth = 50

var (
	while  = p.NewColor(1, 1, 1)
	blue   = p.NewColor(0.5, 0.7, 1)
	purple = p.NewColor(0.54, 0.16, 0.88)
	orange = p.NewColor(1, 0.27, 0)
	brown  = p.NewColor(0.8, 0.6, 0.2)
	lgray  = p.NewColor(0.8, 0.8, 0.8)
)

var (
	purpleMatte = p.NewSphere(p.NewPoint(0, 0, -1), 0.5, p.NewMatte(purple))
	orangeMatte = p.NewSphere(p.NewPoint(0, -100.5, -1), 100, p.NewMatte(orange))
	brownMetal  = p.NewSphere(p.NewPoint(1, 0, -1), 0.5, p.NewMetal(brown, 0.3))
	sphereGlass = p.NewSphere(p.NewPoint(-1, 0, -1), 0.5, p.NewGlass(1.5))
)

func color(r p.Ray, w p.World, depth int) p.Color {
	hited, rec := w.Hit(r, tmin, tmax)
	if hited {
		if depth < 50 {
			bounced, bouncedRay := rec.Scatter(r, rec)
			if bounced {
				attenuation := rec.Material.Color()
				newColor := color(bouncedRay, w, depth+1)
				return attenuation.Mul(newColor)
			}
		}
		return p.Color{}
	}
	dir := r.Direction.Normalize()
	t := 0.8 * (dir.Y + 1.0)
	return while.Vec().Mul(1 - t).Add(blue.Vec().Mul(t)).Color()
}

const (
	nthreads = 16
	samples  = 500
	fov      = 90.0
	width    = 1920
	height   = 1080
	aperture = 0.0
)

func render(camera p.Camera, world p.World, samples int, x, y int, rnd *rand.Rand) c.NRGBA {
	var col p.Color = p.NewColor(0, 0, 0)
	for s := 0; s < samples; s++ {
		r := camera.GetRay(float64(x)+rnd.Float64(), float64(y)+rnd.Float64())
		col = col.Add(color(r, world, 0))
	}
	col = col.Div(float64(samples)).Gamma()
	ir := uint8(255.99 * col.R)
	ig := uint8(255.99 * col.G)
	ib := uint8(255.99 * col.B)
	return c.NRGBA{
		R: ir,
		G: ig,
		B: ib,
		A: 255,
	}
}

func main() {
	lookfrom := p.NewPoint(-2, 2, 1)
	lookat := p.NewPoint(0, 0, -1)
	camera := p.NewCamera(lookfrom, lookat, p.NewVector(0, 1, 0), fov, width, height, aperture)
	world := p.World{}
	world.Add(purpleMatte)
	world.Add(orangeMatte)
	world.Add(brownMetal)
	world.Add(sphereGlass)
	img := image.NewNRGBA(image.Rect(0, 0, camera.Width, camera.Height))

	start := time.Now()
	var wg sync.WaitGroup
	chunkSize := (camera.Height + nthreads - 1) / nthreads
	done := 0
	total := width * height
	for i := 0; i < nthreads; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println("Spawning process", i)
			pstart := time.Now()
			start := i * chunkSize
			end := (i + 1) * chunkSize
			if end > camera.Width {
				end = camera.Width
			}
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for y := start; y < end; y++ {
				for x := 0; x < camera.Width; x++ {
					rgb := render(camera, world, samples, x, y, rnd)
					img.Set(x, camera.Height-y-1, rgb)
					done++
					fmt.Printf("\rDone: %.2f%%", float64(done)/float64(total)*100)
				}
			}
			fmt.Println("\rProcess", i, "has finished in", time.Since(pstart))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done. Elapsed", time.Since(start))

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
