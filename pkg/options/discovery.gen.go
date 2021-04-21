package options

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// pkg/options/discovery.yaml

type (
	DiscoveryOpt struct {
		Enabled bool `env:"DISCOVERY_ENABLED"`
	}
)

// Discovery initializes and returns a DiscoveryOpt with default values
func Discovery() (o *DiscoveryOpt) {
	o = &DiscoveryOpt{
		Enabled: false,
	}

	fill(o)

	// Function that allows access to custom logic inside the parent function.
	// The custom logic in the other file should be like:
	// func (o *Discovery) Defaults() {...}
	func(o interface{}) {
		if def, ok := o.(interface{ Defaults() }); ok {
			def.Defaults()
		}
	}(o)

	return
}
