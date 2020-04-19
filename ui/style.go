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

func setupColorStyle() {
	tview.Styles.TitleColor = tcell.Color208
	tview.Styles.BorderColor = tcell.Color236
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack

	tview.Styles.PrimaryTextColor = tcell.Color249
	tview.Styles.SecondaryTextColor = tcell.Color241
	tview.Styles.TertiaryTextColor = tcell.Color238

	// These have not been configured yet and are all set to pink so that they
	// stand out when do they do appear.
	tview.Styles.ContrastBackgroundColor = tcell.ColorPink
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorPink
	tview.Styles.GraphicsColor = tcell.ColorPink
	tview.Styles.InverseTextColor = tcell.ColorPink
	tview.Styles.ContrastSecondaryTextColor = tcell.ColorPink
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

func newList() *tview.List {
	list := tview.NewList()
	list.SetHighlightFullLine(true)
	list.SetShortcutColor(tcell.Color24)
	list.SetSelectedTextColor(tcell.ColorWhite)
	list.SetSelectedBackgroundColor(tcell.Color24)
	return list
}
