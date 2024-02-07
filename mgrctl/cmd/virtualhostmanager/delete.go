package virtualhostmanager

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uyuni-project/uyuni-tools/shared/api"
	apiTypes "github.com/uyuni-project/uyuni-tools/shared/api/types"
	"github.com/uyuni-project/uyuni-tools/shared/api/virtualhostmanager"
	"github.com/uyuni-project/uyuni-tools/shared/types"
	"github.com/uyuni-project/uyuni-tools/shared/utils"
)

type deleteFlags struct {
	api.ConnectionDetails `mapstructure:"api"`
	Label                 string
}

func deleteCommand(globalFlags *types.GlobalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Virtual Host Manager with a given label",
		RunE: func(cmd *cobra.Command, args []string) error {
			var flags deleteFlags
			return utils.CommandHelper(globalFlags, cmd, args, &flags, delete)
		},
	}

	cmd.Flags().String("Label", "", "Virtual Host Manager label")

	return cmd
}

func delete(globalFlags *types.GlobalFlags, flags *deleteFlags, cmd *cobra.Command, args []string) error {

	res, err := virtualhostmanager.Virtualhostmanager(&flags.ConnectionDetails, flags.Label)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %v", res)

	return nil
}
