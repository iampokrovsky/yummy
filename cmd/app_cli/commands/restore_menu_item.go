package commands

import (
	"github.com/spf13/cobra"
	menumodel "yummy/internal/app/menu/model"
)

func (cli *CLI) restoreMenuItemCmd() *cobra.Command {
	var id int64

	cmd := &cobra.Command{
		Use:   "menu-item",
		Short: "Restore menu item by ID",
		Run: func(cmd *cobra.Command, args []string) {
			ok, err := cli.coreService.RestoreMenuItem(cmd.Context(), menumodel.ID(id))
			if !ok || err != nil {
				cmd.PrintErrf("failed to restore menu item: %v\n", err)
			}

			cmd.Println("Menu item successfully restored.")
		},
	}

	cmd.Flags().Int64VarP(&id, "id", "i", 0, "Menu item ID")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		return nil
	}

	return cmd
}
