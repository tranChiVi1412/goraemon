package module

type Module struct {
	Name         string                 // Name fo modude(ex: redis, mysql, etc)
	Type         ModuleType             // Type of module(ex: db, cache, logging, etc)
	Version      string                 // Version of module
	Config       map[string]interface{} // Config of module
	Dependencies []string               // Dependencies of module
	Enabled      bool                   // Enabled of module
}

type ModuleType string

const (
	ModuleTypeDB      ModuleType = "db"
	ModuleTypeCache   ModuleType = "cache"
	ModuleTypeLogging ModuleType = "logging"
	ModuleTypeMQ      ModuleType = "mq"
	ModuleTypeCloud   ModuleType = "cloud"
)
