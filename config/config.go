package config

import (
	"github.com/pelletier/go-toml"
	"io"
	"os"
)

var Config struct {
	General struct {
		Address string `toml:"address"`
	}
	Database struct {
		Address  string `toml:"address"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}
}

func init() {
	if file, err := os.Open("config.toml"); err == nil {
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		bytes, _ := io.ReadAll(file)

		err = toml.Unmarshal(bytes, &Config)
		if err != nil {
			panic(err)
		}
	} else {
		data, _ := toml.Marshal(Config)

		err = os.WriteFile("config.toml", data, 0644)
		if err != nil {
			panic(err)
		}
	}
}
