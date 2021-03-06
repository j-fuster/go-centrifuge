// +build unit

package configstore

import (
	"os"
	"reflect"
	"testing"

	"github.com/centrifuge/go-centrifuge/config"

	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testlogging"

	"github.com/centrifuge/go-centrifuge/storage"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/stretchr/testify/assert"
)

var dbFiles []string
var ctx = map[string]interface{}{}
var cfg config.Configuration

func TestMain(m *testing.M) {
	ibootstappers := []bootstrap.TestBootstrapper{
		&testlogging.TestLoggingBootstrapper{},
		&config.Bootstrapper{},
		&storage.Bootstrapper{},
	}
	bootstrap.RunTestBootstrappers(ibootstappers, ctx)
	configdb := ctx[storage.BootstrappedConfigDB].(storage.Repository)
	cfg = ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	nc := NewNodeConfig(cfg)
	// clean db
	_ = configdb.Delete(getConfigKey())
	_ = configdb.Delete(getTenantKey(nc.MainIdentity.IdentityID))
	result := m.Run()
	cleanupDBFiles()
	os.Exit(result)
}

func TestNewLevelDBRepository(t *testing.T) {
	repo, _, _ := getRandomStorage()
	assert.NotNil(t, repo)
}

func TestUnregisteredModel(t *testing.T) {
	repo, _, _ := getRandomStorage()
	assert.NotNil(t, repo)
	id := utils.RandomSlice(32)
	newTenant := &TenantConfig{
		IdentityID:                 id,
		EthereumDefaultAccountName: "main",
	}
	err := repo.CreateTenant(id, newTenant)
	assert.Nil(t, err)

	// Error on non registered model
	_, err = repo.GetTenant(id)
	assert.NotNil(t, err)

	repo.RegisterTenant(&TenantConfig{})

	_, err = repo.GetTenant(id)
	assert.Nil(t, err)
}

func TestTenantOperations(t *testing.T) {
	repo, _, _ := getRandomStorage()
	assert.NotNil(t, repo)
	id := utils.RandomSlice(32)
	newTenant := &TenantConfig{
		IdentityID:                 id,
		EthereumDefaultAccountName: "main",
	}
	repo.RegisterTenant(&TenantConfig{})
	err := repo.CreateTenant(id, newTenant)
	assert.Nil(t, err)

	// Create tenant already exist
	err = repo.CreateTenant(id, newTenant)
	assert.NotNil(t, err)

	readTenant, err := repo.GetTenant(id)
	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(newTenant), readTenant.Type())
	assert.Equal(t, newTenant.IdentityID, readTenant.IdentityID)

	// Update tenant
	newTenant.EthereumDefaultAccountName = "secondary"
	err = repo.UpdateTenant(id, newTenant)
	assert.Nil(t, err)

	// Update tenant does not exist
	newId := utils.RandomSlice(32)
	err = repo.UpdateTenant(newId, newTenant)
	assert.NotNil(t, err)

	// Delete tenant
	err = repo.DeleteTenant(id)
	assert.Nil(t, err)
	_, err = repo.GetTenant(id)
	assert.NotNil(t, err)
}

func TestConfigOperations(t *testing.T) {
	repo, _, _ := getRandomStorage()
	assert.NotNil(t, repo)
	newConfig := &NodeConfig{
		NetworkID: 4,
	}
	repo.RegisterConfig(&NodeConfig{})
	err := repo.CreateConfig(newConfig)
	assert.Nil(t, err)

	// Create config already exist
	err = repo.CreateConfig(newConfig)
	assert.NotNil(t, err)

	readDoc, err := repo.GetConfig()
	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(newConfig), readDoc.Type())
	assert.Equal(t, newConfig.NetworkID, readDoc.NetworkID)

	// Update config
	newConfig.NetworkID = 42
	err = repo.UpdateConfig(newConfig)
	assert.Nil(t, err)

	// Delete config
	err = repo.DeleteConfig()
	assert.Nil(t, err)
	_, err = repo.GetConfig()
	assert.NotNil(t, err)

	// Update config does not exist
	err = repo.UpdateConfig(newConfig)
	assert.NotNil(t, err)
}

func TestLevelDBRepo_GetAllTenants(t *testing.T) {
	repo, _, _ := getRandomStorage()
	assert.NotNil(t, repo)
	repo.RegisterTenant(&TenantConfig{})
	ids := [][]byte{utils.RandomSlice(32), utils.RandomSlice(32), utils.RandomSlice(32)}
	ten1 := &TenantConfig{
		IdentityID:                 ids[0],
		EthereumDefaultAccountName: "main",
	}
	ten2 := &TenantConfig{
		IdentityID:                 ids[1],
		EthereumDefaultAccountName: "main",
	}
	ten3 := &TenantConfig{
		IdentityID:                 ids[2],
		EthereumDefaultAccountName: "main",
	}

	err := repo.CreateTenant(ids[0], ten1)
	assert.Nil(t, err)
	err = repo.CreateTenant(ids[1], ten2)
	assert.Nil(t, err)
	err = repo.CreateTenant(ids[2], ten3)
	assert.Nil(t, err)

	tenants, err := repo.GetAllTenants()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(tenants))
	t0Id := tenants[0].ID()
	t1Id := tenants[1].ID()
	t2Id := tenants[2].ID()
	assert.Contains(t, ids, t0Id)
	assert.Contains(t, ids, t1Id)
	assert.Contains(t, ids, t2Id)
}

func cleanupDBFiles() {
	for _, db := range dbFiles {
		err := os.RemoveAll(db)
		if err != nil {
			apiLog.Warningf("Cleanup warn: %v", err)
		}
	}
}

func getRandomStorage() (Repository, string, error) {
	randomPath := storage.GetRandomTestStoragePath()
	db, err := storage.NewLevelDBStorage(randomPath)
	if err != nil {
		return nil, "", err
	}
	dbFiles = append(dbFiles, randomPath)
	return NewDBRepository(storage.NewLevelDBRepository(db)), randomPath, nil
}
