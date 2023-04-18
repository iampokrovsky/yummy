package cli

import (
	"github.com/spf13/cobra"
	menumodel "yummy/internal/app/menu/model"
)

func (cli *CLI) deleteMenuItemCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "menu-item",
		Short: "Delete menu item by ID",
		Run: func(cmd *cobra.Command, args []string) {
			ok, err := cli.menuService.Delete(cmd.Context(), menumodel.ID(id))
			if !ok || err != nil {
				cmd.PrintErrf("failed to delete menu item: %v\n", err)
			}

			cmd.Println("Menu item successfully deleted.")
		},
	}

	cmd.Flags().Int64VarP(&id, "id", "i", 0, "Menu item ID")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		return nil
	}

	return cmd
}
