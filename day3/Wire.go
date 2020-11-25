package main

// Wire represents multiple segments connected with each other
type Wire struct {
	Segments []Segment
}

// NewWire initializes a wire with a slice of points
func NewWire(points []Point) Wire {
	var wire Wire

	wire.Segments = make([]Segment, 0, len(points)-1)
	for i, pointsLen := 1, len(points); i < pointsLen; i++ {
		wire.Segments = append(wire.Segments, NewSegment(points[i-1], points[i]))
	}
	return wire
}

// GetCrossingPoints returns a slice of points where both wires cross
// Returns nil if the two wires do not cross
func GetCrossingPoints(fWire, sWire Wire) []Point {
	var crossingPoints []Point

	for _, seg1 := range fWire.Segments {
		for _, seg2 := range sWire.Segments {
			// If cross but not in the starting point
			if crossingPoint := seg1.Crosses(seg2); crossingPoint != nil && seg1.FirstPoint.X != 1 && seg1.FirstPoint.Y != 1 {
				crossingPoints = append(crossingPoints, *crossingPoint)
			}
		}
	}
	return crossingPoints
}
