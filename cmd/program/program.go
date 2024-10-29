package program

import (
  "log"
  "os"
  "path/filepath"
  "strings"
  "text/template"
  "github.com/JakeNorman007/klarah/cmd/flags"
  tpl "github.com/JakeNorman007/klarah/cmd/templates"
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
  Main()          []byte
  MainNoDB()      []byte
  Api()           []byte
  NoDBApi()       []byte
  Handlers()      []byte
  Middleware()    []byte
  Migrations()    []byte
  Routes()        []byte
  Queries()       []byte
  Tests()         []byte
  Models()        []byte
  Utils()         []byte
}

type DBDriverTemplater interface {
  Service() []byte
}

var (
  echoPackage = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}
  chiPackage = []string{"github.com/go-chi/chi/v5", "github.com/go-chi/chi/v5/middleware"}
  ginPackage = []string{"github.com/gin-gonic/gin"}

  postgresqlPackage = []string{"github.com/jackc/pgx/v5/stdlib"}
  sqlitePackage = []string{"github.com/mattn/go-sqlite3"}

  godotenvPackage = []string{"github.com/joho/godotenv"}
  goosePackage = []string{"github.com/pressly/goose/v3/cmd/goose@latest"}
)

const (
  root = "/"
  apiPath = "api"
  cmdPath = "cmd"
  dbPath = "db"
  handlersPath = "handlers"
  middlewarePath = "middleware"
  migrationsPath = "migrations"
  routesPath = "routes"
  queriesPath = "queries"
  testsPath = "tests"
  modelsPath = "models"
  utilsPath = "utils"
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
  p.FrameworkMap[flags.Echo] = Framework {
    packageName: echoPackage,
    templater: frameworkTemp.EchoTemplate{},
  }
  p.FrameworkMap[flags.Chi] = Framework {
    packageName: chiPackage,
    templater: frameworkTemp.ChiTemplate{},
  }
  p.FrameworkMap[flags.Gin] = Framework {
    packageName: ginPackage,
    templater: frameworkTemp.GinTemplate{},
  }
}

func (p *Project) createDBDriverMap() {
  p.DBDriverMap[flags.Postgresql] = Driver {
    packageName: postgresqlPackage,
    templater: dbDriverTemp.PostgresqlTemplate{},
  }
  p.DBDriverMap[flags.Sqlite] = Driver {
    packageName: sqlitePackage,
    templater: dbDriverTemp.SqliteTemplate{},
  }
}

