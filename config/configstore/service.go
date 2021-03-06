package configstore

// Service exposes functions over the config objects
type Service interface {
	GetConfig() (*NodeConfig, error)
	GetTenant(identifier []byte) (*TenantConfig, error)
	GetAllTenants() ([]*TenantConfig, error)
	CreateConfig(data *NodeConfig) (*NodeConfig, error)
	CreateTenant(data *TenantConfig) (*TenantConfig, error)
	UpdateConfig(data *NodeConfig) (*NodeConfig, error)
	UpdateTenant(data *TenantConfig) (*TenantConfig, error)
	DeleteConfig() error
	DeleteTenant(identifier []byte) error
}

type service struct {
	repo Repository
}

// DefaultService returns an implementation of the config.Service
func DefaultService(repository Repository) Service {
	return &service{repo: repository}
}

func (s service) GetConfig() (*NodeConfig, error) {
	return s.repo.GetConfig()
}

func (s service) GetTenant(identifier []byte) (*TenantConfig, error) {
	return s.repo.GetTenant(identifier)
}

func (s service) GetAllTenants() ([]*TenantConfig, error) {
	return s.repo.GetAllTenants()
}

func (s service) CreateConfig(data *NodeConfig) (*NodeConfig, error) {
	return data, s.repo.CreateConfig(data)
}

func (s service) CreateTenant(data *TenantConfig) (*TenantConfig, error) {
	return data, s.repo.CreateTenant(data.IdentityID, data)
}

func (s service) UpdateConfig(data *NodeConfig) (*NodeConfig, error) {
	return data, s.repo.UpdateConfig(data)
}

func (s service) UpdateTenant(data *TenantConfig) (*TenantConfig, error) {
	return data, s.repo.UpdateTenant(data.IdentityID, data)
}

func (s service) DeleteConfig() error {
	return s.repo.DeleteConfig()
}

func (s service) DeleteTenant(identifier []byte) error {
	return s.repo.DeleteTenant(identifier)
}
