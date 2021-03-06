package config

import "github.com/centrifuge/go-centrifuge/bootstrap"

// Bootstrap constants are keys to the value mappings in context bootstrap.
const (
	BootstrappedConfigFile string = "BootstrappedConfigFile"
)

// Bootstrapper implements bootstrap.Bootstrapper to initialise config package.
type Bootstrapper struct{}

// Bootstrap takes the passed in config file, loads the config and puts the config back into context.
func (*Bootstrapper) Bootstrap(context map[string]interface{}) error {
	if _, ok := context[BootstrappedConfigFile]; !ok {
		return ErrConfigFileBootstrapNotFound
	}
	cfgFile := context[BootstrappedConfigFile].(string)
	context[bootstrap.BootstrappedConfig] = LoadConfiguration(cfgFile)

	return nil
}
