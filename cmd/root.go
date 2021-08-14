package cmd

import (
	"github.com/feitian124/migrate-helper/migrate"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "migrate-helper [path]",
	Short: "help you migrate gocheck tests to testify tests",
	Long: `help you migrate gocheck tests to testify tests. For example:
migrate-helper -f gocheck_tests.go
"c.Assert(err, IsNil)" will be changed to "require.NoError(t, err)"`,
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
	//rootCmd.Flags().StringVarP(&path, "file", "f", "", "goChecker test file path to migrate, can be file or folder")
	//rootCmd.MarkFlagRequired("file")
}
