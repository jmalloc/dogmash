package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jmalloc/dogmash/internal/globalflags"
	"github.com/jmalloc/dogmash/ui"
	"github.com/spf13/cobra"
)

func main() {
	rand.Seed(time.Now().Unix())

	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	root := &cobra.Command{
		Use:   "dogmash",
		Short: "An interactive shell for exploring Dogma applications",
		Args:  cobra.NoArgs,
		RunE:  ui.Command,
	}

	globalflags.DefineLoadPlugin(root)

	return root.Execute()
}
