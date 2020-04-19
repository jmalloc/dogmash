package ui

import (
	"github.com/jmalloc/dogmash/internal/globalflags"
	"github.com/spf13/cobra"
)

// Command is the CLI command for the interactive UI.
func Command(c *cobra.Command, args []string) error {
	plugins, err := globalflags.LoadPlugin(c)
	if err != nil {
		return err
	}

	ct := &controller{}
	return ct.Run(plugins)
}
