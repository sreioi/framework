package foundation

type ServiceProvider interface {
	Register(app Application)
	Boot(app Application)
}
