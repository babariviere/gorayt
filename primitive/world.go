package primitive

// World contains all hitable objects
type World struct {
	objs []Object
}

// Add a new object to the world
func (w *World) Add(obj Object) {
	w.objs = append(w.objs, obj)
}

// Hit check if there is an hit on an object
func (w World) Hit(r Ray, tmin, tmax float64) (bool, Hit) {
	hitAny := false
	record := Hit{Distance: tmax}

	for _, obj := range w.objs {
		hited, rec := obj.Hit(r, tmin, record.Distance)
		if hited {
			hitAny = true
			record = rec
		}
	}
	return hitAny, record
}
