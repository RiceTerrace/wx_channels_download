package config

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func findRegistryItem(key string) (ConfigItem, bool) {
	for _, item := range Registry {
		if item.Key == key {
			return item, true
		}
	}
	return ConfigItem{}, false
}

func TestLoadConfigDefaultsToOriginalVideoDownload(t *testing.T) {
	Registry = nil
	viper.Reset()
	t.Cleanup(func() {
		Registry = nil
		viper.Reset()
	})

	cfg := &Config{}
	if err := cfg.LoadConfig(); err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	item, ok := findRegistryItem("download.defaultHighest")
	if !ok {
		t.Fatal("download.defaultHighest not registered")
	}

	defaultValue, ok := item.Default.(bool)
	if !ok {
		t.Fatalf("download.defaultHighest default is not bool: %T", item.Default)
	}
	if !defaultValue {
		t.Fatalf("download.defaultHighest default = %v, want true", defaultValue)
	}
	if !viper.GetBool("download.defaultHighest") {
		t.Fatal("viper default for download.defaultHighest = false, want true")
	}
}

func TestConfigTemplateDefaultsToOriginalVideoDownload(t *testing.T) {
	data, err := os.ReadFile("config.template.yaml")
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if !strings.Contains(string(data), "defaultHighest: true") {
		t.Fatal("config.template.yaml does not default download.defaultHighest to true")
	}
}
