package wallet_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMigrator(t *testing.T) {

	app := intellisearchdoc.New()

	m := app.Migration()
	assert.NotNil(t, m)
}
