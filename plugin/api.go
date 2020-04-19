package plugin

import v1 "github.com/jmalloc/dogmash/plugin/v1"

// Plugin is the interface used to interact with plugins of any version.
type Plugin interface {
	v1.API
}
