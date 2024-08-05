package frameworkTemp

import _ "embed"

var ginApiTemplate []byte

var ginHandlersTemplate []byte

var ginMigrationsTemplate []byte

var ginRoutesTemplate []byte

var ginStoresTemplate []byte

var ginTypesTemplate []byte

var ginMiddlewareTemplate []byte

var ginUtilsTemplate []byte

type GinTemplate struct {}

func (e GinTemplate) Main() []byte {
    return mainTemplate
}

func (e GinTemplate) Api() []byte {
    return ginApiTemplate
}

func (e GinTemplate) Handlers() []byte {
    return ginHandlersTemplate
}

func (e GinTemplate) Migrations() []byte {
    return ginMigrationsTemplate
}

func (e GinTemplate) Routes() []byte {
    return ginRoutesTemplate
}

func (e GinTemplate) Stores() []byte {
    return ginStoresTemplate
}

func (e GinTemplate) Types() []byte {
    return ginTypesTemplate
}

func (s GinTemplate) Middleware() []byte {
    return nil
}

func (s GinTemplate) Utils() []byte {
    return nil
}
