package cmd

import (
	"github.com/feitian124/migrate-helper/migrate"
	"github.com/spf13/cobra"
)

var Overwrite bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "migrate-helper [-w] [path]",
	Short: "help you migrate gocheck tests to testify tests",
	Long: `help you migrate gocheck tests to testify tests. For example:
migrate-helper -f gocheck_tests.go
"c.Assert(err, IsNil)" will be changed to "require.NoError(t, err)"`,
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Process(args[0])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolVarP(&Overwrite, "overwrite", "w", false, "Overwrite the original file instead of create new one")
}
