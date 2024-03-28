package detection

type Motion struct {
	// Video stream or file
	Source VideoSource

	// Position where motion is detected (optional)
	Position *VideoPosition
}

type VideoSource struct {
	// Uniform Resource Identifier
	URI string

	// Width(X) and Height(Y) of the source, on which consumers could rely on
	Dimensions VideoDimensions
}

// Might used image.Point, but implemented for better marshal control
type VideoDimensions struct {
	Width  int
	Height int
}

// Might used image.Rectangle, but will not
type VideoPosition struct {
	X      int
	Y      int
	Width  int
	Height int
}
