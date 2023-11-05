package cmd

import (
	"fmt"
	"strings"

	"github.com/sivaosorg/govm/logger"
)

func NewCommandManager() *CommandManager {
	return &CommandManager{
		commands: make(map[string]Command),
	}
}

func (c *CommandManager) AddCommand(value Command) {
	c.commands[value.Name()] = value
}

func (c *CommandManager) AddCommands(values ...Command) {
	for _, v := range values {
		c.AddCommand(v)
	}
}

func (c *CommandManager) Size() int {
	return len(c.commands)
}

func (c *CommandManager) AvailableCommand() bool {
	return c.Size() > 0
}

func (c *CommandManager) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("No command provided")
	}
	name := args[1]
	command, ok := c.commands[name]
	if !ok {
		return fmt.Errorf("Unknown command: %s", name)
	}
	_args := args[2:]
	logger.Infof(fmt.Sprintf("[App] Running command: %s", command.Name()))
	logger.Infof(fmt.Sprintf("[App] Command description: %s", command.Description()))
	logger.Infof(fmt.Sprintf("[App] Command args: %s", strings.Join(_args, ",")))
	return command.Execute(_args)
}
