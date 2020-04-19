package globalflags

import (
	"github.com/jmalloc/dogmash/plugin"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
)

// DefineLoadPlugin defines the --load-plugin flag.
func DefineLoadPlugin(c *cobra.Command) {
	c.PersistentFlags().StringArrayP(
		"load-plugin",
		"p",
		nil,
		"load a Dogma application plugin",
	)
}

// LoadPlugin loads and returns all plugins specified via the --load-plugin
// flag.
//
// It returns the plugins that load successfully, even if there is an error.
func LoadPlugin(c *cobra.Command) ([]*plugin.Plugin, error) {
	files, err := c.Flags().GetStringArray("load-plugin")
	if err != nil {
		return nil, err
	}

	var (
		plugins []*plugin.Plugin
		errors  error
	)

	for _, f := range files {
		p, err := plugin.Load(f)

		if err != nil {
			errors = multierr.Append(errors, err)
		} else {
			plugins = append(plugins, p)
		}
	}

	return plugins, errors
}
