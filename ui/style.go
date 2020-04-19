package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func setupStyle() func() {
	styles := tview.Styles
	borders := tview.Borders

	setupColorStyle()
	setupBorderStyle()

	return func() {
		tview.Styles = styles
		tview.Borders = borders
	}
}

// See https://upload.wikimedia.org/wikipedia/commons/1/15/Xterm_256color_chart.svg
func setupColorStyle() {
	tview.Styles.TitleColor = tcell.Color208
	tview.Styles.BorderColor = tcell.Color236
	tview.Styles.GraphicsColor = tcell.Color24

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
	tview.Styles.ContrastBackgroundColor = tcell.Color235
	tview.Styles.MoreContrastBackgroundColor = tcell.Color238

	tview.Styles.PrimaryTextColor = tcell.Color249
	tview.Styles.SecondaryTextColor = tcell.Color241
	tview.Styles.TertiaryTextColor = tcell.Color238

	// These have not been configured yet and are all set to bright colors so
	// that they stand out when do they do appear.
	tview.Styles.InverseTextColor = tcell.ColorGreen
	tview.Styles.ContrastSecondaryTextColor = tcell.Color233
}

func setupBorderStyle() {
	// tview.Borders.Horizontal = tview.BoxDrawingsLightHorizontal
	// tview.Borders.Vertical = tview.BoxDrawingsLightVertical
	// tview.Borders.TopLeft = tview.BoxDrawingsLightDownAndRight
	// tview.Borders.TopRight = tview.BoxDrawingsLightDownAndLeft
	// tview.Borders.BottomLeft = tview.BoxDrawingsLightUpAndRight
	// tview.Borders.BottomRight = tview.BoxDrawingsLightUpAndLeft

	// tview.Borders.LeftT = tview.BoxDrawingsLightVerticalAndRight
	// tview.Borders.RightT = tview.BoxDrawingsLightVerticalAndLeft
	// tview.Borders.TopT = tview.BoxDrawingsLightDownAndHorizontal
	// tview.Borders.BottomT = tview.BoxDrawingsLightUpAndHorizontal
	// tview.Borders.Cross = tview.BoxDrawingsLightVerticalAndHorizontal

	tview.Borders.HorizontalFocus = tview.BoxDrawingsHeavyHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsHeavyVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsHeavyDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsHeavyDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsHeavyUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsHeavyUpAndLeft
}
