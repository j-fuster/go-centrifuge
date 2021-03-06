// +build unit

package configstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetConfig_NoConfig(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterConfig(&NodeConfig{})
	svc := DefaultService(repo)
	cfg, err := svc.GetConfig()
	assert.NotNil(t, err)
	assert.Nil(t, cfg)
}

func TestService_GetConfig(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterConfig(&NodeConfig{})
	svc := DefaultService(repo)
	nodeCfg := NewNodeConfig(cfg)
	err = repo.CreateConfig(nodeCfg)
	assert.Nil(t, err)
	cfg, err := svc.GetConfig()
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
}

func TestService_GetTenant_NoTenant(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterTenant(&TenantConfig{})
	svc := DefaultService(repo)
	cfg, err := svc.GetTenant([]byte("0x123456789"))
	assert.NotNil(t, err)
	assert.Nil(t, cfg)
}

func TestService_GetTenant(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterTenant(&TenantConfig{})
	svc := DefaultService(repo)
	tenantCfg, err := NewTenantConfig("main", cfg)
	assert.Nil(t, err)
	err = repo.CreateTenant(tenantCfg.IdentityID, tenantCfg)
	assert.Nil(t, err)
	cfg, err := svc.GetTenant(tenantCfg.IdentityID)
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
}

func TestService_CreateConfig(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterConfig(&NodeConfig{})
	svc := DefaultService(repo)
	nodeCfg := NewNodeConfig(cfg)
	cfgpb, err := svc.CreateConfig(nodeCfg)
	assert.Nil(t, err)
	assert.Equal(t, nodeCfg.StoragePath, cfgpb.StoragePath)

	//Config already exists
	_, err = svc.CreateConfig(nodeCfg)
	assert.NotNil(t, err)
}

func TestService_CreateTenant(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterTenant(&TenantConfig{})
	svc := DefaultService(repo)
	tenantCfg, err := NewTenantConfig("main", cfg)
	assert.Nil(t, err)
	newCfg, err := svc.CreateTenant(tenantCfg)
	assert.Nil(t, err)
	assert.Equal(t, tenantCfg.IdentityID, newCfg.IdentityID)

	//Tenant already exists
	_, err = svc.CreateTenant(tenantCfg)
	assert.NotNil(t, err)
}

func TestService_UpdateConfig(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterConfig(&NodeConfig{})
	svc := DefaultService(repo)
	nodeCfg := NewNodeConfig(cfg)

	//Config doesn't exists
	_, err = svc.UpdateConfig(nodeCfg)
	assert.NotNil(t, err)

	newCfg, err := svc.CreateConfig(nodeCfg)
	assert.Nil(t, err)
	assert.Equal(t, nodeCfg.StoragePath, newCfg.StoragePath)

	nodeCfg.NetworkString = "something"
	newCfg, err = svc.UpdateConfig(nodeCfg)
	assert.Nil(t, err)
	assert.Equal(t, nodeCfg.NetworkString, newCfg.NetworkString)
}

func TestService_UpdateTenant(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterTenant(&TenantConfig{})
	svc := DefaultService(repo)
	tenantCfg, err := NewTenantConfig("main", cfg)

	// Tenant doesn't exist
	newCfg, err := svc.UpdateTenant(tenantCfg)
	assert.NotNil(t, err)

	newCfg, err = svc.CreateTenant(tenantCfg)
	assert.Nil(t, err)
	assert.Equal(t, tenantCfg.IdentityID, newCfg.IdentityID)

	tenantCfg.EthereumDefaultAccountName = "other"
	newCfg, err = svc.UpdateTenant(tenantCfg)
	assert.Nil(t, err)
	assert.Equal(t, tenantCfg.EthereumDefaultAccountName, newCfg.EthereumDefaultAccountName)
}

func TestService_DeleteConfig(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterConfig(&NodeConfig{})
	svc := DefaultService(repo)

	//No config, no error
	err = svc.DeleteConfig()
	assert.Nil(t, err)

	nodeCfg := NewNodeConfig(cfg)
	_, err = svc.CreateConfig(nodeCfg)
	assert.Nil(t, err)

	err = svc.DeleteConfig()
	assert.Nil(t, err)

	_, err = svc.GetConfig()
	assert.NotNil(t, err)
}

func TestService_DeleteTenant(t *testing.T) {
	repo, _, err := getRandomStorage()
	assert.Nil(t, err)
	repo.RegisterTenant(&TenantConfig{})
	svc := DefaultService(repo)
	tenantCfg, err := NewTenantConfig("main", cfg)
	assert.Nil(t, err)

	//No config, no error
	err = svc.DeleteTenant(tenantCfg.IdentityID)
	assert.Nil(t, err)

	_, err = svc.CreateTenant(tenantCfg)
	assert.Nil(t, err)

	err = svc.DeleteTenant(tenantCfg.IdentityID)
	assert.Nil(t, err)

	_, err = svc.GetTenant(tenantCfg.IdentityID)
	assert.NotNil(t, err)
}
