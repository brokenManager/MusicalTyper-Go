package body

import (
	"math"
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
)

type TimeGauge struct{}

var (
	foregroundColor = DrawHelper.GetMoreBlackishColor(Constants.BackgroundColor, 50)
	backgroundColor = DrawHelper.GetMoreBlackishColor(Constants.BackgroundColor, 25)
)

func (s TimeGauge) Draw(c *DrawComponent.DrawContext) {
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

	RebodyingTimeGaugeWidth := int(math.Floor(Ratio * Constants.WindowWidth))
	DrawHelper.DrawFillRect(c.Renderer, backgroundColor, 0, 60, Constants.WindowWidth, 130)
	DrawHelper.DrawFillRect(c.Renderer, foregroundColor, 0, 60, RebodyingTimeGaugeWidth, 130)
}