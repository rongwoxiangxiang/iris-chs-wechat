package qq

import (
	"chs/config"
	"github.com/pelletier/go-toml"
)

type Configuration struct {
	app_id  string
	app_key string
}

func (configuration Configuration) ToMap() map[string]string {
	return map[string]string{
		"app_id":  configuration.app_id,
		"app_key": configuration.app_key,
	}
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		app_id:  "xxxxxx",
		app_key: "xxxxxx",
	}
}

func TomlBytes(data []byte) *Configuration {
	toml, err := toml.LoadBytes(data)
	if err != nil {
		config.Logger().Errorf("QQ ai Configuration with bytes toml err: ", err)
		return nil
	}
	return TomlConfiguration(toml)
}

func TomlFileToml(filePath string) *Configuration {
	toml, err := toml.LoadFile(filePath)
	if err != nil {
		config.Logger().Errorf("QQ ai Configuration with bytes toml err: ", err)
		return nil
	}
	return TomlConfiguration(toml)
}

func TomlConfiguration(tree *toml.Tree) *Configuration {
	app_id := tree.Get("qq.ai.app_id")
	app_key := tree.Get("qq.ai.app_key")
	if app_id == nil || app_key == nil {
		config.Logger().Errorf("QQ ai Configuration with toml tree empty!")
		return nil
	}
	return &Configuration{
		app_id:  app_id.(string),
		app_key: app_key.(string),
	}
}

func SetConfiguration(app_id, app_key string) *Configuration {
	if app_id == "" || app_key == "" {
		config.Logger().Errorf("QQ ai Configuration with toml tree empty!")
		return nil
	}
	return &Configuration{
		app_id:  app_id,
		app_key: app_key,
	}
}
