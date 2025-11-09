package config

import (
	"os"

	"github.com/ghinknet/json"
	"github.com/spf13/viper"
)

var C *viper.Viper
var Field = make(map[rune]int64)

// staticConfig is constructor of static config
func staticConfig() *viper.Viper {
	// Init viper
	cfg := viper.New()

	// Set config type
	cfg.SetConfigType("yaml")

	// Set config path
	cfg.AddConfigPath("./")

	// Set config file
	cfg.SetConfigName("config")

	// Read the config file
	err := cfg.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if cfg.GetBool("debug") {
		// Init config file
		cfg.SetConfigName("config_debug")

		// Read the debug config file
		err = cfg.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	// Watch config change
	cfg.WatchConfig()

	return cfg
}

// LoadField loads fieldMap from config
func LoadField() {
	// Read field
	var data []byte
	var err error
	if C.GetBool("debug") {
		data, err = os.ReadFile("field_debug.json")
	} else {
		data, err = os.ReadFile("field.json")
	}
	if err != nil {
		panic(err)
	}

	var resultMap map[string]int64

	// Parse json
	if err = json.Unmarshal(data, &resultMap); err != nil {
		panic(err)
	}

	// Fill field
	for k, v := range resultMap {
		Field[[]rune(k)[0]] = v
	}
}

// LoadStatic loads static config
func LoadStatic() *viper.Viper {
	C = staticConfig()
	LoadField()
	return C
}
