package registry

type RegistryProvider interface {
	LoadRegistry() (*ModuleRegistry, error)
	FindModuleMetaData(name string) (*ModuleMetaData, error)
	ListAllMetaData() ([]*ModuleMetaData, error)
}
