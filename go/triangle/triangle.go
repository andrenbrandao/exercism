// Package triangle defines methods to identify different types of triangles.
package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides identifies the type of a triangle given its sides.
func KindFromSides(a, b, c float64) Kind {
	if isNotTriangle(a, b, c) {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	}

	if isScalene(a, b, c) {
		return Sca
	}

	return Iso
}

func isNotTriangle(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isScalene(a, b, c float64) bool {
	return a != b && b != c && a != c
}
