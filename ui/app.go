package ui

import (
	"github.com/dogmatiq/configkit"
	"github.com/rivo/tview"
)

func appPage(
	c *controller,
	cfg configkit.Application,
) tview.Primitive {
	menu := newList()

	addBackItem(c, menu)

	return menu
}
