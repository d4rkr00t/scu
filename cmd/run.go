package cmd

import (
	"errors"
	"os"
	"path"
	"scu/main/lib"
	"scu/main/lib/cache"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <command>",
	Short: "Run a project's commmand",
	Long:  "Run a project's command",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("command name is required")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		var cwd, cwd_err = cmd.Flags().GetString("cwd")
		var os_cwd, _ = os.Getwd()
		if cwd_err != nil {
			cwd = os_cwd
		} else {
			cwd = path.Join(os_cwd, cwd)
		}

		var verbose, _ = cmd.Flags().GetBool("verbose")

		var pkg_json = lib.NewPackageJson(path.Join(cwd, "package.json"))
		var ctx = lib.NewContext(
			cwd,
			cwd,
			args[0],
			pkg_json,
			cache.NewCache(cwd),
			lib.NewLogger(verbose),
			lib.NewStats(),
			pkg_json.GetConfig(),
		)

		if verbose {
			spew.Dump(ctx)
		}

		lib.Run(ctx)
	},
}