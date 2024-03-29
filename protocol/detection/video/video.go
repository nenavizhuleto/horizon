package video

type Source struct {
	// Uniform Resource Identifier
	URI string

	// Width(X) and Height(Y) of the source, on which consumers could rely on
	Dimensions Dimensions
}

// Might used image.Point, but implemented for better marshal control
type Dimensions struct {
	Width  int
	Height int
}

// Might used image.Rectangle, but will not
type Position struct {
	X      int
	Y      int
	Width  int
	Height int
}
