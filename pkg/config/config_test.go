package config_test

import (
	"testing"

	"github.com/vgxbj/mailinone/pkg/config"
)

func TestEmailsConfig(t *testing.T) {
	var configs *config.Configs
	var err error

	if configs, err = config.ReadConfigFromFile("./config-test.toml"); err != nil {
		t.Errorf("%v", err)
	}

	if err = configs.Verify(); err != nil {
		t.Errorf("%v", err)
	}
}
