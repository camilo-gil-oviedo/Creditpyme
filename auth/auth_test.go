package auth

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// fakeTokenCreator implementa TokenCreator para tests
type fakeTokenCreator struct{}

func (f *fakeTokenCreator) CustomToken(ctx context.Context, uid string) (string, error) {
	return "token-for-" + uid, nil
}

func setupInMemoryDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)

	require.NoError(t, db.AutoMigrate(&User{}))
	return db
}

func TestRegisterAndLogin(t *testing.T) {
	db := setupInMemoryDB(t)
	svc := &AuthService{DB: db, FireAuth: &fakeTokenCreator{}}

	// Registrar
	token, err := svc.Register("alice@example.com", "password123")
	require.NoError(t, err)
	require.Contains(t, token, "token-for-")

	// Login
	loginToken, err := svc.Login("alice@example.com", "password123")
	require.NoError(t, err)
	require.Contains(t, loginToken, "token-for-")
}

func TestRegisterDuplicate(t *testing.T) {
	db := setupInMemoryDB(t)
	svc := &AuthService{DB: db, FireAuth: &fakeTokenCreator{}}

	_, err := svc.Register("bob@example.com", "password123")
	require.NoError(t, err)

	_, err = svc.Register("bob@example.com", "password123")
	require.Error(t, err)
}
