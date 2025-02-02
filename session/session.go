package session

type Session struct {
	Config  *Config
	Handler Handlers
	option  Options
}

// TODO: add extra field and methods
type Options struct{}

// TODO: add extra field and methods
type Config struct {
	Region string
}

// TODO: add extra field and methods
type Handlers struct{}

func NewSession(cfgs *Config) (Session, error) {
	session := Session{}
	session.Config = cfgs
	return session, nil
}
