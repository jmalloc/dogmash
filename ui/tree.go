package ui

import "github.com/rivo/tview"

func expandOnSelect(n *tview.TreeNode, expanded bool) {
	n.SetExpanded(expanded)
	n.SetSelectedFunc(func() {
		n.SetExpanded(!n.IsExpanded())
	})
}
