package p2p

import (
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/p2p/receiver"
)

// Bootstrapped constants that are used as key in bootstrap context
const (
	BootstrappedP2PClient string = "BootstrappedP2PClient"
)

// Bootstrapper implements Bootstrapper with p2p details
type Bootstrapper struct{}

// Bootstrap initiates p2p server and client into context
func (b Bootstrapper) Bootstrap(ctx map[string]interface{}) error {
	cfg, ok := ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	if !ok {
		return errors.New("config not initialised")
	}

	registry, ok := ctx[documents.BootstrappedRegistry].(*documents.ServiceRegistry)
	if !ok {
		return errors.New("registry not initialised")
	}

	srv := &p2pServer{config: cfg, handlerCreator: func() *receiver.Handler {
		return receiver.New(cfg, registry)
	}}
	ctx[bootstrap.BootstrappedP2PServer] = srv
	ctx[BootstrappedP2PClient] = srv
	return nil
}
