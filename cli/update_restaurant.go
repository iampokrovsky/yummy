package cli

import (
	"github.com/spf13/cobra"
	restmodel "yummy/internal/app/restaurant/model"
)

func (cli *CLI) updateRestaurantCmd() *cobra.Command {
	var restRaw restmodel.Restaurant

	cmd := &cobra.Command{
		Use:   "restaurant",
		Short: "Update restaurant by ID",
		Run: func(cmd *cobra.Command, args []string) {
			ok, err := cli.restaurantService.Update(cmd.Context(), restRaw)
			if !ok || err != nil {
				cmd.PrintErrf("failed to update the restaurant: %v\n", err)
			}

			cmd.Println("Restaurant successfully updated.")
		},
	}

	cmd.Flags().Int64VarP((*int64)(&restRaw.ID), "id", "i", 0, "ID")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		return nil
	}
	cmd.Flags().StringVarP(&restRaw.Name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&restRaw.Address, "address", "a", "", "Address")
	cmd.Flags().StringVarP((*string)(&restRaw.Cuisine), "cuisine", "c", "", "Cuisine")

	return cmd
}
