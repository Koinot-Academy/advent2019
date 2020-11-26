package main

import (
	"testing"
)

func TestIsHorizontal(t *testing.T) {
	var seg Segment
	var expected, actual bool

	seg = Segment{FirstPoint: Point{1, 8}, SecPoint: Point{7, 8}}
	expected = true
	actual = seg.IsHorizontal()
	if actual != expected {
		t.Errorf("Segment should be horizontal %v", seg)
	}

	seg = Segment{FirstPoint: Point{1, 1}, SecPoint: Point{1, 8}}
	expected = false
	actual = seg.IsHorizontal()
	if actual != expected {
		t.Errorf("Segment should not be horizontal %v", seg)
	}
}

func TestIsVertical(t *testing.T) {
	var seg Segment
	var expected, actual bool

	seg = Segment{FirstPoint: Point{1, 8}, SecPoint: Point{7, 8}}
	expected = false
	actual = seg.IsVertical()
	if actual != expected {
		t.Errorf("Segment should not be vertical %v", seg)
	}

	seg = Segment{FirstPoint: Point{1, 1}, SecPoint: Point{1, 8}}
	expected = true
	actual = seg.IsVertical()
	if actual != expected {
		t.Errorf("Segment should be vertical %v", seg)
	}
}

func TestAreParallel(t *testing.T) {
	var fSeg Segment
	var sSeg Segment
	var expected bool
	var actual bool

	// Both horizontals
	fSeg = Segment{FirstPoint: Point{1, 8}, SecPoint: Point{7, 8}}
	sSeg = Segment{FirstPoint: Point{3, 4}, SecPoint: Point{9, 4}}
	expected = true

	actual = AreParallel(fSeg, sSeg)
	if actual != expected {
		t.Error("fSeg and sSeg should be parallel.")
	}

	// Both verticals
	fSeg = Segment{FirstPoint: Point{1, 1}, SecPoint: Point{1, 8}}
	sSeg = Segment{FirstPoint: Point{6, 5}, SecPoint: Point{6, 2}}
	expected = true

	actual = AreParallel(fSeg, sSeg)
	if actual != expected {
		t.Error("fSeg and sSeg should be parallel.")
	}

	// First vertical second horizontal
	fSeg = Segment{FirstPoint: Point{1, 1}, SecPoint: Point{1, 8}}
	sSeg = Segment{FirstPoint: Point{3, 4}, SecPoint: Point{8, 4}}
	expected = false

	actual = AreParallel(fSeg, sSeg)
	if actual != expected {
		t.Error("fSeg and sSeg should be perpendicular.")
	}
}

func TestIntersects(t *testing.T) {
	var fSeg, sSeg Segment
	var expected Point
	var actual *Point

	// First vertical, second horizontal
	fSeg = Segment{FirstPoint: Point{7, 8}, SecPoint: Point{7, 4}}
	sSeg = Segment{FirstPoint: Point{4, 6}, SecPoint: Point{9, 6}}
	expected = Point{7, 6}

	actual = fSeg.Intersects(sSeg)
	if *actual != expected {
		t.Errorf("Segments should have crossed at %v, found %v.", expected, actual)
	}

	// First horizontal, second vertical
	sSeg = Segment{FirstPoint: Point{9, 6}, SecPoint: Point{4, 6}}
	fSeg = Segment{FirstPoint: Point{7, 8}, SecPoint: Point{7, 4}}
	expected = Point{7, 6}

	actual = fSeg.Intersects(sSeg)
	if *actual != expected {
		t.Errorf("Segments should have crossed at %v, found %v.", expected, actual)
	}

	// First horizontal, second vertical, not crossing
	fSeg = Segment{FirstPoint: Point{76, -29}, SecPoint: Point{159, -29}}
	sSeg = Segment{FirstPoint: Point{101, 118}, SecPoint: Point{101, 47}}
	actual = fSeg.Intersects(sSeg)
	if actual != nil {
		t.Errorf("Segments should not have crossed, found %v.", actual)
	}

	// First vertical, second horizontal, not crossing
	fSeg = Segment{FirstPoint: Point{1, 118}, SecPoint: Point{1, -29}}
	sSeg = Segment{FirstPoint: Point{24, 45}, SecPoint: Point{180, 45}}
	actual = fSeg.Intersects(sSeg)
	if actual != nil {
		t.Errorf("Segments should not have crossed, found %v.", actual)
	}

	// Parallel segments
	sSeg = Segment{FirstPoint: Point{9, 6}, SecPoint: Point{4, 6}}
	fSeg = Segment{FirstPoint: Point{2, 4}, SecPoint: Point{7, 4}}

	actual = fSeg.Intersects(sSeg)
	if actual != nil {
		t.Errorf("Segments should not have crossed, found %v.", actual)
	}

}

func TestContainsPoint(t *testing.T) {
	var expected, actual bool
	var segment Segment
	var point Point

	// Vertical contains point
	segment = Segment{Point{1, 1}, Point{1, 8}}
	point = Point{1, 4}
	expected = true
	actual = segment.ContainsPoint(point)
	if actual != expected {
		t.Errorf("Point %v should be considered in segment %v.", point, segment)
	}

	// Horizontal contains point
	segment = Segment{Point{4, 5}, Point{10, 5}}
	point = Point{8, 5}
	expected = true
	actual = segment.ContainsPoint(point)
	if actual != expected {
		t.Errorf("Point %v should be considered in segment %v.", point, segment)
	}

	// Point is one segment end
	segment = Segment{Point{4, 5}, Point{10, 5}}
	point = Point{4, 5}
	expected = true
	actual = segment.ContainsPoint(point)
	if actual != expected {
		t.Errorf("Point %v should be considered in segment %v.", point, segment)
	}

	// Point is one segment end
	segment = Segment{Point{1, 1}, Point{1, 8}}
	point = Point{1, 8}
	expected = true
	actual = segment.ContainsPoint(point)
	if actual != expected {
		t.Errorf("Point %v should be considered in segment %v.", point, segment)
	}

	// Segment horizontal point outside
	segment = Segment{Point{1, 1}, Point{1, 8}}
	point = Point{4, 5}
	expected = false
	actual = segment.ContainsPoint(point)
	if actual != expected {
		t.Errorf("Point %v should be considered outside segment %v.", point, segment)
	}

	// Segment vertical point outside
	segment = Segment{Point{4, 5}, Point{10, 5}}
	point = Point{2, 4}
	expected = false
	actual = segment.ContainsPoint(point)
	if actual != expected {
		t.Errorf("Point %v should be considered outside segment %v.", point, segment)
	}
}
