package config

import (
	"github.com/gookit/color"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/sreioi/framework/contracts/config"
	"github.com/sreioi/framework/support/file"
	"os"
)

var _ config.Config = &Application{}

type Application struct {
	vip *viper.Viper
}

func NewApplication(envPath string) *Application {
	app := &Application{}
	app.vip = viper.New()
	// 读取环境变量
	app.vip.AutomaticEnv()

	if file.Exists(envPath) {
		// 设置配置文件类型 支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
		app.vip.SetConfigType("env")
		// 设置env相对文件路径
		app.vip.SetConfigFile(envPath)

		if err := app.vip.ReadInConfig(); err != nil {
			color.Redln("Invalid Config error: " + err.Error())
			os.Exit(0)
		}
	}

	return app
}

func (app *Application) Env(name string, defaultValue ...any) any {
	return app.Get(name, defaultValue...)
}

func (app *Application) Add(name string, configuration any) {
	app.vip.Set(name, configuration)
}

func (app *Application) Get(key string, defaultValue ...any) any {
	if !app.vip.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return app.vip.Get(key)
}

func (app *Application) GetString(key string, defaultValue ...string) string {
	return cast.ToString(app.Get(key, defaultValue))
}

func (app *Application) GetInt(key string, defaultValue ...any) int {
	return cast.ToInt(app.Get(key, defaultValue))
}

func (app *Application) GetBool(key string, defaultValue ...any) bool {
	return cast.ToBool(app.Get(key, defaultValue))
}

func (app *Application) GetAllKeys() []string {
	return app.vip.AllKeys()
}
