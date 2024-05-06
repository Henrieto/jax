package jax

import "github.com/henrieto/jax/command"

func LoadConfig(config *Config) {
	for _, plugin := range config.Plugins {
		RegisterCommand(plugin.Commands...)
	}
}

func RegisterCommand(cmds ...*command.Command) {
	command.Register(cmds...)
}
