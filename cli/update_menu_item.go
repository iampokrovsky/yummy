package cli

import (
	"github.com/spf13/cobra"
	menumodel "yummy/internal/app/menu/model"
)

func (cli *CLI) updateMenuItemCmd() *cobra.Command {
	var itemRaw menumodel.MenuItem

	cmd := &cobra.Command{
		Use:   "menu-item",
		Short: "Update menu item by ID",
		Run: func(cmd *cobra.Command, args []string) {
			ok, err := cli.menuService.Update(cmd.Context(), itemRaw)
			if !ok || err != nil {
				cmd.PrintErrf("failed to update menu item: %v\n", err)
			}

			cmd.Println("Menu item successfully updated.")
		},
	}

	cmd.Flags().Int64VarP((*int64)(&itemRaw.ID), "id", "i", 0, "Menu item ID")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		return nil
	}
	cmd.Flags().Int64VarP((*int64)(&itemRaw.RestaurantID), "rest-id", "r", 0, "Restaurant ID")
	cmd.Flags().StringVarP(&itemRaw.Name, "name", "n", "", "Name")
	cmd.Flags().Int64VarP((*int64)(&itemRaw.Price), "price", "p", 0, "Price")

	return cmd
}
