package step

import "github.com/JakeNorman007/klarah/cmd/flags"

type Schema struct {
    Name    string
    Options []Item
    Headers string
    Field   string
}

type Step struct {
    Steps map[string]Schema
}

type Item struct {
    Flag        string
    Title       string
    Description string
}

func InitStep(frameworkType flags.Framework, databaseType flags.Database) *Step {
    step := &Step {
        map[string]Schema {
            "framework": {
                Name: "Project framework",
                Options: []Item{
                    {
                        Title: "standard-library",
                        Description: "Standard go library for creating http servers",
                    },
                    {
                        Title: "echo",
                        Description: "High performing, extensible, framework with minimal overhead",
                    },
                    {
                        Title: "chi",
                        Description: "Lightweight, fast http router that's flexible and powerfull",
                    },
                },
                Headers: "----------------------------- Frameworks -----------------------------",
                Field: frameworkType.String(),
            },

            "database driver": {
                Name: "Project database driver",
                Options: []Item{
                    {
                        Title: "postgresql",
                        Description: "Powerful, open-source relational database management system",
                    },
                    {
                        Title: "sqlite",
                        Description: "Lightweight, self-contained SQL database engine",
                    },
                },
                Headers: "-------------------------- Database drivers --------------------------",
                Field: frameworkType.String(),
            },
        },
    }

    return step
}
