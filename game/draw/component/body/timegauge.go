package body

import (
	"math"
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/area"

	"github.com/veandco/go-sdl2/sdl"
)

func foregroundColor() *sdl.Color {
	return DrawHelper.GetMoreBlackishColor(Constants.BackgroundColor, 50)
}
func backgroundColor() *sdl.Color {
	return DrawHelper.GetMoreBlackishColor(Constants.BackgroundColor, 25)
}

// TimeGauge draws remainings time gauge
func TimeGauge(c *DrawComponent.DrawContext) {
	var Ratio float64
	if len(c.GameState.Beatmap.Notes) <= c.GameState.CurrentSentenceIndex+1 {
		Ratio = 1
	} else {
		CurrentSentenceStartTime := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex].Time
		NextSentenceStartTime := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex+1].Time
		CurrentSentenceDuration := NextSentenceStartTime - CurrentSentenceStartTime
		CurrentTimeInCurrentSentence := CurrentSentenceDuration - c.GameState.CurrentTime + CurrentSentenceStartTime
		Ratio = CurrentTimeInCurrentSentence / CurrentSentenceDuration
	}

	RemainingTimeGaugeWidth := int(math.Floor(Ratio * Constants.WindowWidth))
	DrawHelper.DrawFillRect(c.Renderer, backgroundColor(), area.FromXYWH(0, 60, Constants.WindowWidth, 130))
	DrawHelper.DrawFillRect(c.Renderer, foregroundColor(), area.FromXYWH(0, 60, RemainingTimeGaugeWidth, 130))
}
