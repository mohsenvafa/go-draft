package loader

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type YAMLLoader struct {
	Path string
}

func NewYAMLLoader(path string) *YAMLLoader {
	return &YAMLLoader{Path: path}
}

func (y *YAMLLoader) Name() string {
	return "yaml"
}

func (y *YAMLLoader) Load(target any) error {
	v := viper.New()
	v.SetConfigFile(y.Path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(*os.PathError); ok {
			return nil // ignore if not found
		}
		return fmt.Errorf("read config: %w", err)
	}

	if err := v.Unmarshal(target); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}

	return nil
}
