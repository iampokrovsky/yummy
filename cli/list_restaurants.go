package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (cli *CLI) listRestaurantsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restaurants",
		Short: "List all restaurants",
		Run: func(cmd *cobra.Command, args []string) {
			restaurants, err := cli.restaurantService.List(cmd.Context())
			if err != nil {
				cmd.PrintErrf("failed to list restaurants: %v\n", err)
			}

			fmt.Println("Restaurants:")
			for _, restaurant := range restaurants {
				cmd.Printf("%+v\n", restaurant)
			}
		},
	}

	return cmd
}
