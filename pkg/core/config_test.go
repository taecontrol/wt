package core_test

import (
	"testing"
	"wt/pkg/core"

	"gopkg.in/yaml.v3"
)

func TestDefaultConfigPath(t *testing.T) {
	t.Run("It returns the absolute path to the .wt file", func(t *testing.T) {
		AbsHandlerMock := func(path string) (string, error) {
			return "/absolute/path/.wt", nil
		}

		config := &core.Config{
			AbsolutePath: AbsHandlerMock,
		}

		defaultPath := config.DefaultConfigPath()

		if defaultPath != "/absolute/path/.wt" {
			t.Error("DefaultConfigPath should return the absolute path to the .wt file")
		}
	})
}

func TestLoadConfiguration(t *testing.T) {
	t.Run("It loads the configuration from the .wt file", func(t *testing.T) {
		config := &core.Config{
			ParseYaml: func(b []byte, i interface{}) error {
				return yaml.Unmarshal(b, i)
			},
			ReadFile: func(s string) ([]byte, error) {
				return []byte("init_commands:\n  - echo 'hello world'\nterminate_commands:\n  - echo 'bye'"), nil
			},
			AbsolutePath: func(path string) (string, error) {
				return path, nil
			},
		}

		err := config.LoadConfig()

		if len(config.InitCommands) != 1 && config.InitCommands[0] != "echo 'hello world'" {
			t.Errorf("LoadConfiguration should load the init_commands from the .wt file: %v", err)
		}

		if len(config.TerminateCommands) != 1 && config.TerminateCommands[0] != "echo 'bye'" {
			t.Errorf("LoadConfiguration should load the terminate_commands from the .wt file: %v", err)
		}
	})
}
