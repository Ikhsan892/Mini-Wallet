package wallet

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	godotenv.Load()
	cfg := newConfig()

	assert.NotNil(t, cfg.DB.Driver)
}
