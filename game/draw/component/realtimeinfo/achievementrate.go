package realtimeinfo

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"
)

// AchievementRate presents achievement rate
type AchievementRate struct{}

// Draw draws achievement rate by percent text
func (s AchievementRate) Draw(c *DrawComponent.DrawContext) {
	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(Constants.Margin+320, 430),
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"達成率", Constants.TypedTextColor)

	Text := fmt.Sprintf("%05.1f%%", c.GameState.GetAchievementRate(false)*100)
	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(Constants.Margin+330, 430),
		DrawHelper.LeftAlign, DrawHelper.BigFont,
		Text, Constants.BlueThickColor)
}
