package configs

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {

	box := packr.NewBox("./")
	configType := "yml"
	defaultConfig, _ := box.Find("default.yml")
	v := viper.New()
	v.SetConfigType(configType)

	err := v.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		return
	}
	configs := v.AllSettings()

	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	env := os.Getenv("ENV")

	if env != "" {
		envConfig, _ := box.Find(env + ".yml")
		viper.SetConfigType(configType)
		err = viper.ReadConfig(bytes.NewReader(envConfig))
		if err != nil {
			return
		}
	}
	viper.AutomaticEnv()
}
