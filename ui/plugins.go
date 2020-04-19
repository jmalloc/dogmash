package ui

import (
	"context"
	"fmt"
	"path"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/dodeca/logging"
	"github.com/gdamore/tcell"
	"github.com/jmalloc/dogmash/plugin"
	"github.com/rivo/tview"
)

func pluginsPage(
	c *controller,
	plugins []*plugin.Plugin,
) tview.Primitive {
	menu := newList()

	c.Pages.AddPage(
		"plugin-load",
		pluginLoadModal(c, menu),
		true,
		false,
	)

	menu.AddItem(
		"load plugin",
		"load a Dogma application plugin",
		'l',
		func() {
			c.Pages.SendToFront("plugin-load")
			c.Pages.ShowPage("plugin-load")
		},
	)

	addBackItem(c, menu)

	for _, p := range plugins {
		addPageForPlugin(c, p, menu)
	}

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

func pluginLoadModal(
	c *controller,
	menu *tview.List,
) tview.Primitive {
	form := tview.NewForm()
	form.SetButtonsAlign(tview.AlignRight)
	form.SetLabelColor(tcell.ColorWhite)

	file := tview.NewInputField()
	file.SetLabel("Path")
	file.SetPlaceholder("/path/to/file.so")
	file.SetAutocompleteFunc(completeFilePath)

	ok := func() {
		p, err := plugin.Load(file.GetText())
		if err != nil {
			// TODO:
		} else {
			addPageForPlugin(c, p, menu)
		}

		c.Pages.HidePage("plugin-load")
	}

	cancel := func() {
		c.Pages.HidePage("plugin-load")
	}

	form.AddFormItem(file)
	form.AddButton("OK", ok)
	form.AddButton("CANCEL", cancel)
	form.SetCancelFunc(cancel)

	frame := tview.NewFrame(form)
	frame.SetTitle(" load plugin ")
	frame.SetBorder(true)
	frame.SetBorders(1, 0, 0, 0, 2, 2)
	frame.SetBorderColor(tview.Styles.TitleColor)

	frame.AddText(
		"Enter the path to the Dogma plugin file ...",
		true,
		tview.AlignLeft,
		tview.Styles.PrimaryTextColor,
	)

	container := tview.NewFlex()
	container.SetBorderPadding(2, 2, 10, 10)
	container.SetDirection(tview.FlexRow)
	container.AddItem(frame, 9, 0, true)

	return container
}

func addPageForPlugin(
	c *controller,
	p *plugin.Plugin,
	menu *tview.List,
) {
	base := path.Base(p.File)

	id := c.AddPage(
		"",
		pluginPage(c, p),
	)

	menu.InsertItem(
		-3, // before "load" and "back" options
		fmt.Sprintf("%s plugin", base),
		fmt.Sprintf("browse the applications in the '%s' plugin", base),
		0,
		c.Push(id),
	)

	menu.SetCurrentItem(0)
}
