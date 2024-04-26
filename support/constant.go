package support

const Version = "v1.0.0"

const (
	EnvTest    = "test"
	EnvRuntime = "runtime"
	EnvArtisan = "artisan"
)

var (
	Env          = EnvRuntime
	EnvPath      = ".env"
	RelativePath string
	RootPath     string
)
