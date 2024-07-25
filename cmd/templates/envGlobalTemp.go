package templates

import _"embed"

//go:embed frameworkTemp/files/globalEnvironment.tmpl
var globalEnvironmentVariableTemp []byte

func GlobalEnvironmentVariableTemp() []byte {
    return globalEnvironmentVariableTemp
}
