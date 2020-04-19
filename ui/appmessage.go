package ui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/configkit/message"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (v *visitor) visitMessages() {
	var sorted []message.Name
	for n := range v.config.MessageNames().All() {
		sorted = append(sorted, n)
	}
	sort.Slice(sorted, func(i, j int) bool {
		a := sorted[i].String()
		b := sorted[i].String()
		return typeName(a) < typeName(b)
	})

	nodes := map[message.Role]*tview.TreeNode{}

	for _, r := range message.Roles {
		n := tview.NewTreeNode(r.String() + "s")
		expandOnSelect(n, false)
		v.messages.AddChild(n)
		nodes[r] = n
	}

	roles := v.config.MessageNames().All()

	for _, n := range sorted {
		n := n // capture loop variable
		r := roles[n]

		node := tview.NewTreeNode(typeName(n.String()))
		nodes[r].AddChild(node)

		v.pages.AddPage(
			n.String(),
			v.messageDetailsPage(n, r),
			true,
			false,
		)

		node.SetSelectedFunc(func() {
			v.pages.SwitchToPage(n.String())
		})
	}
}

func (v *visitor) messageDetailsPage(
	n message.Name,
	r message.Role,
) tview.Primitive {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	props := tview.NewTable()
	flex.AddItem(props, 4, 0, false)

	props.SetCellSimple(0, 0, "Go Type:")
	props.SetCellSimple(0, 1, typeName(n.String()))
	props.SetCellSimple(1, 0, "Go Package:")
	props.SetCellSimple(1, 1, packageName(n.String()))
	props.SetCellSimple(2, 0, "Role:")
	props.SetCellSimple(2, 1, r.String())

	producers := sortedHandlerList(v.config.Handlers().ProducersOf(n))
	consumers := sortedHandlerList(v.config.Handlers().ConsumersOf(n))
	var lines []string

	switch r {
	case message.CommandRole:
		lines = append(lines, "handled by")
		lines = append(lines, consumers...)
		lines = append(lines, "\nexecuted by:")
		lines = append(lines, producers...)
	case message.EventRole:
		lines = append(lines, "handled by")
		lines = append(lines, consumers...)
		lines = append(lines, "\nrecorded by:")
		lines = append(lines, producers...)
	case message.TimeoutRole:
		lines = append(lines, "scheduled by:")
		lines = append(lines, consumers...)
	}

	handlers := tview.NewTextView()
	handlers.SetText(strings.Join(lines, "\n"))
	handlers.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyESC:
			v.controller.PopNow()
		case tcell.KeyTAB:
			v.controller.App.SetFocus(v.tree)
		}
	})

	flex.AddItem(handlers, 0, 1, true)

	return flex
}

func sortedHandlerList(handlers configkit.HandlerSet) []string {
	var lines []string

	if len(handlers) == 0 {
		return []string{
			" • (none)",
		}
	}

	for _, h := range handlers {
		lines = append(
			lines,
			tview.Escape(
				fmt.Sprintf(
					" • %s %s",
					h.Identity().Name,
					h.HandlerType(),
				),
			),
		)
	}

	sort.Strings(lines)

	return lines
}
