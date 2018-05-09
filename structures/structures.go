package structures

// Structure visible within the current package and outside
type Point struct {
	X int // Visible
	Y int // Visible
}

// Structure visible only within the current package
type point2 struct {
	x int // Not visible
	y int // Not visible
}
