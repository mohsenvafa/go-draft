package loader

import "fmt"

func LoadConfig(globalCfg, serviceCfg any, sources ...ConfigSource) error {
	for _, src := range sources {
		if globalCfg != nil {
			if err := src.Load(globalCfg); err != nil {
				return fmt.Errorf("load global config from %s: %w", src.Name(), err)
			}
		}
		if serviceCfg != nil {
			if err := src.Load(serviceCfg); err != nil {
				return fmt.Errorf("load service config from %s: %w", src.Name(), err)
			}
		}
	}
	return nil
}
