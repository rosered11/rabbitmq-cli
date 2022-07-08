package rabbitctl

import (
	"rabbitmq-cli/pkg/rabbitmq"

	"github.com/spf13/cobra"
)

var receiveCmd = &cobra.Command{
	Use:     "receive",
	Aliases: []string{"rec"},
	Short:   "Receive data from rabbitmq",
	Run: func(cmd *cobra.Command, args []string) {
		rabbitmq.Consume()
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}
