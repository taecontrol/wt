package core

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ConfigContract interface {
	GetInitCommands() []string
	GetTerminateCommands() []string
	LoadConfig() error
	DefaultConfigPath() string
}

type Config struct {
	InitCommands      []string `yaml:"init_commands"`
	TerminateCommands []string `yaml:"terminate_commands"`

	AbsolutePath func(string) (string, error)
	ReadFile     func(string) ([]byte, error)
	ParseYaml    func([]byte, interface{}) error
}

func NewConfig() *Config {
	return &Config{
		AbsolutePath: filepath.Abs,
		ReadFile:     os.ReadFile,
		ParseYaml:    yaml.Unmarshal,
	}
}

func (c Config) GetInitCommands() []string {
	return c.InitCommands
}

func (c Config) GetTerminateCommands() []string {
	return c.TerminateCommands
}

func (c *Config) LoadConfig() error {
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
