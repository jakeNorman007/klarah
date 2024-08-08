package frameworkTemp

import _ "embed"

//go:embed files/api/chiApi.go.tmpl
var chiApiTemplate []byte

//go:embed files/handlers/helloWorld_handler.go.tmpl
var chiHandlersTemplate []byte

var chiMiddlewareTemplate []byte

//go:embed files/migrations/001_posts.sql.tmpl
var chiMigrationsTemplate []byte

//go:embed files/routes/chi_posts_routes.go.tmpl
var chiRoutesTemplate []byte

//go:embed files/stores/posts_data.go.tmpl
var chiStoresTemplate []byte

//go:embed files/types/posts.go.tmpl
var chiTypesTemplate []byte

//go:embed files/utils/json_utils.go.tmpl
var chiUtilsTemplate []byte

type ChiTemplate struct {}

func (e ChiTemplate) Main() []byte {
    return mainTemplate
}

func (e ChiTemplate) Api() []byte {
    return chiApiTemplate
}

func (e ChiTemplate) Handlers() []byte {
    return chiHandlersTemplate
}

func (e ChiTemplate) Migrations() []byte {
    return chiMigrationsTemplate
}

func (e ChiTemplate) Routes() []byte {
    return chiRoutesTemplate
}

func (e ChiTemplate) Stores() []byte {
    return chiStoresTemplate
}

func (e ChiTemplate) Types() []byte {
    return chiTypesTemplate
}

func (s ChiTemplate) Middleware() []byte {
    return nil
}

func (s ChiTemplate) Utils() []byte {
    return chiUtilsTemplate
}

