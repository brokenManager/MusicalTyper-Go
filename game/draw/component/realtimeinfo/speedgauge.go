package realtimeinfo

import (
	"fmt"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	fastSpeedGaugeAnimateColor      = constants.RedColor.Darker(30)
	normalSpeedGaugeForegroundColor = constants.GreenThinColor.Darker(50)
)

// SpeedGauge draws players's typing speed by text and color
func SpeedGauge(typingSpeed int, FrameCount int) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		helper.DrawText(Renderer,
			pos.FromXY(constants.Margin, 382),
			helper.LeftAlign, helper.SystemFont,
			"タイピング速度", constants.TypedTextColor)

		Area := area.FromXYWH(constants.Margin, 405, constants.WindowWidth-constants.Margin*2, 20)

		if typingSpeed > 4 {
			//4key/secを超えていたら、赤色でアニメーションs
			Color := constants.RedColor
			if !(FrameCount%10 < 5) {
				Color = fastSpeedGaugeAnimateColor
			}
			helper.DrawFillRect(Renderer, Color, Area)
		} else {
			//そうでなければ普通に描画。
			helper.DrawFillRect(Renderer, constants.GreenThinColor, Area)

			GaugeWidth := typingSpeed / 4 * (constants.WindowWidth * 2)
			helper.DrawFillRect(Renderer, normalSpeedGaugeForegroundColor,
				area.FromXYWH(constants.Margin, 405, GaugeWidth, 20),
			)
		}
		helper.DrawText(Renderer,
			pos.FromXY(constants.WindowWidth/2, 402),
			helper.Center, helper.SystemFont,
			fmt.Sprintf("%2d Char/sec", typingSpeed), constants.TextColor,
		)
	}
}
