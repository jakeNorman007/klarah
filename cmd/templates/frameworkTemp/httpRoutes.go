package frameworkTemp

import _"embed"

type StandardLibTemplate struct{}

var standardServerTemplate []byte

var standardRoutesTemplate []byte

func (s StandardLibTemplate) Main() []byte {
    return mainTemplate
}

func (s StandardLibTemplate) Server() []byte {
    return standardServerTemplate
}

func (s StandardLibTemplate) Routes() []byte {
    return standardRoutesTemplate
}
