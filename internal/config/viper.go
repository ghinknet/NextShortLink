package config

import (
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/ghinknet/json"
	"github.com/spf13/viper"
)

var C *viper.Viper
var OnChange []func()
var lastLoad time.Time
var Debug = false
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

	// Is debug mode?
	if _, err = os.Stat("config_debug.yaml"); err == nil {
		// Init config file
		cfg.SetConfigName("config_debug")

		// Set debug status
		Debug = true

		// Read the debug config file
		err = cfg.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	// Watch config change
	cfg.WatchConfig()

	// Record first load
	lastLoad = time.Now()

	// Trigger to reload
	cfg.OnConfigChange(func(e fsnotify.Event) {
		// Debounce
		if lastLoad.Add(time.Second * 1).After(time.Now()) {
			return
		}

		for _, fn := range OnChange {
			fn()
		}

		lastLoad = time.Now()
	})

	// Replace default value
	if cfg.GetString("host") == "" {
		cfg.Set("host", "[::]")
	}

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
