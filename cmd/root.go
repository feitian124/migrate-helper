package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var File string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "migrate-helper",
	Short: "help you migrate gocheck tests to testify tests",
	Long: `help you migrate gocheck tests to testify tests. For example:
migrate-helper -f gocheck_tests.go
"c.Assert(err, IsNil)" will be changed to "require.NoError(t, err)"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args: " + strings.Join(args, " "))
		fmt.Println("file: " + File)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVarP(&File, "file", "f", "", "goChecker test file to migrate")
	rootCmd.MarkFlagRequired("file")
}

