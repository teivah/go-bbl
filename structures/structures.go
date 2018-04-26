package structures

type Point struct {
	// Structure visible within the current package and outside
	X int // Visible
	Y int // Visible
}

type point2 struct {
	// Structure visible only within the current package
	x int // Not visible
	y int // Not visible
}
