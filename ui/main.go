package ui

import "github.com/rivo/tview"

func mainPage(c *controller) tview.Primitive {
	menu := newList()

	menu.AddItem(
		"plugins",
		"browse applications by plugin",
		'p',
		c.Push("plugins"),
	)

	menu.AddItem(
		"quit",
		"quit dogmash",
		'q',
		c.Pop(),
	)

	return menu
}
