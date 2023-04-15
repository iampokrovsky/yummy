package cli

import (
	"context"
	"github.com/spf13/cobra"
	menumodel "yummy/internal/app/menu/model"
)

func (cli *CLI) getMenuItemCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "menu-item",
		Short: "Get menu item by ID",
		Run: func(cmd *cobra.Command, args []string) {
			item, err := cli.menuService.GetByID(context.Background(), menumodel.ID(id))
			if err != nil {
				cmd.PrintErrf("failed to get menu item: %v\n", err)
			}

			cmd.Printf("Menu item:\n%+v\n", item)
		},
	}

	cmd.Flags().Int64VarP(&id, "id", "i", 0, "Menu item ID")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		return nil
	}

	return cmd
}
