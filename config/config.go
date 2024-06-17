package config

import (
	"context"
	"github.com/spf13/viper"
	"io"
	"log"
	"strings"
)

type State interface {
	GetString(string) string
}

type Source interface {
	Name() string
	Load(context.Context, State) (io.Reader, error)
}

type Configuration struct {
	Port string `mapstructure:"port"`
}

func New(ctx context.Context, filePath string, sources ...Source) *Configuration {
	conf := &Configuration{}

	vip := viper.New()
	vip.SetConfigType("yaml")

	bindEnvironmentVariables(vip)
	readConfigFile(vip, filePath)
	for _, s := range sources {
		r, err2 := s.Load(ctx, vip)
		if err2 != nil {
			log.Fatalf("could not load config from source %q: %s", s.Name(), err2)
		}

		err2 = vip.MergeConfig(r)
		if err2 != nil {
			log.Fatalf("could not merge config from source %q: %s", s.Name(), err2)
		}
	}

	err := vip.Unmarshal(conf)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return conf
}

func bindEnvironmentVariables(vip *viper.Viper) {
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	vip.MustBindEnv("port") // PORT
}

func readConfigFile(vip *viper.Viper, configFilePath string) {
	if len(configFilePath) == 0 {
		return
	}
	vip.SetConfigFile(configFilePath)
	if err := vip.ReadInConfig(); err != nil {
		log.Fatalf("error reading configuration file, %s", err)
	}
}
