package ui

import "github.com/rivo/tview"

func addBackItem(c *controller, list *tview.List) {
	list.AddItem(
		"back",
		"return to the previous page",
		'b',
		c.Pop(),
	)
}
