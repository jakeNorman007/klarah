package cmd

import (
	"os"
	"fmt"
	"log"
	"sync"
	"strings"
	"github.com/spf13/cobra"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/JakeNorman007/klarah/cmd/step"
	"github.com/JakeNorman007/klarah/cmd/flags"
	"github.com/JakeNorman007/klarah/cmd/program"
	"github.com/JakeNorman007/klarah/cmd/ui/mInput"
	"github.com/JakeNorman007/klarah/cmd/ui/spinner"
	"github.com/JakeNorman007/klarah/cmd/ui/textInput"
)

const logo = `
               __      __                            __
              |  |  __|  |                          |  | 
              |  | /  /  |                          |  | 
              |  |/  /|  |                          |  |
              |  /  / |  |    _____    ____  _____  |  |____
              |    /  |  |   /  _  \  /  __|/  _  \ |       \
              |    \  |  |  |  / \  \|  /  |  / \  \|  |---  \
              |  \  \ |  |_ |  | |  ||  |  |  | |  ||  |  |  |
              |  |\  \|    ||  \_/  ||  |  |  \_/  ||  |  |  |
              |__| \__\____| \____/|||__|   \____/|||__|  |__|

                      *** Build backends quickly ***
            `


var (
    logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#59DA00")).Bold(true)
	endingMsgStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#50DA00")).Bold(true)
)

func init() {
    var flagFramework flags.Framework
    var flagDBDriver flags.Database

    rootCmd.AddCommand(createCmd)

    createCmd.Flags().StringP("name", "n", "", "Name of project")
    createCmd.Flags().VarP(&flagFramework, "framework", "f", fmt.Sprintf("Frameworks to use: %s", strings.Join(flags.FrameworkTypes, ", ")))
    createCmd.Flags().VarP(&flagDBDriver, "database driver", "d", fmt.Sprintf("Databases to use: %s", strings.Join(flags.DatabaseTypes, ", ")))
}

type Options struct {
    ProjectName *textInput.Output
    ProjectType *mInput.Selection
    DBDriver    *mInput.Selection
}

var createCmd = &cobra.Command {
	Use:   "new",
	Short: "Create a project using the klarah scaffolding tool",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
        var tprogram *tea.Program
        var err error

        flagName := cmd.Flag("name").Value.String()
        if flagName != "" && doesDirectoryExistAndIsNotEmpty(flagName) {
            err = fmt.Errorf("%s already exists, is not empty. Choose different name", flagName)
            cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
        }


        flagFramework := flags.Framework(cmd.Flag("framework").Value.String())
		flagDBDriver := flags.Database(cmd.Flag("database driver").Value.String())

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

        if project.DBDriver == "" {
            step := steps.Steps["database driver"]
            tprogram = tea.NewProgram(mInput.InitModelMulti(step.Options, options.DBDriver, step.Headers, project))
            if _, err := tprogram.Run(); err != nil {
                cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
            }

            project.ExitCLI(tprogram)

            project.DBDriver = flags.Database(strings.ToLower(options.DBDriver.Choice))
            err := cmd.Flag("database driver").Value.Set(project.DBDriver.String())
            if err != nil {
                log.Fatal("Failed to set the driver flag value", err)
            }
        }

        currentWorkingDirectory, err := os.Getwd()
        if err != nil {
            log.Printf("Could not get current working directory: %v", err)
            cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
        }

        project.AbsolutePath = currentWorkingDirectory

        spinner := tea.NewProgram(spinner.InitialModelNew())

        wg := sync.WaitGroup{}
        wg.Add(1)

        go func() {
            defer wg.Done()
            if _, err := spinner.Run(); err != nil {
                cobra.CheckErr(err)
            }
        }()

        defer func() {
			if r := recover(); r != nil {
				fmt.Println("The program encountered an unexpected issue and had to exit. The error was:", r)//better error goes here
				fmt.Println("If you continue to experience this issue, please post a message to our GitHub")
				if releaseErr := spinner.ReleaseTerminal(); releaseErr != nil {
					log.Printf("Problem releasing terminal: %v", releaseErr)
				}
			}
		}()

        err = project.CreateMainFile()
        if err != nil {
            if releaseErr := spinner.ReleaseTerminal(); releaseErr != nil {
                log.Printf("Issue releasing termainl: %v", err)
            }

            log.Printf("Issue creating files for project: %v", err)
            cobra.CheckErr(textInput.CreateErrorInputModel(err).Err())
        }

        fmt.Println(endingMsgStyle.Render(fmt.Sprintf("Next steps: cd into your newly created project with: cd %s", project.ProjectName)))

        err = spinner.ReleaseTerminal()
        if err != nil {
            log.Printf("Could not release terminal: %v", err)
            cobra.CheckErr(err)
        }
    },
}

func doesDirectoryExistAndIsNotEmpty(name string) bool {
    if _, err := os.Stat(name); err == nil {
        directoryEntries, err := os.ReadDir(name)
        if err != nil {
            log.Printf("Could not read directory: %v", err)
            cobra.CheckErr(textInput.CreateErrorInputModel(err))
        }
        if len(directoryEntries) > 0 {
            return true
        }
    }

    return false
}
