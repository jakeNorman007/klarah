package frameworkTemp

import _"embed"

//go:embed files/cmd/main.go.tmpl
var mainTemplate []byte

//go:embed files/README.md.tmpl
var readmeTemplate []byte

//go:embed files/makefile.tmpl
var makeTemplate []byte

//go:embed files/gitignore.tmpl
var gitIgnoreTemplate []byte

func MakeTemplate() []byte {
    return makeTemplate
}

func ReadmeTemplate() []byte {
    return readmeTemplate
}

func GitIgnoreTemplate() []byte {
    return gitIgnoreTemplate
}
