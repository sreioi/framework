package config

type Config interface {
	// Env get config from env.
	Env(name string, defaultValue ...any) any
	// Add config to application.
	Add(name string, configuration any)
	// Get config from application.
	Get(key string, defaultValue ...any) any
	// GetString get string type config from application.
	GetString(key string, defaultValue ...string) string
	// GetInt get int type config from application.
	GetInt(key string, defaultValue ...any) int
	// GetBool get bool type config from application.
	GetBool(key string, defaultValue ...any) bool
	// GetAllKeys get all keys from application.
	GetAllKeys() []string
}
