package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (cli *CLI) listRestaurantsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restaurants",
		Short: "ListAll all restaurants",
		Run: func(cmd *cobra.Command, args []string) {
			restaurants, err := cli.coreService.ListAllRestaurants(cmd.Context())
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
