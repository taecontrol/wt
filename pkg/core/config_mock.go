package core

import "github.com/stretchr/testify/mock"

type ConfigMock struct {
	mock.Mock

	InitCommands      []string `yaml:"init_commands"`
	TerminateCommands []string `yaml:"terminate_commands"`
}

func NewConfigMock() *ConfigMock {
	return &ConfigMock{}
}

func (configMock *ConfigMock) GetInitCommands() []string {
	args := configMock.Called()
	return args.Get(0).([]string)
}

func (configMock *ConfigMock) GetTerminateCommands() []string {
	args := configMock.Called()
	return args.Get(0).([]string)
}

func (configMock *ConfigMock) LoadConfig() error {
	args := configMock.Called()

	return args.Error(0)
}

func (configMock *ConfigMock) DefaultConfigPath() string {
	args := configMock.Called()

	return args.String(0)
}
