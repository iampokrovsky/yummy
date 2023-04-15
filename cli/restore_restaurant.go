package cli

import (
	"github.com/spf13/cobra"
	restmodel "yummy/internal/app/restaurant/model"
)

func (cli *CLI) restoreRestaurantCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "restaurant [flags]",
		Short: "Restore restaurant by ID",
		Run: func(cmd *cobra.Command, args []string) {
			ok, err := cli.restaurantService.Restore(cmd.Context(), restmodel.ID(id))
			if !ok || err != nil {
				cmd.PrintErrf("failed to restore restaurant: %v\n", err)
			}

			cmd.Println("Restaurant successfully restored.")
		},
	}

	cmd.Flags().Int64VarP(&id, "id", "i", 0, "Restaurant ID")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		return nil
	}

	return cmd
}
