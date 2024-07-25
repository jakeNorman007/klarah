package program

import (
	"log"
	"os"
	"path/filepath"

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
    standardLibraryPackage = []string{""}

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

func (p *Project) CreateMainFile() error {
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

func (p *Project) CreateFileAndInjectTemp() error {
    return nil
}
