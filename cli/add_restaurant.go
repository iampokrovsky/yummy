package cli

import (
	"github.com/spf13/cobra"
	restmodel "yummy/internal/app/restaurant/model"
)

func (cli *CLI) addRestaurantCmd() *cobra.Command {
	var restRaw restmodel.Restaurant

	cmd := &cobra.Command{
		Use:   "restaurant",
		Short: "Add a new restaurant",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := cli.restaurantService.Create(cmd.Context(), restRaw)
			if err != nil {
				cmd.PrintErrf("failed to add restaurant: %v\n", err)
			}

			cmd.Printf("Restaurant added with ID %d.\n", id)

		},
	}

	cmd.Flags().StringVarP(&restRaw.Name, "name", "n", "", "Name")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		return nil
	}

	cmd.Flags().StringVarP(&restRaw.Address, "address", "a", "", "Address")
	if err := cmd.MarkFlagRequired("address"); err != nil {
		return nil
	}

	cmd.Flags().StringVarP((*string)(&restRaw.Cuisine), "cuisine", "c", "", "Cuisine")
	if err := cmd.MarkFlagRequired("cuisine"); err != nil {
		return nil
	}

	return cmd
}
