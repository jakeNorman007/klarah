package utilities

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
)

const ProgramName = "klarah"

func NonInteractiveCommand(use string, flagSet *pflag.FlagSet) string {
    nonInteractiveCommand := fmt.Sprintf("%s %s", ProgramName, use)

    visitFn := func(flag *pflag.Flag) {
        if flag.Name != "help" {
            featureFlagString := ""

            for _, k := range strings.Split(flag.Value.String(), ",") {
                if k != "" {
                    featureFlagString += fmt.Sprintf(" --feature %s", k)
                }
            }

            nonInteractiveCommand += featureFlagString
        } else if flag.Value.Type() == "bool" {
            if flag.Value.String() == "true" {
                nonInteractiveCommand = fmt.Sprintf("%s --%s", nonInteractiveCommand, flag.Name)
            }
        } else {
            nonInteractiveCommand = fmt.Sprintf("%s --%s %s", nonInteractiveCommand, flag.Name, flag.Value.String())
        }
    }

    flagSet.SortFlags = false
    flagSet.VisitAll(visitFn)

    return nonInteractiveCommand
}

func ExecuteCmd(name string, args []string, directory string) error {
    command := exec.Command(name, args...)
    command.Dir = directory

    var out bytes.Buffer
    command.Stdout = &out
    if err := command.Run(); err != nil {
        return err
    }

    return nil
}

func InitGoModFile(appDirectory string) error {
    err := ExecuteCmd("go", []string{"mod", "init"}, appDirectory)
    if err != nil {
        return err
    }

    return nil
}

func GoGetPackage(packageName string) error {
    err := ExecuteCmd("go", []string{"get", "-u"}, packageName)
    if err != nil {
        return err
    }

    return nil
}

func GoFormat(appDirectory string) error {
    err := ExecuteCmd("gofmt", []string{"-s", "-w", "."}, appDirectory)
    if err != nil {
        return err
    }

    return nil
}

func GoModTidy(appDirectory string) error {
    err := ExecuteCmd("go", []string{"mod", "tidy"}, appDirectory)
    if err != nil {
        return err
    }

    return nil
}