func (p *Project) CreateMainFile() error {
  if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
    if err := os.Mkdir(p.AbsolutePath, 0o754); err != nil {
      log.Printf("Could not create directory: %v", err)
      return err
    }
  }

  p.ProjectName = strings.TrimSpace(p.ProjectName)

  projectPath := filepath.Join(p.AbsolutePath, p.ProjectName)
  if _, err := os.Stat(projectPath); os.IsNotExist(err) {
    err := os.MkdirAll(projectPath, 0o751)
    if err != nil {
      log.Printf("Error creating projects root directory %v", err)
      return err
    }
  }

  p.createFrameworkMap()

  err := utilities.InitGoModFile(p.ProjectName, projectPath)
  if err != nil {
    log.Printf("Could not initialize go.mod file in new project: %v", err)
    cobra.CheckErr(err)
  }

  if p.ProjectType != flags.StandardLibrary {
    err = utilities.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
    if err != nil {
      log.Printf("Could not install go dependencies for the chosen framework, %v", err)
      cobra.CheckErr(err)
    }
  }

  if p.DBDriver != "none" {
    err = p.CreatePath(apiPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", apiPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(apiPath, projectPath, "api.go", "api")
    if err != nil {
      log.Printf("Error injecting api.go file: %s", apiPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(cmdPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", cmdPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(cmdPath, projectPath, "main.go", "main")
    if err != nil {
      log.Printf("Error injecting main.go file: %s", cmdPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(handlersPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", handlersPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(handlersPath, projectPath, "helloWorld_handler.go", "handlers")
    if err != nil {
      log.Printf("Error injecting handlers.go file: %s", handlersPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(testsPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", testsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(testsPath, projectPath, "handlers_test.go", "tests")
    if err != nil {
      log.Printf("Error injecting handlers_test.go file: %s", testsPath)
      cobra.CheckErr(err)
      return err
    }

    p.createDBDriverMap()
    err = utilities.GoGetPackage(projectPath, p.DBDriverMap[p.DBDriver].packageName)
    if err != nil {
      log.Printf("Could not install dependency for chosen driver %v", err)
      cobra.CheckErr(err)
    }

    err = p.CreatePath(dbPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", dbPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(dbPath, projectPath, "database.go", "database")
    if err != nil {
      log.Printf("Error injecting database.go file: %s", dbPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(modelsPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", modelsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(modelsPath, projectPath, "posts.go", "models")
    if err != nil {
      log.Printf("Error injecting posts.go file: %s", modelsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(migrationsPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", migrationsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(migrationsPath, projectPath, "001_posts.sql", "migrations")
    if err != nil {
      log.Printf("Error injecting 001_posts.sql file: %s", migrationsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(routesPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", routesPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(routesPath, projectPath, "posts_routes.go", "routes")
    if err != nil {
      log.Printf("Error injecting posts_routes.go file: %s", routesPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(queriesPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", queriesPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(queriesPath, projectPath, "posts_data.go", "queries")
    if err != nil {
      log.Printf("Error injecting posts_data.go file: %s", queriesPath)
      cobra.CheckErr(err)
      return err
    }
  }

  if p.DBDriver == "none" {
    err = p.CreatePath(apiPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", apiPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(apiPath, projectPath, "api.go", "noApi")
    if err != nil {
      log.Printf("Error injecting api.go file: %s", apiPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(cmdPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", cmdPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(cmdPath, projectPath, "main.go", "noMain")
    if err != nil {
      log.Printf("Error injecting main.go file: %s", cmdPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(handlersPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", handlersPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(handlersPath, projectPath, "helloWorld_handler.go", "handlers")
    if err != nil {
      log.Printf("Error injecting handlers.go file: %s", handlersPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(testsPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", testsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(testsPath, projectPath, "handlers_test.go", "tests")
    if err != nil {
      log.Printf("Error injecting handlers_test.go file: %s", testsPath)
      cobra.CheckErr(err)
      return err
    }
  }

  err = utilities.GoGetPackage(projectPath, godotenvPackage)
  if err != nil {
    log.Printf("Could not install dependency: %v", err)
    cobra.CheckErr(err)
  }

  err = utilities.GoGetPackage(projectPath, goosePackage)
  if err != nil {
    log.Printf("Could not install dependency: %v", err)
    cobra.CheckErr(err)
  }

  if p.ProjectType == "standard-library" {
    err = p.CreatePath(middlewarePath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", middlewarePath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(middlewarePath, projectPath, "logging.go", "middleware")
    if err != nil {
      log.Printf("Error injecting logging.go file: %s", middlewarePath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreatePath(utilsPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", utilsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(utilsPath, projectPath, "json_utils.go", "utils")
    if err != nil {
      log.Printf("Error injecting json_utils.go file: %s", utilsPath)
      cobra.CheckErr(err)
      return err
    }

  } else if p.ProjectType == "chi" {
    err = p.CreatePath(utilsPath, projectPath)
    if err != nil {
      log.Printf("Error in creating path: %s", utilsPath)
      cobra.CheckErr(err)
      return err
    }

    err = p.CreateFileAndInjectTemp(utilsPath, projectPath, "json_utils.go", "utils")
    if err != nil {
      log.Printf("Error injecting json_utils.go file: %s", utilsPath)
      cobra.CheckErr(err)
      return err
    }
  }

  err = p.CreateFileAndInjectTemp(root, projectPath, ".env", "env")
  if err != nil {
    log.Printf("Error injecting .env file: %v", err)
    cobra.CheckErr(err)
    return err
  }

  makeFile, err := os.Create(filepath.Join(projectPath, "Makefile"))
  if err != nil {
    cobra.CheckErr(err)
    return err
  }

  defer makeFile.Close()

  makeFileTemplate := template.Must(template.New("makefile").Parse(string(frameworkTemp.MakeTemplate())))
  err = makeFileTemplate.Execute(makeFile, p)
  if err != nil {
    return err
  }

  readmeFile, err := os.Create(filepath.Join(projectPath, "README.md"))
  if err != nil {
    cobra.CheckErr(err)
    return err
  }

  defer readmeFile.Close()

  readmeFileTemplate := template.Must(template.New("readme").Parse(string(frameworkTemp.ReadmeTemplate())))
  err = readmeFileTemplate.Execute(readmeFile, p)
  if err != nil {
    return err
  }

  gitignoreFile, err := os.Create(filepath.Join(projectPath, ".gitignore"))
  if err != nil {
    cobra.CheckErr(err)
    return err
  }

  defer gitignoreFile.Close()

  gitignoreFileTemplate := template.Must(template.New(".gitignore").Parse(string(frameworkTemp.GitIgnoreTemplate())))
  err = gitignoreFileTemplate.Execute(gitignoreFile, p)
  if err != nil {
    return err
  }

  envFile, err := os.Create(filepath.Join(projectPath, ".env"))
  if err != nil {
    cobra.CheckErr(err)
    return err
  }

  defer envFile.Close()

  envFileTemplate := template.Must(template.New(".env").Parse(string(tpl.GlobalEnvironmentVariableTemp())))
  err = envFileTemplate.Execute(envFile, p)
  if err != nil {
    return err
  }

  err = utilities.GoModTidy(projectPath)
  if err != nil {
    log.Printf("Could not go tidy in project: %v", err)
    cobra.CheckErr(err)
  }

  return nil
}

func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
  path := filepath.Join(projectPath, pathToCreate)
  if _, err := os.Stat(path); os.IsNotExist(err) {
    err := os.MkdirAll(path, 0o751) //0o751 specifies the permissions: owner full, group: r, x, others: r
    if err != nil {
      log.Printf("Error creating directory %v\n", err)

      return err
    }
  }

  return nil
}

func (p *Project) CreateFileAndInjectTemp(pathToCreate string, projectPath string, fileName string, methodName string) error {
  createdFile, err := os.Create(filepath.Join(projectPath, pathToCreate, fileName))
  if err != nil {
    return err
  }

  defer createdFile.Close()

  switch methodName {
  case "main":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Main())))
    err = createdTemplate.Execute(createdFile, p)
  case "noMain":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.MainNoDB())))
    err = createdTemplate.Execute(createdFile, p)
  case "database":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.DBDriverMap[p.DBDriver].templater.Service())))
    err = createdTemplate.Execute(createdFile, p)
  case "api":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Api())))
    err = createdTemplate.Execute(createdFile, p)
  case "noApi":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.NoDBApi())))
    err = createdTemplate.Execute(createdFile, p)
  case "handlers":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Handlers())))
    err = createdTemplate.Execute(createdFile, p)
  case "middleware":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Middleware())))
    err = createdTemplate.Execute(createdFile, p)
  case "migrations":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Migrations())))
    err = createdTemplate.Execute(createdFile, p)
  case "routes":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Routes())))
    err = createdTemplate.Execute(createdFile, p)
  case "queries":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Queries())))
    err = createdTemplate.Execute(createdFile, p)
  case "tests":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Tests())))
    err = createdTemplate.Execute(createdFile, p)
  case "models":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Models())))
    err = createdTemplate.Execute(createdFile, p)
  case "utils":
    createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Utils())))
    err = createdTemplate.Execute(createdFile, p)
  }

  if err != nil {
    return err
  }

  return nil
}
