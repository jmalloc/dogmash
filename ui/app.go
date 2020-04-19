package ui

import (
	"context"
	"strings"

	"github.com/dogmatiq/configkit"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func appPage(
	c *controller,
	cfg configkit.Application,
) tview.Primitive {
	v := &visitor{
		controller: c,
		config:     cfg,
		tree:       tview.NewTreeView(),
		pages:      tview.NewPages(),
	}

	cfg.AcceptVisitor(context.Background(), v)

	grid := tview.NewGrid()
	grid.SetColumns(0, 0)
	grid.SetRows(0)
	grid.SetGap(1, 2)

	appFrame := tview.NewFrame(v.tree)
	appFrame.SetTitle(" application structure ")
	appFrame.SetBorder(true)
	appFrame.SetBorders(1, 0, 0, 0, 2, 2)
	appFrame.SetBorderColor(tview.Styles.TitleColor)

	grid.AddItem(
		appFrame,
		0, 0, // row/col
		1, 1, // row/col span
		0, 0, // mins,
		true, // focus
	)

	detailsFrame := tview.NewFrame(v.pages)
	detailsFrame.SetTitle(" details ")
	detailsFrame.SetBorder(true)
	detailsFrame.SetBorders(1, 0, 0, 0, 2, 2)
	detailsFrame.SetBorderColor(tview.Styles.TitleColor)

	grid.AddItem(
		detailsFrame,
		0, 1, // row/col
		1, 1, // row/col span
		0, 0, // mins
		false, // focus
	)

	v.tree.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyESC:
			c.PopNow()
		case tcell.KeyTAB:
			c.App.SetFocus(detailsFrame)
		}
	})

	return grid
}

type visitor struct {
	controller *controller
	config     configkit.Application

	tree  *tview.TreeView
	pages *tview.Pages

	app      *tview.TreeNode
	handlers *tview.TreeNode
	messages *tview.TreeNode

	aggregates   *tview.TreeNode
	processes    *tview.TreeNode
	integrations *tview.TreeNode
	projections  *tview.TreeNode
}

func (v *visitor) VisitApplication(ctx context.Context, cfg configkit.Application) error {
	v.app = tview.NewTreeNode(cfg.Identity().Name)
	v.tree.SetRoot(v.app)
	v.tree.SetCurrentNode(v.app)

	v.handlers = tview.NewTreeNode("handlers")
	expandOnSelect(v.handlers, true)
	v.app.AddChild(v.handlers)

	v.messages = tview.NewTreeNode("messages")
	expandOnSelect(v.messages, true)
	v.app.AddChild(v.messages)

	v.visitMessages()

	return cfg.Handlers().AcceptVisitor(ctx, v)
}

func (v *visitor) VisitAggregate(_ context.Context, cfg configkit.Aggregate) error {
	if v.aggregates == nil {
		v.aggregates = tview.NewTreeNode("aggregates")
		expandOnSelect(v.aggregates, false)
		v.handlers.AddChild(v.aggregates)
	}

	v.aggregates.AddChild(
		v.handlerNode(cfg),
	)

	return nil
}

func (v *visitor) VisitProcess(_ context.Context, cfg configkit.Process) error {
	if v.processes == nil {
		v.processes = tview.NewTreeNode("processes")
		expandOnSelect(v.processes, false)
		v.handlers.AddChild(v.processes)
	}

	v.processes.AddChild(
		v.handlerNode(cfg),
	)

	return nil
}

func (v *visitor) VisitIntegration(_ context.Context, cfg configkit.Integration) error {
	if v.integrations == nil {
		v.integrations = tview.NewTreeNode("integrations")
		expandOnSelect(v.integrations, false)
		v.handlers.AddChild(v.integrations)
	}

	v.integrations.AddChild(
		v.handlerNode(cfg),
	)

	return nil
}

func (v *visitor) VisitProjection(_ context.Context, cfg configkit.Projection) error {
	if v.projections == nil {
		v.projections = tview.NewTreeNode("projections")
		expandOnSelect(v.projections, false)
		v.handlers.AddChild(v.projections)
	}

	v.projections.AddChild(
		v.handlerNode(cfg),
	)

	return nil
}

func typeName(t string) string {
	i := strings.LastIndexByte(t, '.')
	if i != -1 {
		return t[i+1:]
	}

	return t
}

func packageName(t string) string {
	i := strings.LastIndexByte(t, '.')
	if i != -1 {
		return t[:i]
	}

	return ""
}
