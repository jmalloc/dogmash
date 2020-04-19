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

func (v *visitor) handlerNode(cfg configkit.Handler) *tview.TreeNode {
	id := cfg.Identity()
	node := tview.NewTreeNode(id.Name)

	v.pages.AddPage(
		id.Key,
		v.handlerDetailsPage(cfg),
		true,
		false,
	)

	node.SetSelectedFunc(func() {
		v.pages.SwitchToPage(id.Key)
	})

	return node
}

func (v *visitor) handlerDetailsPage(cfg configkit.Handler) tview.Primitive {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	props := tview.NewTable()
	flex.AddItem(props, 6, 0, false)

	props.SetCellSimple(0, 0, "Go Type:")
	props.SetCellSimple(0, 1, typeName(cfg.TypeName()))
	props.SetCellSimple(1, 0, "Go Package:")
	props.SetCellSimple(1, 1, packageName(cfg.TypeName()))
	props.SetCellSimple(2, 0, "Identity Name:")
	props.SetCellSimple(2, 1, cfg.Identity().Name)
	props.SetCellSimple(3, 0, "Identity Key:")
	props.SetCellSimple(3, 1, cfg.Identity().Key)
	props.SetCellSimple(4, 0, "Handler Type:")
	props.SetCellSimple(4, 1, cfg.HandlerType().String())

	names := cfg.MessageNames()
	var lines []string

	switch cfg.HandlerType() {
	case configkit.AggregateHandlerType,
		configkit.IntegrationHandlerType:
		consumed := sortedMessageList(names.Consumed)
		produced := sortedMessageList(names.Produced)
		lines = append(lines, "handles commands:")
		lines = append(lines, consumed...)
		lines = append(lines, "\nrecords events:")
		lines = append(lines, produced...)
	case configkit.ProcessHandlerType:
		consumed := sortedMessageList(names.Consumed)
		commands := sortedMessageList(names.Produced.FilterByRole(message.CommandRole))
		timeouts := sortedMessageList(names.Produced.FilterByRole(message.TimeoutRole))
		lines = append(lines, "handles events:")
		lines = append(lines, consumed...)
		lines = append(lines, "\nexecutes commands:")
		lines = append(lines, commands...)
		lines = append(lines, "\nschedules timeouts:")
		lines = append(lines, timeouts...)
	case configkit.ProjectionHandlerType:
		consumed := sortedMessageList(names.Consumed)
		lines = append(lines, "handles events:")
		lines = append(lines, consumed...)
	}

	messages := tview.NewTextView()
	messages.SetText(strings.Join(lines, "\n"))
	messages.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyESC:
			v.controller.PopNow()
		case tcell.KeyTAB:
			v.controller.App.SetFocus(v.tree)
		}
	})

	flex.AddItem(messages, 0, 1, true)

	return flex
}

func sortedMessageList(names message.NameRoles) []string {
	var lines []string

	if len(names) == 0 {
		return []string{
			" • (none)",
		}
	}

	for n := range names {
		lines = append(
			lines,
			tview.Escape(
				fmt.Sprintf(
					" • %s",
					typeName(n.String()),
				),
			),
		)
	}

	sort.Strings(lines)

	return lines
}
