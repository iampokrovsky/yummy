package cli

import (
	"github.com/spf13/cobra"
	restmodel "yummy/internal/app/restaurant/model"
)

func (cli *CLI) getRestaurantCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "restaurant",
		Short: "Get restaurant by ID",
		Run: func(cmd *cobra.Command, args []string) {
			restaurant, err := cli.restaurantService.GetByID(cmd.Context(), restmodel.ID(id))
			if err != nil {
				cmd.PrintErrf("failed to get restaurant: %v\n", err)
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
