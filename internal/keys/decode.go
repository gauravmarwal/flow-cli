package keys

import (
	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowcli/services"
)

type flagsDecode struct{}

var decodeFlags = flagsDecode{}

var DecodeCommand = &command.Command{
	Cmd: &cobra.Command{
		Use:     "decode <public key>",
		Short:   "Decode a public account key hex string",
		Args:    cobra.ExactArgs(1),
		Example: "flow keys decode 4a22246...31bce1e71a7b6d11",
	},
	Flags: &decodeFlags,
	Run: func(
		cmd *cobra.Command,
		args []string,
		globalFlags command.GlobalFlags,
		services *services.Services,
	) (command.Result, error) {
		accountKey, err := services.Keys.Decode(
			args[0], // public key
		)
		if err != nil {
			return nil, err
		}

		pubKey := accountKey.PublicKey
		return &KeyResult{publicKey: &pubKey, accountKey: accountKey}, err
	},
}
