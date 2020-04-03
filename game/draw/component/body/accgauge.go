package body

import (
	"musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/rank"

	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// AccGauge draws accuracy guage and player rank
func AccGauge(CurrentSentence beatmap.Sentence, achievementRate float64, rank rank.Rank) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		RankPosX := int(Constants.WindowWidth * achievementRate)

		//正解率ゲージ　100%なら赤色
		Acc := CurrentSentence.GetAccuracy()
		GaugeArea := area.FromXYWH(0, 60, int(Acc), 3)
		if Acc == 1 {
			helper.DrawFillRect(Renderer, Constants.RedColor, GaugeArea)
		} else {
			helper.DrawFillRect(Renderer, Constants.GreenThickColor, GaugeArea)
		}
		//正解率ゲージの上に出るランク
		helper.DrawText(Renderer,
			pos.FromXY(RankPosX, 168),
			helper.RightAlign, helper.SystemFont,
			rank.Text(), Constants.TypedTextColor)
	}
}
