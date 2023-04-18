package cli

import (
	"context"
	"github.com/spf13/cobra"
	"os"
	"yummy/cli/utils"
	menumodel "yummy/internal/app/menu/model"
	restmodel "yummy/internal/app/restaurant/model"
)

type ID uint64

type RestaurantService interface {
	Create(ctx context.Context, item restmodel.Restaurant) (restmodel.ID, error)
	GetByID(ctx context.Context, id restmodel.ID) (restmodel.Restaurant, error)
	List(ctx context.Context) ([]restmodel.Restaurant, error)
	ListByName(ctx context.Context, name string) ([]restmodel.Restaurant, error)
	ListByCuisine(ctx context.Context, cuisine string) ([]restmodel.Restaurant, error)
	Update(ctx context.Context, item restmodel.Restaurant) (bool, error)
	Delete(ctx context.Context, id restmodel.ID) (bool, error)
	Restore(ctx context.Context, id restmodel.ID) (bool, error)
}

type MenuService interface {
	Create(ctx context.Context, item menumodel.MenuItem) (menumodel.ID, error)
	GetByID(ctx context.Context, id menumodel.ID) (menumodel.MenuItem, error)
	ListByRestaurantID(ctx context.Context, restId menumodel.ID) ([]menumodel.MenuItem, error)
	ListByName(ctx context.Context, name string) ([]menumodel.MenuItem, error)
	Update(ctx context.Context, item menumodel.MenuItem) (bool, error)
	Delete(ctx context.Context, id menumodel.ID) (bool, error)
	Restore(ctx context.Context, id menumodel.ID) (bool, error)
}

type CLI struct {
	restaurantService RestaurantService
	menuService       MenuService
	rootCommand       *cobra.Command
}

// New initializes the CLI object with service methods and returns a pointer to it.
func New(restaurantService RestaurantService, menuService MenuService) *CLI {
	var cli CLI

	cli = CLI{
		restaurantService: restaurantService,
		menuService:       menuService,
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
	rootCmd.SetHelpFunc(utils.RootHelpFunc())

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
		cmd.SetHelpFunc(utils.DefaultHelpFunc())
		rootCmd.AddCommand(cmd)

	}

	cli.rootCommand = rootCmd
}
