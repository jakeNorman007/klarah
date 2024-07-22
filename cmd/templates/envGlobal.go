package templates

import _ "embed"

//go:embed framework/files/envGlobal.tmpl
var globalEnvironmentVariablesTemplate []byte

func EnvGlobalTemplate() []byte {
    return globalEnvironmentVariablesTemplate
}
