package config

import (
	"github.com/zer0day88/brick-test/pkg/environment"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type config struct {
	Port        string            `yaml:"port"`
	Environment environment.Level `yaml:"environment"`
	LogLevel    zerolog.Level     `yaml:"log_level"`
	Database    struct {
		Postgres struct {
			Host                  string `yaml:"host"`
			User                  string `yaml:"user"`
			Password              string `yaml:"password"`
			DbName                string `yaml:"dbname"`
			Port                  int    `yaml:"port"`
			SSLMode               string `yaml:"ssl_mode"`
			ConnectionMaxLifetime int    `yaml:"connection_max_lifetime"`
			ConnectionMaxOpen     int    `yaml:"connection_max_open"`
			ConnectionMaxIdle     int    `yaml:"connection_max_idle"`
			ConnectionMaxIdleTime int    `yaml:"connection_max_idle_time"`
		}
	}
	External struct {
		BankABC struct {
			BaseUrl string `yaml:"baseurl"`
		} `yaml:"bankabc"`
	} `yaml:"external"`
}

var (
	Key config
)

// Load config file will be injected by sidecar vault if using kubernetes
func Load(log zerolog.Logger) {

	readConfig(log)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Warn().Msgf("Config file changed: %s", e.Name)
		readConfig(log)
	})

}

func readConfig(log zerolog.Logger) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	viper.SetDefault("environment", "development")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal().Msgf("config file not found, %s", err.Error())
		} else {
			log.Fatal().Msgf("%s", err.Error())
		}
	}

	err := viper.Unmarshal(&Key)
	if err != nil {
		log.Fatal().Msgf("failed to unmarshal config struct, %v", err.Error())
	}
}
