package cli

import (
	"github.com/spf13/cobra"
	menumodel "yummy/internal/app/menu/model"
)

func (cli *CLI) addMenuItemCmd() *cobra.Command {
	var itemRaw menumodel.MenuItem

	cmd := &cobra.Command{
		Use:   "menu-item",
		Short: "Add a new menu item",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := cli.menuService.Create(cmd.Context(), itemRaw)
			if err != nil {
				cmd.PrintErrf("failed to add menu item: %v\n", err)
			}

			cmd.Printf("Menu item added with ID %d.\n", id)
		},
	}

	cmd.Flags().Int64VarP((*int64)(&itemRaw.RestaurantID), "rest-id", "r", 0, "Restaurant ID")
	if err := cmd.MarkFlagRequired("rest-id"); err != nil {
		return nil
	}

	cmd.Flags().StringVarP(&itemRaw.Name, "name", "n", "", "Name")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		return nil
	}

	cmd.Flags().Int64VarP((*int64)(&itemRaw.Price), "price", "p", 0, "Price")
	if err := cmd.MarkFlagRequired("price"); err != nil {
		return nil
	}

	return cmd
}
