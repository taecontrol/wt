package utils

import "os"

type Env struct {
	envVars map[string]string
}

func NewEnv() *Env {
	return &Env{
		envVars: make(map[string]string),
	}
}

func (env *Env) Set(key string, value string) {
	env.envVars[key] = value
}

func (env *Env) SetEnvVars(envVars map[string]string) {
	env.envVars = envVars
}

func (env *Env) Get(key string) string {
	return env.envVars[key]
}

func (env *Env) ToEnviron() []string {
	environ := os.Environ()

	for key, value := range env.envVars {
		environ = append(environ, key+"="+value)
	}

	return environ
}
