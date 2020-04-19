package plugin

import (
	"fmt"
	"path"
	"plugin"

	v1 "github.com/jmalloc/dogmash/plugin/v1"
)

// Load loads a Dogma plugin.
func Load(file string) (Plugin, error) {
	base := path.Base(file)

	p, err := plugin.Open(file)
	if err != nil {
		return nil, err
	}

	s, err := p.Lookup(v1.Symbol)
	if err != nil {
		return nil, fmt.Errorf(
			"%s is not a valid plugin, the %s symbol is not defined",
			base,
			v1.Symbol,
		)
	}

	x, ok := s.(v1.API)
	if !ok {
		return nil, fmt.Errorf(
			"%s is not a valid plugin, the %s symbol has type %T, which does not impement the v1.API interface",
			base,
			v1.Symbol,
			s,
		)
	}

	return x, nil
}
