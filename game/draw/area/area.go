package area

import (
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/draw/size"

	"github.com/veandco/go-sdl2/sdl"
)

// Area expresses area on screen
type Area struct {
	p pos.Pos
	s size.Size
}

// FromXYWH makes Area from x and y corrdinates, width, and height
func FromXYWH(x, y, w, h int) Area {
	return Area{
		p: pos.FromXY(x, y),
		s: size.FromWH(w, h),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if 0 < a {
		return a
	}
	return -a
}

// FromTwoPos returns area that contains specified two points
func FromTwoPos(a, b pos.Pos) Area {
	x := min(a.X(), b.X())
	y := min(a.Y(), b.Y())
	w := abs(a.X() - b.X())
	h := abs(a.Y() - b.Y())
	return FromXYWH(x, y, w, h)
}

// X returns x coordinate of area
func (a Area) X() int {
	return a.p.X()
}

// Y returns y coordinate of area
func (a Area) Y() int {
	return a.p.Y()
}

// W returns width of area
func (a Area) W() int {
	return a.s.W()
}

// H returns height of area
func (a Area) H() int {
	return a.s.H()
}

// ToRect casts Area to sdl.Rect
func (a Area) ToRect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(a.X()),
		Y: int32(a.Y()),
		W: int32(a.W()),
		H: int32(a.H()),
	}
}
