package command

import (
	"chat_socket/server/pkg/constant"
	"github.com/spf13/cobra"
)

type baseCommand struct {
	command    cobra.Command
	serverName string
	wsPort     int
	grpcPort   int
}

// new base command, init config and log
func newBaseCommand(name string) baseCommand {
	return baseCommand{
		serverName: name,
		command: cobra.Command{
			Use: "pre init config",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				// init config
				path := ""
				return nil
			},
		},
	}
}

func (b baseCommand) addServerName(name string) {
	b.command.Flags().StringP(constant.ServerName, "server_name", name, "server name")
}

func (b baseCommand) addWsPort(port int) {
	b.command.Flags().IntP(constant.ServerWsPort, "ws_port", port, "ws port")
}

func (b baseCommand) addGrpcPost(port int) {
	b.command.Flags().IntP(constant.ServerGrpcPort, "grpc_port", port, "grpc port")
}

func (b baseCommand) getServerName() string {
	name, _ := b.command.Flags().GetString(constant.ServerName)
	return name
}

func (b baseCommand) getWsPort(cmd *cobra.Command) int {
	port, _ := cmd.Flags().GetInt(constant.ServerWsPort)
	return port
}

func (b baseCommand) getGrpcPort(cmd *cobra.Command) int {
	port, _ := cmd.Flags().GetInt(constant.ServerGrpcPort)
	return port
}

func (b baseCommand) addConfigPath(path string) {
	b.command.Flags().StringP(constant.ServerConfigPath, "config_path", path, "config file path")
}

func (b baseCommand) getConfigPath(cmd *cobra.Command) string {
	path, _ := cmd.Flags().GetString(constant.ServerConfigPath)
	return path
}
