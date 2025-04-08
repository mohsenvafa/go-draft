package global

type GlobalConfig struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	APIURL       string `mapstructure:"api_url"`
}
