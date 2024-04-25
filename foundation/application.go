package foundation

import (
	"github.com/sreioi/framework/config"
	"github.com/sreioi/framework/contracts/foundation"
)

var App foundation.Application

type Application struct {
	foundation.Container
}

func init() {
	// 加载配置文件
	app := &Application{}
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

}

// getBaseServiceProviders 获取框架的基本服务提供者
func (app *Application) getBaseServiceProviders() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&config.ServiceProvider{},
	}
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
