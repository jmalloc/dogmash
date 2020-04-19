package ui

import (
	"github.com/dogmatiq/dodeca/logging"
	"github.com/google/uuid"
	"github.com/jmalloc/dogmash/plugin"
	"github.com/rivo/tview"
)

type controller struct {
	App    *tview.Application
	Pages  *tview.Pages
	Logger logging.Logger

	stack   []string
	current string
}

func (c *controller) Run(
	plugins []*plugin.Plugin,
) error {
	tearDown := setupStyle()
	defer tearDown()

	c.App = tview.NewApplication()
	c.Pages = tview.NewPages()

	c.current = c.AddPage(
		"",
		mainPage(c),
	)

	c.AddPage(
		"plugins",
		pluginsPage(c, plugins),
	)

	inner := tview.NewFrame(c.Pages)
	inner.SetTitle(" dogmash ")
	// inner.SetTitleColor(tcell.Color208)
	// inner.SetBorderColor(tcell.Color236)
	inner.SetBorderPadding(1, 1, 2, 2)
	inner.SetBorder(true)

	outer := tview.NewFrame(inner)
	outer.SetBorders(1, 1, 0, 0, 2, 2)
	c.App.SetRoot(outer, true)

	return c.App.Run()
}

func (c *controller) AddPage(id string, item tview.Primitive) string {
	if id == "" {
		id = uuid.New().String()
	}

	c.Pages.AddPage(
		id,
		item,
		true,
		c.current == "",
	)

	return id
}

func (c *controller) Push(page string) func() {
	return func() {
		c.PushNow(page)
	}
}

func (c *controller) PushNow(page string) {
	c.stack = append(c.stack, c.current)
	c.current = page
	c.Pages.SwitchToPage(c.current)
}

func (c *controller) Pop() func() {
	return func() {
		c.PopNow()
	}
}
func (c *controller) PopNow() {
	i := len(c.stack) - 1

	if i < 0 {
		c.App.Stop()
		return
	}

	c.current = c.stack[i]
	c.stack = c.stack[:i]
	c.Pages.SwitchToPage(c.current)
}
