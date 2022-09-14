package example

import (
	cfg "github.com/ferditatlisu/generic-client-go/genericclient"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config interface {
	GetConfig() (*ApplicationConfig, error)
}

type config struct{}

type ApplicationConfig struct {
	LocalApi LocalApiConfig
}

type LocalApiConfig struct {
	Api           cfg.ApiConfig
	Default       cfg.EndPoint
	LatePoint     cfg.EndPoint
	PostDefault   cfg.EndPoint
	PutDefault    cfg.EndPoint
	DeleteDefault cfg.EndPoint
}

func (c *config) GetConfig() (*ApplicationConfig, error) {
	configuration := ApplicationConfig{}
	env := getGoEnv()

	viperInstance := getViperInstance()
	err := viperInstance.ReadInConfig()

	if err != nil {
		return nil, err
	}

	sub := viperInstance.Sub(env)
	envSubst(sub)
	err = sub.Unmarshal(&configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func envSubst(sub *viper.Viper) {
	for _, k := range sub.AllKeys() {
		value := sub.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			sub.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}

func CreateConfigInstance() *config {
	return &config{}
}

func getViperInstance() *viper.Viper {
	viperInstance := viper.New()
	viperInstance.SetConfigFile("resources/config.yml")
	return viperInstance
}

func getGoEnv() string {
	env := os.Getenv("GO_ENV")
	if env != "" {
		return env
	}
	return "stage"
}
