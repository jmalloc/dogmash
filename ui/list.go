package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func newList() *tview.List {
	list := tview.NewList()
	list.SetHighlightFullLine(true)
	list.SetShortcutColor(tcell.Color24)
	list.SetSelectedTextColor(tcell.ColorWhite)
	list.SetSelectedBackgroundColor(tcell.Color24)
	return list
}

func addBackItem(c *controller, list *tview.List) {
	list.AddItem(
		"back",
		"return to the previous page",
		'b',
		c.Pop(),
	)
}
