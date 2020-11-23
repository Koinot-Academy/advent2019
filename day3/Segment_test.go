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

func TestCrosses(t *testing.T) {
	var fSeg, sSeg Segment
	var expected Point
	var actual *Point

	// First vertical, second horizontal
	fSeg = Segment{FirstPoint: Point{7, 8}, SecPoint: Point{7, 4}}
	sSeg = Segment{FirstPoint: Point{4, 6}, SecPoint: Point{9, 6}}
	expected = Point{7, 6}

	actual = fSeg.Crosses(sSeg)
	if *actual != expected {
		t.Errorf("Segments should have crossed at %v, found %v.", expected, actual)
	}

	// First horizontal, second vertical
	sSeg = Segment{FirstPoint: Point{9, 6}, SecPoint: Point{4, 6}}
	fSeg = Segment{FirstPoint: Point{7, 8}, SecPoint: Point{7, 4}}
	expected = Point{7, 6}

	actual = fSeg.Crosses(sSeg)
	if *actual != expected {
		t.Errorf("Segments should have crossed at %v, found %v.", expected, actual)
	}

	// Parallel segments
	sSeg = Segment{FirstPoint: Point{9, 6}, SecPoint: Point{4, 6}}
	fSeg = Segment{FirstPoint: Point{2, 4}, SecPoint: Point{7, 4}}

	actual = fSeg.Crosses(sSeg)
	if actual != nil {
		t.Errorf("Segments should not have crossed, found %v.", actual)
	}

}
