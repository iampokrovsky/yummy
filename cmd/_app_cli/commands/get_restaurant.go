package commands

import (
	"github.com/spf13/cobra"
	restmodel "yummy/internal/app/_restaurant/model"
)

func (cli *CLI) getRestaurantCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "_restaurant",
		Short: "Get _restaurant by ID",
		Run: func(cmd *cobra.Command, args []string) {
			restaurant, err := cli.coreService.GetRestaurantByID(cmd.Context(), restmodel.ID(id))
			if err != nil {
				cmd.PrintErrf("failed to get _restaurant: %v\n", err)
				return
			}

			cmd.Printf("Restaurant:\n%+v\n", restaurant)
		},
	}

	cmd.Flags().Int64VarP(&id, "id", "i", 0, "Restaurant ID")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		return nil
	}

	return cmd
}