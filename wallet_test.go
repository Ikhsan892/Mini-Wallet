package wallet_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet"
)

func TestDBIsNotNIl(t *testing.T) {
	g := wallet.New()
	g.Start()
	assert.NotNil(t, g.DB())
}

func TestSettingsIsNotNull(t *testing.T) {
	g := wallet.New()

	assert.NotEmpty(t, g.Settings().DB.Driver)
}
