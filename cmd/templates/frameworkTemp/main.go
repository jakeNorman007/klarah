package frameworktemp

import _"embed"

var mainTemplate []byte

var readmeTemplate []byte

var makeTemplate []byte

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
