package loader

type ConfigSource interface {
	Load(target any) error
	Name() string
}
