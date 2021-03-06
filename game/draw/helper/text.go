package helper

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/draw/size"
	Logger "musicaltyper-go/game/logger"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// FontSize means size of font
type FontSize uint8

type textureWithSize struct {
	Width, Height int32
	Texture       *sdl.Texture
}

var (
	fontCache    = map[FontSize]*ttf.Font{}
	textureCache = map[FontSize]map[string]*textureWithSize{}
)

const (
	// BigFont is bigger font size
	BigFont FontSize = 72
	// JapaneseFont is font size for japanese
	JapaneseFont = 48
	// AlphabetFont is font size for ASCII
	AlphabetFont = 32
	// FullFont is font size for full width
	FullFont = 24
	// RankFont is font size for rank of player
	RankFont = 20
	// SystemFont is default font size
	SystemFont = 16
)

func genTextureCacheKey(Text string, Color color.Color) string {
	return fmt.Sprintf("%s,%s", Text, Color.String())
}

// return -> texture, isFromCache
func makeTexture(Renderer *sdl.Renderer, Size FontSize, Text string, Color color.Color, ShouldCache bool) (*textureWithSize, bool) {
	logger := Logger.NewLogger("makeTexture")

	CacheKey := genTextureCacheKey(Text, Color)
	Texture, TextureExists := textureCache[Size][CacheKey]

	if !TextureExists {
		Begin := time.Now()

		Font, FontExists := fontCache[Size]
		if !FontExists {
			LoadedFont, Error := ttf.OpenFont("./mplus-1m-medium.ttf", int(Size))
			logger.CheckError(Error)
			fontCache[Size] = LoadedFont
			Font = LoadedFont
		}

		RenderedText, Error := Font.RenderUTF8Blended(Text, *Color.ToSDLColor())
		logger.CheckError(Error)
		defer RenderedText.Free()

		TextureFromSurface, Error := Renderer.CreateTextureFromSurface(RenderedText)
		logger.CheckError(Error)
		Result := &textureWithSize{RenderedText.W, RenderedText.H, TextureFromSurface}

		if _, MapExists := textureCache[Size]; !MapExists {
			textureCache[Size] = map[string]*textureWithSize{}
		}

		if ShouldCache {
			textureCache[Size][CacheKey] = Result
			if Constants.PrintTextureLog {
				fmt.Printf("Created \"%s\" texture. Key: %s Size:%d. Took %dμs\n", Text, CacheKey, Size, time.Now().Sub(Begin).Microseconds())
			}
		}
		return Result, false
	}

	return Texture, true
}

func drawText(Renderer *sdl.Renderer, p pos.Pos, alignment AlignmentType, Size FontSize, Text string, Color color.Color, ShouldCache bool) (int, int) {
	logger := Logger.NewLogger("DrawText")
	if Text == "" {
		return 0, 0
	}

	Texture, isCached := makeTexture(Renderer, Size, Text, Color, ShouldCache)
	x, y := p.X(), p.Y()

	var Rect sdl.Rect
	switch alignment {
	case LeftAlign:
		Rect = sdl.Rect{
			X: int32(x),
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	case RightAlign:
		Rect = sdl.Rect{
			X: int32(x) - Texture.Width,
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	case Center:
		Rect = sdl.Rect{
			X: int32(x) - Texture.Width/2,
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	}

	Error := Renderer.Copy(Texture.Texture, nil, &Rect)
	logger.CheckError(Error)

	if !ShouldCache && !isCached {
		Texture.Texture.Destroy()
	}

	return int(Texture.Width), int(Texture.Height)
}

// DrawText renders text
func DrawText(Renderer *sdl.Renderer, p pos.Pos, alignment AlignmentType, Size FontSize, Text string, Color color.Color) (int, int) {
	return drawText(Renderer, p, alignment, Size, Text, Color, true)
}

func DrawTextWithoutCache(Renderer *sdl.Renderer, p pos.Pos, alignment AlignmentType, Size FontSize, Text string, Color color.Color) (int, int) {
	return drawText(Renderer, p, alignment, Size, Text, Color, false)
}

// DrawThickLine renders thick line
//fixme: 計算ガバガバなので斜めの線とか引くと多分バグる
func DrawThickLine(Renderer *sdl.Renderer, from, to pos.Pos, Color color.Color, Thickness int) {
	Color.ApplyColor(Renderer)
	Renderer.DrawRect(area.FromTwoPos(from, to).ToRect())
}

// DrawLine render line
func DrawLine(Renderer *sdl.Renderer, from, to pos.Pos, Color color.Color) {
	DrawThickLine(Renderer, from, to, Color, 1)
}

// GetTextSize calculates dimension of text by actual rendering
func GetTextSize(Renderer *sdl.Renderer, Size FontSize, Text string, Color color.Color) size.Size {
	Texture, _ := makeTexture(Renderer, Size, Text, Color, true)
	return size.FromWH(int(Texture.Width), int(Texture.Height))
}

// DrawFillRect renders filled rect
func DrawFillRect(Renderer *sdl.Renderer, Color color.Color, a area.Area) {
	Color.ApplyColor(Renderer)
	Renderer.FillRect(a.ToRect())
}

// DrawLineRect render rect by lines
func DrawLineRect(Renderer *sdl.Renderer, Color color.Color, a area.Area, thickness int) {
	Color.ApplyColor(Renderer)

	var (
		X         = int32(a.X())
		Y         = int32(a.Y())
		Width     = int32(a.W())
		Height    = int32(a.H())
		Thickness = int32(thickness)
	)

	Rects := []sdl.Rect{
		{
			X: X,
			Y: Y,
			W: Width,
			H: Thickness,
		},
		{
			X: X,
			Y: Y,
			W: Thickness,
			H: Height,
		},
		{
			X: X + Width - Thickness,
			Y: Y,
			W: Thickness,
			H: Height,
		},
		{
			X: X,
			Y: Y + Height - Thickness,
			W: Width,
			H: Thickness,
		},
	}

	Renderer.DrawRects(Rects)
}

// Quit destroys textures
func Quit() {
	for _, v := range textureCache {
		for _, t := range v {
			t.Texture.Destroy()
		}
	}

	for _, v := range fontCache {
		v.Close()
	}
}
