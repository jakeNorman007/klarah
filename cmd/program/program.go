package program

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/JakeNorman007/klarah/cmd/flags"
	"github.com/JakeNorman007/klarah/cmd/templates/dbDriverTemp"
	"github.com/JakeNorman007/klarah/cmd/templates/frameworkTemp"
	"github.com/JakeNorman007/klarah/cmd/utilities"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type Project struct {
    ProjectName     string
    Exit            bool
    AbsolutePath    string
    ProjectType     flags.Framework
    DBDriver        flags.Database
    FrameworkMap    map[flags.Framework]Framework
    DBDriverMap     map[flags.Database]Driver
}

type Framework struct {
    packageName []string
    templater   Templater
}

type Driver struct {
    packageName []string
    templater   DBDriverTemplater 
}

type Templater interface {
    Main() []byte
}

type DBDriverTemplater interface {
    Service() []byte
    Env() []byte
}

var (
    postgresqlPackage = []string{"github.com/jackc/pgx/v5/stdlib"}

    godotenvPackage = []string{"github.com/joho/godotenv"}
)

const (
    root = "/"
    apiPath = "/api"
    cmdPath = "/cmd"
    dbPath = "/db"
    handlersPath = "/handlers"
    middlewarePath = "/middleware"
    routesPath = "/routes"
    storesPath = "/stores"
    typesPath = "/types"
    utilsPath = "/utils"
)

func (p *Project) ExitCLI(tprogram *tea.Program) {
    if p.Exit {
        if err := tprogram.ReleaseTerminal(); err != nil {
            log.Fatal(err)
        }

        os.Exit(1)
    }
}

func (p *Project) createFrameworkMap() {
    p.FrameworkMap[flags.StandardLibrary] = Framework {
        packageName: []string{},
        templater: frameworkTemp.StandardLibTemplate{},
    }
}

func (p *Project) createDBDriverMap() {
    p.DBDriverMap[flags.Postgresql] = Driver {
        packageName: postgresqlPackage,
        templater: dbDriverTemp.PostgresqlTemplate{},
    }
}

func (p *Project) CreateMainFile() error {
    if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
        if err := os.Mkdir(p.AbsolutePath, 0o754); err != nil {
            log.Printf("Could not create directory: %v", err)
        }
    }

    p.ProjectName = strings.TrimSpace(p.ProjectName)

    projectPath := filepath.Join(p.AbsolutePath, p.ProjectName)
    if _, err := os.Stat(projectPath); os.IsNotExist(err) {
        err := os.MkdirAll(projectPath, 0o754)
        if err != nil {
            log.Printf("Error creating projects root directory %v\n", err)
            return err
        }
    }

    p.createFrameworkMap()

    err := utilities.InitGoModFile(p.ProjectName, projectPath)
    if err != nil {
        log.Printf("Could not initialize go.mod file in new project: %v\n", err)
        cobra.CheckErr(err)
    }

    if p.ProjectType != flags.StandardLibrary {
        err = utilities.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
        if err != nil {
            log.Printf("Could not install go dependencies for the chosen framework, %v\n", err)
            cobra.CheckErr(err)
        }
    }

    return nil
}

func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
    path := filepath.Join(pathToCreate, projectPath)

    if _, err := os.Stat(path); os.IsNotExist(err) {
        err := os.Mkdir(path, 0o751) //0o751 specifies the permissions: owner full, group: r, x, others: r
        if err != nil {
            log.Printf("Error creating directory %v\n", err)

            return err
        }
    }

    return nil
}

func (p *Project) CreateFileAndInjectTemp(pathToCreate string, projectPath string, fileName string, methodName string) error {
    createdFile, err := os.Create(filepath.Join(projectPath, pathToCreate, fileName, methodName))
    
    if err != nil {
        return err
    }
    
    defer createdFile.Close()

    return nil
}
