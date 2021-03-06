package config_test

import (
	"testing"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := config.New([]string{".", "../../.."})
	assert.Equal(t, "3000", conf.Port())
	assert.Equal(t, "postgres://tanker:tanker@localhost:5436/tanker?sslmode=disable", conf.ConnectionURL())
}
