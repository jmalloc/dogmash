package plugin

import v1 "github.com/jmalloc/dogmash/plugin/v1"

// Plugin is the API used to interact with plugins of any version.
type Plugin struct {
	v1.API

	// File is the path to the plugin file.
	File string
}
