package foundation

import (
	carbon2 "github.com/golang-module/carbon/v2"
	"github.com/sreioi/framework/config"
	"github.com/sreioi/framework/contracts/foundation"
	"github.com/sreioi/framework/support/carbon"
)

var App foundation.Application

type Application struct {
	foundation.Container
}

func init() {
	// 加载配置文件
	app := &Application{
		Container: NewContainer(),
	}
	// 注册框架的基本服务
	app.registerBaseServiceProviders()
	// 实现基础的日志服务
	app.bootBaseServiceProviders()
	// 实现数据库服务
	// 启动框架的基本服务
	// 返回框架的App实例
	App = app
}

func NewApplication() foundation.Application {
	return App
}

func (app *Application) Boot() {
	app.registerConfiguredServiceProviders()
	app.bootConfiguredServiceProviders()
	app.setTimezone()
}

// getBaseServiceProviders 获取框架的基本服务提供者
func (app *Application) getBaseServiceProviders() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&config.ServiceProvider{},
	}
}

// getConfiguredServiceProviders 获取配置的服务提供者
func (app *Application) getConfiguredServiceProviders() []foundation.ServiceProvider {
	return app.MakeConfig().Get("app.providers").([]foundation.ServiceProvider)
}

// registerConfiguredServiceProviders 注册配置的服务提供者
func (app *Application) registerConfiguredServiceProviders() {
	app.registerServiceProviders(app.getConfiguredServiceProviders())
}

// bootConfiguredServiceProviders 启动配置的服务提供者
func (app *Application) bootConfiguredServiceProviders() {
	app.bootServiceProviders(app.getConfiguredServiceProviders())
}

// registerBaseServiceProviders 注册框架的基本服务
func (app *Application) registerBaseServiceProviders() {
	app.registerServiceProviders(app.getBaseServiceProviders())
}

// bootBaseServiceProviders 启动框架的基本服务
func (app *Application) bootBaseServiceProviders() {
	app.bootServiceProviders(app.getBaseServiceProviders())
}

// registerServiceProviders 注册服务提供者
func (app *Application) registerServiceProviders(serviceProviders []foundation.ServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Register(app)
	}
}

// bootServiceProviders 启动服务提供者
func (app *Application) bootServiceProviders(serviceProviders []foundation.ServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Boot(app)
	}
}

func (app *Application) setTimezone() {
	carbon.NewTime().SetTimezone(app.MakeConfig().GetString("app.timezone", carbon2.Shanghai))
}
