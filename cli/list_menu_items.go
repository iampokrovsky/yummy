package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	menumodel "yummy/internal/app/menu/model"
)

func (cli *CLI) listMenuItemsCmd() *cobra.Command {
	var restId int64

	cmd := &cobra.Command{
		Use:   "menu-items",
		Short: "List menu items by restaurant ID",
		Run: func(cmd *cobra.Command, args []string) {
			items, err := cli.menuService.ListByRestaurantID(cmd.Context(), menumodel.ID(restId))
			if err != nil {
				cmd.PrintErrf("failed to list menu items: %v\n", err)
			}

			fmt.Println("Menu items:")
			for _, item := range items {
				cmd.Printf("%+v\n", item)
			}
		},
	}

	cmd.Flags().Int64VarP(&restId, "rest-id", "r", 0, "Restaurant ID")
	if err := cmd.MarkFlagRequired("rest-id"); err != nil {
		return nil
	}

	return cmd
}
