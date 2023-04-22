package commands

import (
	"context"
	"github.com/spf13/cobra"
	"os"
	utils2 "yummy/cmd/_app_cli/commands/utils"
	restmodel "yummy/internal/app/_restaurant/model"
	menumodel "yummy/internal/app/menu/model"
)

type CoreService interface {
	CreateMenuItem(ctx context.Context, item menumodel.MenuItem) (menumodel.ID, error)
	GetMenuItemByID(ctx context.Context, id menumodel.ID) (menumodel.MenuItem, error)
	ListMenuItemsByName(ctx context.Context, name string) ([]menumodel.MenuItem, error)
	ListMenuItemsByRestaurantID(ctx context.Context, restId menumodel.ID) ([]menumodel.MenuItem, error)
	UpdateMenuItem(ctx context.Context, item menumodel.MenuItem) (bool, error)
	DeleteRestaurant(ctx context.Context, id restmodel.ID) (bool, error)
	RestoreMenuItem(ctx context.Context, id menumodel.ID) (bool, error)
	CreateRestaurant(ctx context.Context, item restmodel.Restaurant) (restmodel.ID, error)
	GetRestaurantByID(ctx context.Context, id restmodel.ID) (restmodel.Restaurant, error)
	ListRestaurantsByName(ctx context.Context, name string) ([]restmodel.Restaurant, error)
	ListRestaurantsByCuisine(ctx context.Context, cuisine string) ([]restmodel.Restaurant, error)
	ListAllRestaurants(ctx context.Context) ([]restmodel.Restaurant, error)
	UpdateRestaurant(ctx context.Context, item restmodel.Restaurant) (bool, error)
	DeleteMenuItem(ctx context.Context, id menumodel.ID) (bool, error)
	RestoreRestaurant(ctx context.Context, id restmodel.ID) (bool, error)
}

type CLI struct {
	coreService CoreService
	rootCommand *cobra.Command
}

// New initializes the CLI object with service methods and returns a pointer to it.
func NewCLI(coreService CoreService) *CLI {
	var cli CLI

	cli = CLI{
		coreService: coreService,
	}

	cli.initCommands()

	return &cli
}

// Execute runs the execution of the CLI command.
func (cli *CLI) Execute(ctx context.Context) {
	err := cli.rootCommand.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func (cli *CLI) initCommands() {
	rootCmd := cli.rootCmd()

	// Turn off "completion" command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Set custom help function for root command
	rootCmd.SetHelpFunc(utils2.RootHelpFunc())

	// Create and configure "add" command
	addCmd := cli.addCmd()
	addCmd.AddCommand(cli.addMenuItemCmd(), cli.addRestaurantCmd())

	// Create and configure "get" command
	getCmd := cli.getCmd()
	getCmd.AddCommand(cli.getMenuItemCmd(), cli.getRestaurantCmd())

	// Create and configure "list" command
	listCmd := cli.listCmd()
	listCmd.AddCommand(cli.listMenuItemsCmd(), cli.listRestaurantsCmd())

	// Create and configure "update" command
	updateCmd := cli.updateCmd()
	updateCmd.AddCommand(cli.updateMenuItemCmd(), cli.updateRestaurantCmd())

	// Create and configure "delete" command
	deleteCmd := cli.deleteCmd()
	deleteCmd.AddCommand(cli.deleteMenuItemCmd(), cli.deleteRestaurantCmd())

	// Create and configure "restore" command
	restoreCmd := cli.restoreCmd()
	restoreCmd.AddCommand(cli.restoreMenuItemCmd(), cli.restoreRestaurantCmd())

	// Create "spell" command
	spellCmd := cli.spellCmd()

	// Create "fmt" command
	fmtCmd := cli.fmtCmd()

	commands := []*cobra.Command{
		addCmd,
		getCmd,
		listCmd,
		updateCmd,
		deleteCmd,
		restoreCmd,
		spellCmd,
		fmtCmd,
	}

	for _, cmd := range commands {
		cmd.SetHelpFunc(utils2.DefaultHelpFunc())
		rootCmd.AddCommand(cmd)

	}

	cli.rootCommand = rootCmd
}
