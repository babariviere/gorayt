package primitive

// Hit is a record at where the ray hit an object
type Hit struct {
	Distance float64
	Normal   Ray
}

// Object is an object that can be hitted by light
type Object interface {
	Hit(r Ray, tmin, tmax float64) (bool, Hit)
}
