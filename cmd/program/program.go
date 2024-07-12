package program

import (
	"log"
	"os"

	"github.com/JakeNorman007/klarah/cmd/flags"
	tea "github.com/charmbracelet/bubbletea"
)

type Project struct {
    ProjectName     string
    Exit            bool
    //AbsolutePath  string
    ProjectType     flags.Framework
    DBDriver        flags.Database
    FrameworkMap    map[flags.Framework]Framework
    DBDriverMap     map[flags.Database]Driver
}

type Framework struct {
    packageName string
    templater   Templater
}

type Driver struct {
    packageName []string
    templater   DBDriverTemplater 
}

type Templater interface {
}

type DBDriverTemplater interface {
}

const (
    //here will go the path variables
)

func (p *Project) ExitCLI(tprogram *tea.Program) {
    if p.Exit {
        if err := tprogram.ReleaseTerminal(); err != nil {
            log.Fatal(err)
        }

        os.Exit(1)
    }
}
