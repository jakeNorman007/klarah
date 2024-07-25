package frameworkTemp

import _"embed"

type StandardLibTemplate struct{}

func (s StandardLibTemplate) Main() []byte {
    return mainTemplate
}
