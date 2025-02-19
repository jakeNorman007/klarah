package frameworkTemp

import _"embed"

//go:embed files/api/api.go.tmpl
var standardApiTemplate []byte

//go:embed files/api/noDBApi.go.tmpl
var standardNoDBApiTemplate []byte

//go:embed files/handlers/helloWorld_handler.go.tmpl
var standardHandlersTemplate []byte

//go:embed files/middleware/logging.go.tmpl
var standardMiddlewareTemplate []byte

//go:embed files/migrations/001_posts.sql.tmpl
var standardMigrationsTemplate []byte

//go:embed files/routes/posts_routes.go.tmpl
var standardRoutesTemplate []byte

//go:embed files/queries/posts_data.go.tmpl
var standardStoresTemplate []byte

//go:embed files/models/posts.go.tmpl
var standardTypesTemplate []byte

//go:embed files/utils/json_utils.go.tmpl
var standardUtilsTemplate []byte

//go:embed files/tests/handlers_test.go.tmpl
var standardTestsTemplate []byte

type StandardLibTemplate struct{}

func (s StandardLibTemplate) Main() []byte {
    return mainTemplate
}

func (s StandardLibTemplate) MainNoDB() []byte {
    return mainNoDBTemplate
}

func (s StandardLibTemplate) Api() []byte {
    return standardApiTemplate
}

func (s StandardLibTemplate) NoDBApi() []byte {
    return standardNoDBApiTemplate
}

func (s StandardLibTemplate) Handlers() []byte {
    return standardHandlersTemplate
}

func (s StandardLibTemplate) Middleware() []byte {
    return standardMiddlewareTemplate
}

func (s StandardLibTemplate) Migrations() []byte {
    return standardMigrationsTemplate
}

func (s StandardLibTemplate) Routes() []byte {
    return standardRoutesTemplate
}

func (s StandardLibTemplate) Queries() []byte {
    return standardStoresTemplate
}

func (s StandardLibTemplate) Models() []byte {
    return standardTypesTemplate
}

func (s StandardLibTemplate) Utils() []byte {
    return standardUtilsTemplate
}

func (s StandardLibTemplate) Tests() []byte {
    return standardTestsTemplate
}
