package wt

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Command struct {
}

type Config struct {
	InitCommands      []string `yaml:"init_commands"`
	TerminateCommands []string `yaml:"terminate_commands"`

	AbsolutePath func(string) (string, error)
	ReadFile     func(string) ([]byte, error)
	ParseYaml    func([]byte, interface{}) error
}

func NewConfig() (*Config, error) {
	c := &Config{
		AbsolutePath: filepath.Abs,
		ReadFile:     os.ReadFile,
		ParseYaml:    yaml.Unmarshal,
	}

	err := c.LoadConfiguration()

	return c, err
}

func (c *Config) LoadConfiguration() error {
	fileData, err := c.ReadFile(c.DefaultConfigPath())
	if err != nil {
		return err
	}

	err = c.ParseYaml(fileData, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c Config) DefaultConfigPath() string {
	configPath, err := c.AbsolutePath("./.wt")

	if err != nil {
		panic(err)
	}

	return configPath
}
