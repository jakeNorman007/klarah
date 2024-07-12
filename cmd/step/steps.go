package step

import (
	"github.com/JakeNorman007/klarah/cmd/flags"
)

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
                        Title: "Standard library",
                        Description: "Built in library for creating http servers",
                    },
                },
                Headers: "Choose your framework",
                Field: frameworkType.String(),
            },
            "database driver": {
                Name: "Project database driver",
                Options: []Item{
                    {
                        Title: "Postgresql",
                        Description: "Go Postgresql database driver",
                    },
                },
                Headers: "Choose a database driver",
                Field: frameworkType.String(),
            },
        },
    }

    return step
}
