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
	ModuleTypeDatabase ModuleType = "database"
	ModuleTypeCache    ModuleType = "cache"
	ModuleTypeLogging  ModuleType = "logging"
	ModuleTypeMQ       ModuleType = "mq"
	ModuleTypeCloud    ModuleType = "cloud"
)

// IsEnabled checks if the module is enabled
func (m *Module) IsEnabled() bool {
	return m.Enabled
}

// Enable enables the module
func (m *Module) Enable() {
	m.Enabled = true
}

// Disable disables the module
func (m *Module) Disable() {
	m.Enabled = false
}

// HasDependency checks if the module has a specific dependency
func (m *Module) HasDependency(name string) bool {
	for _, dep := range m.Dependencies {
		if dep == name {
			return true
		}
	}
	return false
}

// AddDependency adds a dependency to the module
func (m *Module) AddDependency(name string) {
	if !m.HasDependency(name) {
		m.Dependencies = append(m.Dependencies, name)
	}
}

// RemoveDependency removes a dependency from the module
func (m *Module) RemoveDependency(name string) {
	for i, dep := range m.Dependencies {
		if dep == name {
			m.Dependencies = append(m.Dependencies[:i], m.Dependencies[i+1:]...)
			return
		}
	}
}
