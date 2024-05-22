package foundation

import (
	"github.com/sreioi/framework/config"
	"github.com/sreioi/framework/contracts/foundation"
	"github.com/sreioi/framework/support/carbon"
	"path/filepath"
)

var App foundation.Application

type Application struct {
	foundation.Container
	publishes     map[string]map[string]string
	publishGroups map[string]map[string]string
}

func init() {
	// 加载配置文件
	app := &Application{
		Container:     NewContainer(),
		publishes:     make(map[string]map[string]string),
		publishGroups: make(map[string]map[string]string),
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

// setTimezone 设置时区
func (app *Application) setTimezone() {
	carbon.SetTimezone(app.MakeConfig().GetString("app.timezone", carbon.Shanghai))
}

func (app *Application) Path(path string) string {
	return filepath.Join("app", path)
}

func (app *Application) BasePath(path string) string {
	return filepath.Join("", path)
}

func (app *Application) ConfigPath(path string) string {
	return filepath.Join("config", path)
}

func (app *Application) DatabasePath(path string) string {
	return filepath.Join("database", path)
}

func (app *Application) StoragePath(path string) string {
	return filepath.Join("storage", path)
}

func (app *Application) PublicPath(path string) string {
	return filepath.Join("public", path)
}

func (app *Application) Publishes(packageName string, paths map[string]string, groups ...string) {
	app.ensurePublishArrayInitialized(packageName)

	for key, value := range paths {
		app.publishes[packageName][key] = value
	}

	for _, group := range groups {
		app.addPublishGroup(group, paths)
	}
}

func (app *Application) ensurePublishArrayInitialized(packageName string) {
	if _, exist := app.publishes[packageName]; !exist {
		app.publishes[packageName] = make(map[string]string)
	}
}

func (app *Application) addPublishGroup(group string, paths map[string]string) {
	if _, exist := app.publishGroups[group]; !exist {
		app.publishGroups[group] = make(map[string]string)
	}

	for key, value := range paths {
		app.publishGroups[group][key] = value
	}
}
