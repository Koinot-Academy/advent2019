package main

// Segment represents a segment in a cartesian coordinate system
type Segment struct {
	FirstPoint Point
	SecPoint   Point
}

// NewSegment initializes a segment with two points
func NewSegment(fPoint, sPoint Point) Segment {
	return Segment{FirstPoint: fPoint, SecPoint: sPoint}
}

// AreParallel returns true if fSeg is parallel with sSeg
func AreParallel(fSeg Segment, sSeg Segment) bool {
	areParallel := false

	if fSeg.IsHorizontal() && sSeg.IsHorizontal() {
		areParallel = true
	} else if fSeg.IsVertical() && sSeg.IsVertical() {
		areParallel = true
	}
	return areParallel
}

// IsHorizontal tells if seg is horizontally oriented
func (seg Segment) IsHorizontal() bool {
	return seg.FirstPoint.Y == seg.SecPoint.Y
}

// IsVertical tells if seg is vertically oriented
func (seg Segment) IsVertical() bool {
	return seg.FirstPoint.X == seg.SecPoint.X
}

// Intersects checks if seg and secSeg intersect with each other. If so, it returns a *Point
// or nil if segments do not cross.
func (seg Segment) Intersects(secSeg Segment) *Point {
	var intersection *Point = nil

	// If both segments are not parallel, they might cross
	if !AreParallel(seg, secSeg) {
		if seg.IsVertical() {
			secY := secSeg.FirstPoint.Y // secSeg is horizontal so Y is constant
			segX := seg.FirstPoint.X    // seg is vertical so X is constant
			if (secSeg.FirstPoint.X <= segX && secSeg.SecPoint.X >= segX) || (secSeg.SecPoint.X <= segX && secSeg.FirstPoint.X >= segX) {
				if (seg.FirstPoint.Y <= secY && seg.SecPoint.Y >= secY) || (seg.SecPoint.Y <= secY && seg.FirstPoint.Y >= secY) {
					intersection = &Point{seg.FirstPoint.X, secY}
				}
			}
		} else {
			// Is horizontal in either case
			secX := secSeg.FirstPoint.X // secSeg is vertical so X is constant
			segY := seg.FirstPoint.Y    // seg is horizontal so Y is constant
			if (secSeg.FirstPoint.Y <= segY && secSeg.SecPoint.Y >= segY) || (secSeg.SecPoint.Y <= segY && secSeg.FirstPoint.Y >= segY) {
				if (seg.FirstPoint.X <= secX && seg.SecPoint.X >= secX) || (seg.SecPoint.X <= secX && seg.FirstPoint.X >= secX) {
					intersection = &Point{secX, seg.FirstPoint.Y}
				}
			}
		}
	}
	return intersection
}

// ContainsPoint reports if a segment contains a point.
// If so, it returns true.
func (seg Segment) ContainsPoint(point Point) bool {
	var isPointInSegment bool

	if point == seg.FirstPoint || point == seg.SecPoint {
		isPointInSegment = true
	} else if seg.IsVertical() && point.X == seg.FirstPoint.X {
		if (seg.FirstPoint.Y > point.Y && seg.SecPoint.Y < point.Y) || (seg.SecPoint.Y > point.Y && seg.FirstPoint.Y < point.Y) {
			isPointInSegment = true
		}
	} else if seg.IsHorizontal() && point.Y == seg.FirstPoint.Y {
		if (seg.FirstPoint.X > point.X && seg.SecPoint.X < point.X) || (seg.SecPoint.X > point.X && seg.FirstPoint.X < point.X) {
			isPointInSegment = true
		}
	}
	return isPointInSegment
}
