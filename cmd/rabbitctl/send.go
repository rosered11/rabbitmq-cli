package rabbitctl

import (
	"rabbitmq-cli/pkg/rabbitmq"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"send"},
	Short:   "Send data to rabbitmq",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// res := stringer.Reverse(args[0])

		rabbitmq.Send(args[0])
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
