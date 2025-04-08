package loader

import (
	"fmt"

	"github.com/spf13/viper"
)

type EnvLoader struct{}

func NewEnvLoader() *EnvLoader {
	return &EnvLoader{}
}

func (e *EnvLoader) Name() string {
	return "env"
}

func (e *EnvLoader) Load(target any) error {
	v := viper.New()
	v.AutomaticEnv()

	if err := v.Unmarshal(target); err != nil {
		return fmt.Errorf("unmarshal env: %w", err)
	}

	return nil
}
