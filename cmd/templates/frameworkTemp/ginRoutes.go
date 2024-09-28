package frameworkTemp

import _ "embed"

//go:embed files/api/ginApi.go.tmpl
var ginApiTemplate []byte

//go:embed
var ginNoDBApiTemplate []byte

//go:embed files/handlers/ginHelloWorld_handler.go.tmpl
var ginHandlersTemplate []byte

//go:embed files/migrations/001_posts.sql.tmpl
var ginMigrationsTemplate []byte

//go:embed files/routes/gin_posts_routes.go.tmpl
var ginRoutesTemplate []byte

//go:embed files/stores/posts_data.go.tmpl
var ginStoresTemplate []byte

//go:embed files/types/posts.go.tmpl
var ginTypesTemplate []byte

var ginMiddlewareTemplate []byte

var ginUtilsTemplate []byte

//go:embed files/tests/ginHandlers_test.go.tmpl
var ginTestsTemplate []byte

type GinTemplate struct {}

func (e GinTemplate) Main() []byte {
    return mainTemplate
}

func (e GinTemplate) Api() []byte {
    return ginApiTemplate
}

func (e GinTemplate) NoDBApi() []byte {
    return ginNoDBApiTemplate
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

func (s GinTemplate) Tests() []byte {
    return ginTestsTemplate
}
