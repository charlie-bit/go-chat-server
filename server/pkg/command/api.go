package command

import "github.com/spf13/cobra"

type apiCommand struct {
	baseCommand
	run func(port int) error
}

func NewApiCommand(run func(port int) error) ServerCommander {
	return &apiCommand{
		newBaseCommand("api"),
		run,
	}
}

func (a apiCommand) AddRunE() {
	a.command.RunE = func(cmd *cobra.Command, args []string) error {
		return a.run(a.getWsPort(cmd))
	}
}

func (a apiCommand) Execute() error {
	a.AddRunE()
	return a.Execute()
}
