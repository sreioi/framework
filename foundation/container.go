package foundation

import (
	"fmt"
	"github.com/sreioi/framework/contracts/foundation"
	"sync"
)

type Container struct {
	bindings  sync.Map
	instances sync.Map
}

type instance struct {
	concrete any
	shared   bool
}

func NewContainer() *Container {
	return &Container{}
}

// Bind 绑定
func (c *Container) Bind(key any, callback func(app foundation.Application) (any, error)) {
	c.bindings.Store(key, instance{
		concrete: callback,
		shared:   false,
	})
}

// BindWith 绑定并且附带参数
func (c *Container) BindWith(key any, callback func(app foundation.Application, parameters map[string]any) (any, error)) {
	c.bindings.Store(key, instance{
		concrete: callback,
		shared:   false,
	})
}

// Instance 实例
func (c *Container) Instance(key any, ins any) {
	c.bindings.Store(key, instance{concrete: ins, shared: true})
}

// Singleton 单例
func (c *Container) Singleton(key any, callback func(app foundation.Application) (any, error)) {
	c.bindings.Store(key, instance{
		concrete: callback,
		shared:   true,
	})
}

// Make 获取示例
func (c *Container) Make(key any) (any, error) {
	return c.make(key, nil)
}

// MakeWith 获取示例并且附带参数
func (c *Container) MakeWith(key any, parameters map[string]any) (any, error) {
	return c.make(key, parameters)
}

// make 获取示例，如果没有示例则初始化并保存
func (c *Container) make(key any, parameters map[string]any) (any, error) {
	binding, ok := c.bindings.Load(key)
	if !ok {
		return nil, fmt.Errorf("binding not found: %+v", key)
	}

	if parameters == nil {
		instance, ok := c.instances.Load(key)
		if ok {
			return instance, nil
		}
	}

	bindingImpl := binding.(instance)
	switch concrete := bindingImpl.concrete.(type) {
	case func(app foundation.Application) (any, error):
		concreteImpl, err := concrete(App)
		if err != nil {
			return nil, err
		}
		if bindingImpl.shared {
			c.instances.Store(key, concreteImpl)
		}

		return concreteImpl, nil
	case func(app foundation.Application, parameters map[string]any) (any, error):
		concreteImpl, err := concrete(App, parameters)
		if err != nil {
			return nil, err
		}

		return concreteImpl, nil
	default:
		c.instances.Store(key, concrete)
		return concrete, nil
	}
}
