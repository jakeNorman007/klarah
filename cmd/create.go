package cmd

import (
	"github.com/JakeNorman007/klarah/cmd/flags"
	"github.com/JakeNorman007/klarah/cmd/step"
	"github.com/spf13/cobra"
)

func Init() {
    //var flagFramework flags.Framework
    //var flagDBDriiver flags.Database
    rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command {
	Use:   "create",
	Short: "",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {

        flagFramework := flags.Framework(cmd.Flag("framework").Value.String())
		flagDBDriver := flags.Database(cmd.Flag("driver").Value.String())

        steps := step.InitStep(flagFramework, flagDBDriver)
    },
}
