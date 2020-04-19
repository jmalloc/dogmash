package v1

import (
	"context"
	"io"

	"github.com/dogmatiq/dogma"
)

// Symbol is the name of the global variable that must implement the plugin API.
const Symbol = "DogmaV1"

// API defines the v1 plugin API.
type API interface {
	ListApplications() []string
	OpenApplication(ctx context.Context, n string) (dogma.Application, io.Closer, error)
}
