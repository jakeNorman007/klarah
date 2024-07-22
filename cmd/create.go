package cmd

import (
	"os"
	"fmt"
	"log"
	"strings"
	"github.com/spf13/cobra"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/JakeNorman007/klarah/cmd/step"
	"github.com/JakeNorman007/klarah/cmd/flags"
	"github.com/JakeNorman007/klarah/cmd/program"
	"github.com/JakeNorman007/klarah/cmd/utilities"
	"github.com/JakeNorman007/klarah/cmd/ui/mInput"
	"github.com/JakeNorman007/klarah/cmd/ui/textInput"
)

const logo = "Klarah"

var (
    logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("Red"))
    tipMessageStyle = lipgloss.NewStyle()
    endingMessageStyle = lipgloss.NewStyle()
)

func init() {
    var flagFramework flags.Framework
    var flagDBDriver flags.Database
    rootCmd.AddCommand(createCmd)

    createCmd.Flags().StringP("name", "n", "", "Name of project")
    createCmd.Flags().VarP(&flagFramework, "framework", "f", fmt.Sprintf("Frameworks to use: %s", strings.Join(flags.FrameworkTypes, ", ")))
    createCmd.Flags().VarP(&flagDBDriver, "driver", "d", fmt.Sprintf("Databases to use: %s", strings.Join(flags.DatabaseTypes, ", ")))
}

type Options struct {
    ProjectName *textInput.Output
    ProjectType *mInput.Selection
    DBDriver    *mInput.Selection
}

var createCmd = &cobra.Command {
	Use:   "create",
	Short: "Create a project using the klarah scaffolding tool",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
        var tprogram *tea.Program
        var err error

        isInteractive := false
        flagName := cmd.Flag("name").Value.String()
        if flagName != "" && doesDirectoryExistAndIsNotEmpty(flagName) {
            err = fmt.Errorf("%s already exists, is not empty. Choose different name", flagName)
            cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
        }


        flagFramework := flags.Framework(cmd.Flag("framework").Value.String())
		flagDBDriver := flags.Database(cmd.Flag("driver").Value.String())

        options := Options {
            ProjectName:  &textInput.Output{},
            ProjectType:  &mInput.Selection{},
            DBDriver:     &mInput.Selection{},
        }

        project := &program.Project {
            ProjectName:    flagName,
            ProjectType:    flagFramework,
            DBDriver:       flagDBDriver,
            FrameworkMap:   make(map[flags.Framework]program.Framework),
            DBDriverMap:    make(map[flags.Database]program.Driver),

        }

        steps := step.InitStep(flagFramework, flagDBDriver)
        fmt.Printf("%s\n", logoStyle.Render(logo))

        if project.ProjectName == "" {
            isInteractive = true
            tprogram := tea.NewProgram(textInput.InitTextInputModel(options.ProjectName, "project name ", project))
            if _, err := tprogram.Run(); err != nil {
                log.Printf("Project name contains an error: %v", err)
                cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
            }
            if doesDirectoryExistAndIsNotEmpty(options.ProjectName.Output) {
                err = fmt.Errorf("%s already exists, is not empty. Choose different name", options.ProjectName.Output)
                cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
            }

            project.ExitCLI(tprogram)

            project.ProjectName = options.ProjectName.Output
            err := cmd.Flag("name").Value.Set(project.ProjectName)
            if err != nil {
                log.Fatal("Failed to set name flag value", err)
            }
        }

        if project.ProjectType == "" {
            isInteractive = true
            step := steps.Steps["framework"]
            tprogram = tea.NewProgram(mInput.InitModelMulti(step.Options, options.ProjectType, step.Headers, project))
            if _, err := tprogram.Run(); err != nil {
                cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
            }

            project.ExitCLI(tprogram)

            step.Field = options.ProjectType.Choice

            project.ProjectType = flags.Framework(strings.ToLower(options.ProjectType.Choice))
            err := cmd.Flag("framework").Value.Set(project.ProjectType.String())
            if err != nil {
                log.Fatal("Failed to set framework flag value", err)
            }
        }

        if isInteractive {
            nonInteractiveCommand := utilities.NonInteractiveCommand(cmd.Use, cmd.Flags())
            fmt.Println(tipMessageStyle.Render("Tip:"))
            fmt.Println(tipMessageStyle.Italic(false).Render(fmt.Sprintf("%s\n", nonInteractiveCommand)))
        }
    },
}

func doesDirectoryExistAndIsNotEmpty(name string) bool {
    if _, err := os.Stat(name); err == nil {
        directoryEntries, err := os.ReadDir(name)
        if err != nil {
            log.Printf("Could not read directory: %v", err)
        }
        if len(directoryEntries) > 0 {
            return true
        }
    }
    return false
}
