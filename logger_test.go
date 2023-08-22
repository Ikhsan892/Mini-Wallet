package wallet_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZapLogger(t *testing.T) {
	assert.NotNil(t, intellisearchdoc.NewZapLogger())
}
