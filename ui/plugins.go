package ui

import (
	"context"
	"fmt"
	"path"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/dodeca/logging"
	"github.com/jmalloc/dogmash/plugin"
	"github.com/rivo/tview"
)

func pluginsPage(
	c *controller,
	plugins []*plugin.Plugin,
) tview.Primitive {
	menu := newList()

	for _, p := range plugins {
		base := path.Base(p.File)

		id := c.AddPage(
			"",
			pluginPage(c, p),
		)

		menu.AddItem(
			fmt.Sprintf("%s plugin", base),
			fmt.Sprintf("browse the applications in the '%s' plugin", base),
			0,
			c.Push(id),
		)
	}

	menu.AddItem(
		"load plugin",
		"load a Dogma application plugin",
		'l',
		func() {

		},
	)

	addBackItem(c, menu)

	return menu
}

func pluginPage(
	c *controller,
	p *plugin.Plugin,
) tview.Primitive {
	menu := newList()

	for _, n := range p.ListApplications() {
		app, closer, err := p.OpenApplication(
			context.Background(),
			n,
		)
		if err != nil {
			logging.Log(c.Logger, "unable to open %s application: %s", err)
			continue
		}
		defer closer.Close()

		cfg := configkit.FromApplication(app)
		k := cfg.Identity().Key

		id := c.AddPage(
			"",
			appPage(c, cfg),
		)

		menu.AddItem(
			fmt.Sprintf("%s application", n),
			fmt.Sprintf("browse the contents of the '%s' application (%s)", n, k),
			0,
			c.Push(id),
		)
	}

	addBackItem(c, menu)

	return menu
}
