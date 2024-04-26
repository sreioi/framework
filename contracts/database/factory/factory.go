package factory

type Factory interface {
	// Definition 定义了模型的默认状态。
	Definition() map[string]any
}

type Model interface {
	// Factory 为模型创建一个新的工厂实例。
	Factory() Factory
}
